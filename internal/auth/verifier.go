package auth

import (
	"context"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/MicahParks/keyfunc/v3"
	"github.com/golang-jwt/jwt/v5"
)

// Verifier checks signature and standard claims of tokens issued by an external IdP.
type Verifier struct {
	parser  *jwt.Parser
	keyfunc jwt.Keyfunc
}

// NewVerifier builds a verifier from cfg. ctx bounds the background JWKS refresh.
func NewVerifier(ctx context.Context, cfg *Config) (*Verifier, error) {
	kf, err := newKeyfunc(ctx, cfg)
	if err != nil {
		return nil, err
	}

	opts := []jwt.ParserOption{
		jwt.WithValidMethods(cfg.Algorithms),
		jwt.WithIssuer(cfg.Issuer),
		jwt.WithLeeway(cfg.ClockSkew),
		jwt.WithExpirationRequired(),
	}
	if cfg.Audience != "" {
		opts = append(opts, jwt.WithAudience(cfg.Audience))
	}

	return &Verifier{parser: jwt.NewParser(opts...), keyfunc: kf}, nil
}

func (v *Verifier) Verify(raw string) (jwt.MapClaims, error) {
	claims := jwt.MapClaims{}
	if _, err := v.parser.ParseWithClaims(raw, claims, v.keyfunc); err != nil {
		return nil, err
	}
	return claims, nil
}

func newKeyfunc(ctx context.Context, cfg *Config) (jwt.Keyfunc, error) {
	if cfg.PublicKeyFile != "" {
		key, err := loadPublicKey(cfg.PublicKeyFile)
		if err != nil {
			return nil, err
		}
		return func(*jwt.Token) (any, error) { return key, nil }, nil
	}

	jwksURL := cfg.JWKSURL
	if jwksURL == "" {
		discovered, err := discoverJWKSURL(ctx, cfg.OIDCIssuerURL)
		if err != nil {
			return nil, err
		}
		jwksURL = discovered
	}

	// Keys are kept in memory and refreshed in background and on unknown kid,
	// so signature verification never goes to the network on the request path.
	k, err := keyfunc.NewDefaultOverrideCtx(ctx, []string{jwksURL}, keyfunc.Override{
		RefreshInterval: cfg.JWKSRefreshInterval,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to init JWKS from %s: %w", jwksURL, err)
	}
	return k.Keyfunc, nil
}

func discoverJWKSURL(ctx context.Context, issuerURL string) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	u := strings.TrimSuffix(issuerURL, "/") + "/.well-known/openid-configuration"
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u, nil)
	if err != nil {
		return "", err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("OIDC discovery %s: %w", u, err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("OIDC discovery %s: unexpected status %d", u, resp.StatusCode)
	}

	var doc struct {
		JWKSURI string `json:"jwks_uri"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&doc); err != nil {
		return "", fmt.Errorf("OIDC discovery %s: %w", u, err)
	}
	if doc.JWKSURI == "" {
		return "", fmt.Errorf("OIDC discovery %s: jwks_uri is empty", u)
	}
	return doc.JWKSURI, nil
}

func loadPublicKey(path string) (any, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read public key: %w", err)
	}
	block, _ := pem.Decode(data)
	if block == nil {
		return nil, fmt.Errorf("public key %s: no PEM block found", path)
	}
	switch block.Type {
	case "PUBLIC KEY":
		return x509.ParsePKIXPublicKey(block.Bytes)
	case "RSA PUBLIC KEY":
		return x509.ParsePKCS1PublicKey(block.Bytes)
	case "CERTIFICATE":
		cert, err := x509.ParseCertificate(block.Bytes)
		if err != nil {
			return nil, err
		}
		return cert.PublicKey, nil
	default:
		return nil, fmt.Errorf("public key %s: unsupported PEM block %q", path, block.Type)
	}
}
