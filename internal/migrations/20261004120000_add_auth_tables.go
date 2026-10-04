package schema

import (
	"context"
	"database/sql"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upCreateAuthTables, downDropAuthTables)
}

func upCreateAuthTables(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.ExecContext(ctx, `
		CREATE TABLE IF NOT EXISTS auth_roles (
			id SERIAL PRIMARY KEY,
			name VARCHAR(255) NOT NULL UNIQUE,
			description TEXT NOT NULL DEFAULT '',
			created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
		);

		-- '*' in any field means "any value".
		CREATE TABLE IF NOT EXISTS auth_role_rules (
			id SERIAL PRIMARY KEY,
			role_id INTEGER NOT NULL REFERENCES auth_roles(id) ON DELETE CASCADE,
			verbs TEXT[] NOT NULL CHECK (cardinality(verbs) > 0),
			resource_group VARCHAR(255) NOT NULL DEFAULT '*',
			namespace VARCHAR(255) NOT NULL DEFAULT '*',
			kind VARCHAR(255) NOT NULL DEFAULT '*',
			name VARCHAR(255) NOT NULL DEFAULT '*',
			shard_id VARCHAR(255) NOT NULL DEFAULT '*'
		);

		CREATE INDEX IF NOT EXISTS idx_auth_role_rules_role_id
			ON auth_role_rules (role_id);

		-- subject_kind = 'subject': subject_value is the token subject claim (sub / client_id).
		-- subject_kind = 'group':   subject_value is one of the token groups.
		CREATE TABLE IF NOT EXISTS auth_role_bindings (
			id SERIAL PRIMARY KEY,
			role_id INTEGER NOT NULL REFERENCES auth_roles(id) ON DELETE CASCADE,
			subject_kind VARCHAR(16) NOT NULL CHECK (subject_kind IN ('subject', 'group')),
			subject_value VARCHAR(255) NOT NULL,
			disabled BOOLEAN NOT NULL DEFAULT FALSE,
			created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
			CONSTRAINT uq_auth_role_bindings UNIQUE (subject_kind, subject_value, role_id)
		);

		-- Lookup on every authenticated request: active bindings of a subject and its groups.
		CREATE INDEX IF NOT EXISTS idx_auth_role_bindings_active
			ON auth_role_bindings (subject_kind, subject_value) INCLUDE (role_id)
			WHERE NOT disabled;
	`)
	return err
}

func downDropAuthTables(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.ExecContext(ctx, `
		DROP TABLE IF EXISTS auth_role_bindings;
		DROP TABLE IF EXISTS auth_role_rules;
		DROP TABLE IF EXISTS auth_roles;
	`)
	return err
}
