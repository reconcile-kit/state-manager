package config

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/reconcile-kit/state-manager/internal/auth"
)

// AuthConfig reads AUTH_* variables. Without AUTH_ENABLED=true authorization is disabled
// and the service works without tokens.
func AuthConfig() (*auth.Config, error) {
	enabled, err := parseBool("AUTH_ENABLED")
	if err != nil || !enabled {
		return &auth.Config{Enabled: false}, err
	}

	cfg := &auth.Config{
		Enabled:           true,
		OIDCIssuerURL:     getEnv("AUTH_OIDC_ISSUER_URL", ""),
		JWKSURL:           getEnv("AUTH_JWKS_URL", ""),
		PublicKeyFile:     getEnv("AUTH_JWT_PUBLIC_KEY_FILE", ""),
		Issuer:            getEnv("AUTH_ISSUER", getEnv("AUTH_OIDC_ISSUER_URL", "")),
		Audience:          getEnv("AUTH_AUDIENCE", ""),
		Algorithms:        splitList(getEnv("AUTH_ALGORITHMS", "RS256,ES256")),
		SubjectClaim:      getEnv("AUTH_SUBJECT_CLAIM", "sub"),
		GroupsClaim:       getEnv("AUTH_GROUPS_CLAIM", "groups"),
		PermissionsClaim:  getEnv("AUTH_PERMISSIONS_CLAIM", "permissions"),
		PermissionsSource: auth.PermissionsSource(getEnv("AUTH_PERMISSIONS_SOURCE", string(auth.SourceAuto))),
	}
	if cfg.ClockSkew, err = parseDuration("AUTH_CLOCK_SKEW", "30s"); err != nil {
		return nil, err
	}
	if cfg.JWKSRefreshInterval, err = parseDuration("AUTH_JWKS_REFRESH_INTERVAL", "1h"); err != nil {
		return nil, err
	}
	if err = cfg.Validate(); err != nil {
		return nil, err
	}
	return cfg, nil
}

func parseBool(key string) (bool, error) {
	v := getEnv(key, "false")
	b, err := strconv.ParseBool(v)
	if err != nil {
		return false, fmt.Errorf("invalid %s %q: %w", key, v, err)
	}
	return b, nil
}

func parseDuration(key, def string) (time.Duration, error) {
	v := getEnv(key, def)
	d, err := time.ParseDuration(v)
	if err != nil {
		return 0, fmt.Errorf("invalid %s %q: %w", key, v, err)
	}
	return d, nil
}

func splitList(s string) []string {
	var out []string
	for _, item := range strings.Split(s, ",") {
		if item = strings.TrimSpace(item); item != "" {
			out = append(out, item)
		}
	}
	return out
}
