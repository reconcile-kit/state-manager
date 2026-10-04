package auth

import (
	"context"
	"errors"
	"fmt"
)

var (
	ErrUnauthenticated = errors.New("unauthenticated")
	ErrForbidden       = errors.New("forbidden")
)

// Principal is an authenticated caller with the rules granted to it.
type Principal struct {
	Subject string
	Groups  []string
	Rules   []Rule
}

type principalKey struct{}

func WithPrincipal(ctx context.Context, p *Principal) context.Context {
	return context.WithValue(ctx, principalKey{}, p)
}

func PrincipalFromContext(ctx context.Context) (*Principal, bool) {
	p, ok := ctx.Value(principalKey{}).(*Principal)
	return p, ok && p != nil
}

// Authorizer decides whether the caller from ctx may perform verb on attrs.
type Authorizer interface {
	Authorize(ctx context.Context, verb Verb, attrs *Attributes) error
}

// NoopAuthorizer allows everything. It is used when authorization is disabled.
type NoopAuthorizer struct{}

func (NoopAuthorizer) Authorize(context.Context, Verb, *Attributes) error {
	return nil
}

// RulesAuthorizer checks the rules of the principal stored in ctx.
// Rules are additive: access is granted if any rule allows it.
type RulesAuthorizer struct{}

func (RulesAuthorizer) Authorize(ctx context.Context, verb Verb, attrs *Attributes) error {
	p, ok := PrincipalFromContext(ctx)
	if !ok {
		return ErrUnauthenticated
	}
	for i := range p.Rules {
		if p.Rules[i].Allows(verb, attrs) {
			return nil
		}
	}
	// No quoting: handlers embed error text into a JSON string as is.
	return fmt.Errorf("%w: %s is not allowed to %s resource_group=%s namespace=%s kind=%s name=%s shard_id=%s",
		ErrForbidden, p.Subject, verb, attrs.ResourceGroup, attrs.Namespace, attrs.Kind, attrs.Name, attrs.ShardID)
}
