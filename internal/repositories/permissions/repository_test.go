package permissions

import (
	"context"
	"database/sql"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
	"github.com/reconcile-kit/state-manager/internal/auth"
	_ "github.com/reconcile-kit/state-manager/internal/migrations"
)

// Integration test, runs only with TEST_DATABASE_URL pointing to an empty Postgres.
func TestRules(t *testing.T) {
	dbURL := os.Getenv("TEST_DATABASE_URL")
	if dbURL == "" {
		t.Skip("TEST_DATABASE_URL is not set")
	}
	ctx := context.Background()

	db, err := sql.Open("pgx", dbURL)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	if err = goose.SetDialect("postgres"); err != nil {
		t.Fatal(err)
	}
	if err = goose.Up(db, t.TempDir()); err != nil {
		t.Fatal(err)
	}

	_, err = db.ExecContext(ctx, `
		TRUNCATE auth_roles CASCADE;
		INSERT INTO auth_roles (id, name) VALUES (1, 'shard-5-reader'), (2, 'user-reader'), (3, 'admin');
		INSERT INTO auth_role_rules (role_id, verbs, shard_id) VALUES (1, '{get,list}', '5');
		INSERT INTO auth_role_rules (role_id, verbs, kind) VALUES (2, '{get,list}', 'user');
		INSERT INTO auth_role_rules (role_id, verbs) VALUES (3, '{*}');
		INSERT INTO auth_role_bindings (role_id, subject_kind, subject_value) VALUES
			(1, 'subject', 'svc-a'),
			(2, 'group', 'readers');
		INSERT INTO auth_role_bindings (role_id, subject_kind, subject_value, disabled) VALUES
			(3, 'subject', 'svc-a', true);
	`)
	if err != nil {
		t.Fatal(err)
	}

	pool, err := pgxpool.New(ctx, dbURL)
	if err != nil {
		t.Fatal(err)
	}
	defer pool.Close()
	repo := NewPermissionsRepository(pool)

	rules, err := repo.Rules(ctx, "svc-a", []string{"readers", "other"})
	if err != nil {
		t.Fatal(err)
	}
	if len(rules) != 2 {
		t.Fatalf("expected 2 rules (disabled binding excluded), got %+v", rules)
	}
	for _, r := range rules {
		if r.Verbs == auth.VerbAll {
			t.Fatalf("disabled admin binding leaked: %+v", r)
		}
	}

	rules, err = repo.Rules(ctx, "svc-a", nil)
	if err != nil {
		t.Fatal(err)
	}
	if len(rules) != 1 || rules[0].ShardID != "5" || rules[0].Kind != "" {
		t.Fatalf("expected shard-5 rule with wildcard kind, got %+v", rules)
	}

	rules, err = repo.Rules(ctx, "unknown", nil)
	if err != nil || len(rules) != 0 {
		t.Fatalf("expected no rules, got %+v, %v", rules, err)
	}
}
