package auth

import (
	"fmt"
	"time"
)

// PermissionsSource defines where the rules of a token are taken from.
type PermissionsSource string

const (
	// SourceAuto uses the permissions claim when the token has it and the database otherwise.
	SourceAuto PermissionsSource = "auto"
	// SourceClaims uses only the permissions claim, the database is never queried.
	SourceClaims PermissionsSource = "claims"
	// SourceDB uses only role bindings from the database.
	SourceDB PermissionsSource = "db"
)

// Asymmetric algorithms only: tokens are issued by a third party,
// so the service must never hold a signing secret.
var supportedAlgorithms = map[string]struct{}{
	"RS256": {}, "RS384": {}, "RS512": {},
	"PS256": {}, "PS384": {}, "PS512": {},
	"ES256": {}, "ES384": {}, "ES512": {},
	"EdDSA": {},
}

type Config struct {
	Enabled bool

	// Exactly one key source must be set.
	OIDCIssuerURL       string
	JWKSURL             string
	PublicKeyFile       string
	JWKSRefreshInterval time.Duration

	Issuer     string
	Audience   string
	Algorithms []string
	ClockSkew  time.Duration

	SubjectClaim     string
	GroupsClaim      string
	PermissionsClaim string

	PermissionsSource PermissionsSource
}

func (c *Config) Validate() error {
	if !c.Enabled {
		return nil
	}

	keySources := 0
	for _, s := range []string{c.OIDCIssuerURL, c.JWKSURL, c.PublicKeyFile} {
		if s != "" {
			keySources++
		}
	}
	if keySources != 1 {
		return fmt.Errorf("exactly one of AUTH_OIDC_ISSUER_URL, AUTH_JWKS_URL, AUTH_JWT_PUBLIC_KEY_FILE must be set")
	}
	if c.Issuer == "" {
		return fmt.Errorf("AUTH_ISSUER is required")
	}
	if len(c.Algorithms) == 0 {
		return fmt.Errorf("AUTH_ALGORITHMS must not be empty")
	}
	for _, alg := range c.Algorithms {
		if _, ok := supportedAlgorithms[alg]; !ok {
			return fmt.Errorf("unsupported algorithm %q in AUTH_ALGORITHMS", alg)
		}
	}
	if c.SubjectClaim == "" {
		return fmt.Errorf("AUTH_SUBJECT_CLAIM must not be empty")
	}
	switch c.PermissionsSource {
	case SourceAuto, SourceClaims:
		if c.PermissionsClaim == "" {
			return fmt.Errorf("AUTH_PERMISSIONS_CLAIM must not be empty for permissions source %q", c.PermissionsSource)
		}
	case SourceDB:
	default:
		return fmt.Errorf("unknown AUTH_PERMISSIONS_SOURCE %q, expected auto, claims or db", c.PermissionsSource)
	}
	return nil
}
