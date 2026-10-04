package auth

import (
	"context"
	"fmt"
	"strings"
)

// RuleStore loads rules bound to a subject or any of its groups.
type RuleStore interface {
	Rules(ctx context.Context, subject string, groups []string) ([]Rule, error)
}

// Authenticator turns a raw bearer token into a Principal with its rules.
type Authenticator struct {
	verifier *Verifier
	store    RuleStore
	source   PermissionsSource

	subjectPath     []string
	groupsPath      []string
	permissionsPath []string
}

// NewAuthenticator creates an authenticator. store may be nil only for SourceClaims.
func NewAuthenticator(cfg *Config, verifier *Verifier, store RuleStore) (*Authenticator, error) {
	if cfg.PermissionsSource != SourceClaims && store == nil {
		return nil, fmt.Errorf("rule store is required for permissions source %q", cfg.PermissionsSource)
	}
	return &Authenticator{
		verifier:        verifier,
		store:           store,
		source:          cfg.PermissionsSource,
		subjectPath:     claimPath(cfg.SubjectClaim),
		groupsPath:      claimPath(cfg.GroupsClaim),
		permissionsPath: claimPath(cfg.PermissionsClaim),
	}, nil
}

// Authenticate verifies the token and resolves rules.
// Token problems are wrapped into ErrUnauthenticated, store failures are returned as is.
func (a *Authenticator) Authenticate(ctx context.Context, raw string) (*Principal, error) {
	claims, err := a.verifier.Verify(raw)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrUnauthenticated, err)
	}

	p := &Principal{}
	if v, ok := lookupClaim(claims, a.subjectPath); ok {
		if p.Subject, err = parseStringClaim(v); err != nil {
			return nil, fmt.Errorf("%w: claim %s: %v", ErrUnauthenticated, strings.Join(a.subjectPath, "."), err)
		}
	}
	if v, ok := lookupClaim(claims, a.groupsPath); ok {
		if p.Groups, err = parseStringsClaim(v); err != nil {
			return nil, fmt.Errorf("%w: claim %s: %v", ErrUnauthenticated, strings.Join(a.groupsPath, "."), err)
		}
	}

	permissions, hasPermissions := lookupClaim(claims, a.permissionsPath)
	useClaims := a.source == SourceClaims || (a.source == SourceAuto && hasPermissions)

	switch {
	case useClaims && hasPermissions:
		if p.Rules, err = parseRulesClaim(permissions); err != nil {
			return nil, fmt.Errorf("%w: claim %s: %v", ErrUnauthenticated, strings.Join(a.permissionsPath, "."), err)
		}
	case useClaims:
		// claims-only mode and the token has no permissions: authenticated, but nothing is allowed.
	case p.Subject == "" && len(p.Groups) == 0:
		// Nothing to look up bindings by, skip the query.
	default:
		if p.Rules, err = a.store.Rules(ctx, p.Subject, p.Groups); err != nil {
			return nil, fmt.Errorf("load rules for %s: %w", p.Subject, err)
		}
	}
	return p, nil
}
