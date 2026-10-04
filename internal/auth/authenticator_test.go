package auth

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const testIssuer = "https://idp.example.com"

type fakeStore struct {
	calls int
	rules []Rule
	err   error
}

func (s *fakeStore) Rules(context.Context, string, []string) ([]Rule, error) {
	s.calls++
	return s.rules, s.err
}

type testEnv struct {
	key *rsa.PrivateKey
	cfg *Config
}

func newTestEnv(t testing.TB, source PermissionsSource) *testEnv {
	t.Helper()
	key, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		t.Fatal(err)
	}
	der, err := x509.MarshalPKIXPublicKey(&key.PublicKey)
	if err != nil {
		t.Fatal(err)
	}
	path := filepath.Join(t.TempDir(), "key.pem")
	if err = os.WriteFile(path, pem.EncodeToMemory(&pem.Block{Type: "PUBLIC KEY", Bytes: der}), 0o600); err != nil {
		t.Fatal(err)
	}
	cfg := &Config{
		Enabled:           true,
		PublicKeyFile:     path,
		Issuer:            testIssuer,
		Audience:          "state-manager",
		Algorithms:        []string{"RS256"},
		SubjectClaim:      "sub",
		GroupsClaim:       "groups",
		PermissionsClaim:  "permissions",
		PermissionsSource: source,
	}
	if err = cfg.Validate(); err != nil {
		t.Fatal(err)
	}
	return &testEnv{key: key, cfg: cfg}
}

func (e *testEnv) token(t testing.TB, extra jwt.MapClaims) string {
	t.Helper()
	claims := jwt.MapClaims{
		"iss": testIssuer,
		"aud": "state-manager",
		"sub": "svc-a",
		"exp": time.Now().Add(time.Hour).Unix(),
	}
	for k, v := range extra {
		if v == nil {
			delete(claims, k)
			continue
		}
		claims[k] = v
	}
	s, err := jwt.NewWithClaims(jwt.SigningMethodRS256, claims).SignedString(e.key)
	if err != nil {
		t.Fatal(err)
	}
	return s
}

func (e *testEnv) authenticator(t testing.TB, store RuleStore) *Authenticator {
	t.Helper()
	v, err := NewVerifier(context.Background(), e.cfg)
	if err != nil {
		t.Fatal(err)
	}
	a, err := NewAuthenticator(e.cfg, v, store)
	if err != nil {
		t.Fatal(err)
	}
	return a
}

var permissionsClaim = []any{
	map[string]any{"verbs": []any{"get", "list"}, "shard_id": "5"},
	map[string]any{"verbs": []any{"get"}, "kind": "user", "namespace": "*"},
}

func TestAuthenticateAutoUsesClaimWithoutDB(t *testing.T) {
	env := newTestEnv(t, SourceAuto)
	store := &fakeStore{}
	p, err := env.authenticator(t, store).Authenticate(context.Background(),
		env.token(t, jwt.MapClaims{"permissions": permissionsClaim, "groups": []any{"g1"}}))
	if err != nil {
		t.Fatal(err)
	}
	if store.calls != 0 {
		t.Fatalf("store must not be queried when the token has permissions, calls=%d", store.calls)
	}
	if p.Subject != "svc-a" || len(p.Groups) != 1 || len(p.Rules) != 2 {
		t.Fatalf("unexpected principal %+v", p)
	}
	if p.Rules[1].Namespace != "" || p.Rules[1].Kind != "user" {
		t.Fatalf("wildcard must be normalized: %+v", p.Rules[1])
	}
}

func TestAuthenticateAutoFallsBackToDB(t *testing.T) {
	env := newTestEnv(t, SourceAuto)
	store := &fakeStore{rules: []Rule{{Verbs: VerbAll}}}
	p, err := env.authenticator(t, store).Authenticate(context.Background(), env.token(t, nil))
	if err != nil {
		t.Fatal(err)
	}
	if store.calls != 1 || len(p.Rules) != 1 {
		t.Fatalf("expected rules from store, calls=%d rules=%v", store.calls, p.Rules)
	}
}

func TestAuthenticateDBIgnoresClaim(t *testing.T) {
	env := newTestEnv(t, SourceDB)
	store := &fakeStore{}
	p, err := env.authenticator(t, store).Authenticate(context.Background(),
		env.token(t, jwt.MapClaims{"permissions": permissionsClaim}))
	if err != nil {
		t.Fatal(err)
	}
	if store.calls != 1 || len(p.Rules) != 0 {
		t.Fatalf("db mode must use store only, calls=%d rules=%v", store.calls, p.Rules)
	}
}

func TestAuthenticateClaimsModeWithoutPermissions(t *testing.T) {
	env := newTestEnv(t, SourceClaims)
	p, err := env.authenticator(t, nil).Authenticate(context.Background(), env.token(t, nil))
	if err != nil {
		t.Fatal(err)
	}
	if len(p.Rules) != 0 {
		t.Fatalf("expected no rules, got %v", p.Rules)
	}
}

func TestAuthenticateNestedClaim(t *testing.T) {
	env := newTestEnv(t, SourceClaims)
	env.cfg.PermissionsClaim = "resource_access.state-manager.permissions"
	p, err := env.authenticator(t, nil).Authenticate(context.Background(), env.token(t, jwt.MapClaims{
		"resource_access": map[string]any{"state-manager": map[string]any{"permissions": permissionsClaim}},
	}))
	if err != nil {
		t.Fatal(err)
	}
	if len(p.Rules) != 2 {
		t.Fatalf("expected 2 rules, got %v", p.Rules)
	}
}

func TestAuthenticateSkipsDBWithoutIdentity(t *testing.T) {
	env := newTestEnv(t, SourceDB)
	store := &fakeStore{}
	if _, err := env.authenticator(t, store).Authenticate(context.Background(), env.token(t, jwt.MapClaims{"sub": nil})); err != nil {
		t.Fatal(err)
	}
	if store.calls != 0 {
		t.Fatalf("store must not be queried without subject and groups")
	}
}

func TestAuthenticateRejectsInvalidTokens(t *testing.T) {
	env := newTestEnv(t, SourceAuto)
	a := env.authenticator(t, &fakeStore{})
	other := newTestEnv(t, SourceAuto)

	hs, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss": testIssuer, "aud": "state-manager", "exp": time.Now().Add(time.Hour).Unix(),
	}).SignedString([]byte("secret"))
	if err != nil {
		t.Fatal(err)
	}

	tests := map[string]string{
		"expired":           env.token(t, jwt.MapClaims{"exp": time.Now().Add(-time.Hour).Unix()}),
		"no exp":            env.token(t, jwt.MapClaims{"exp": nil}),
		"wrong issuer":      env.token(t, jwt.MapClaims{"iss": "https://evil"}),
		"wrong audience":    env.token(t, jwt.MapClaims{"aud": "other"}),
		"foreign key":       other.token(t, nil),
		"hs256":             hs,
		"garbage":           "not-a-jwt",
		"bad permissions":   env.token(t, jwt.MapClaims{"permissions": "admin"}),
		"bad verb":          env.token(t, jwt.MapClaims{"permissions": []any{map[string]any{"verbs": []any{"patch"}}}}),
		"non-string groups": env.token(t, jwt.MapClaims{"groups": []any{1}}),
	}
	for name, token := range tests {
		t.Run(name, func(t *testing.T) {
			if _, err := a.Authenticate(context.Background(), token); !errors.Is(err, ErrUnauthenticated) {
				t.Fatalf("expected ErrUnauthenticated, got %v", err)
			}
		})
	}
}

func TestMiddleware(t *testing.T) {
	env := newTestEnv(t, SourceAuto)
	store := &fakeStore{}
	handler := Middleware(env.authenticator(t, store))(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		p, ok := PrincipalFromContext(r.Context())
		if !ok || p.Subject != "svc-a" {
			t.Errorf("principal not in context: %+v", p)
		}
		w.WriteHeader(http.StatusOK)
	}))

	do := func(header string) *httptest.ResponseRecorder {
		req := httptest.NewRequest(http.MethodGet, "/api/v1/resources", nil)
		if header != "" {
			req.Header.Set("Authorization", header)
		}
		rec := httptest.NewRecorder()
		handler.ServeHTTP(rec, req)
		return rec
	}

	if rec := do(""); rec.Code != http.StatusUnauthorized || rec.Header().Get("WWW-Authenticate") == "" {
		t.Fatalf("missing token: got %d", rec.Code)
	}
	if rec := do("Basic abc"); rec.Code != http.StatusUnauthorized {
		t.Fatalf("basic auth: got %d", rec.Code)
	}
	if rec := do("Bearer not-a-jwt"); rec.Code != http.StatusUnauthorized {
		t.Fatalf("invalid token: got %d", rec.Code)
	}
	if rec := do("bearer " + env.token(t, nil)); rec.Code != http.StatusOK {
		t.Fatalf("valid token: got %d %s", rec.Code, rec.Body)
	}

	store.err = errors.New("db down")
	if rec := do("Bearer " + env.token(t, nil)); rec.Code != http.StatusInternalServerError {
		t.Fatalf("store error: got %d", rec.Code)
	}
}

func TestConfigValidate(t *testing.T) {
	base := func() *Config {
		return &Config{
			Enabled: true, JWKSURL: "https://idp/jwks", Issuer: testIssuer, Algorithms: []string{"RS256"},
			SubjectClaim: "sub", PermissionsClaim: "permissions", PermissionsSource: SourceAuto,
		}
	}
	if err := base().Validate(); err != nil {
		t.Fatalf("valid config: %v", err)
	}
	if err := (&Config{}).Validate(); err != nil {
		t.Fatalf("disabled config must be valid: %v", err)
	}

	bad := map[string]func(c *Config){
		"two key sources": func(c *Config) { c.PublicKeyFile = "/key.pem" },
		"no key source":   func(c *Config) { c.JWKSURL = "" },
		"no issuer":       func(c *Config) { c.Issuer = "" },
		"hs256":           func(c *Config) { c.Algorithms = []string{"HS256"} },
		"none":            func(c *Config) { c.Algorithms = []string{"none"} },
		"unknown source":  func(c *Config) { c.PermissionsSource = "ldap" },
	}
	for name, mutate := range bad {
		t.Run(name, func(t *testing.T) {
			c := base()
			mutate(c)
			if err := c.Validate(); err == nil {
				t.Fatal("expected error")
			}
		})
	}
}

func BenchmarkAuthenticateClaims(b *testing.B) {
	env := newTestEnv(b, SourceAuto)
	a := env.authenticator(b, &fakeStore{})
	token := env.token(b, jwt.MapClaims{"permissions": permissionsClaim, "groups": []any{"g1", "g2"}})
	ctx := context.Background()
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, err := a.Authenticate(ctx, token); err != nil {
			b.Fatal(err)
		}
	}
}
