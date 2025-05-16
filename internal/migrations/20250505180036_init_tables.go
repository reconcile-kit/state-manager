package schema

import (
	"context"
	"database/sql"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upCreateResourcesTables, downDropResourcesTables)
}

func upCreateResourcesTables(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.Exec(`
		CREATE TABLE IF NOT EXISTS resources (
			id SERIAL PRIMARY KEY,
			shard_id VARCHAR(255) NOT NULL DEFAULT 'default',
			resource_group VARCHAR(255) NOT NULL,
			kind VARCHAR(255) NOT NULL,
			namespace VARCHAR(255) NOT NULL,
			name VARCHAR(255) NOT NULL,
			created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
			deletion_timestamp TIMESTAMP WITH TIME ZONE DEFAULT NULL,
			finalizers JSONB,    
			annotations JSONB,    
			spec JSONB,
			status JSONB,    
			version BIGINT NOT NULL DEFAULT 1,
			current_version BIGINT NOT NULL DEFAULT 0
		);

		CREATE UNIQUE INDEX IF NOT EXISTS idx_resources_resource_group_kind_namespace_name 
            ON resources (resource_group, kind, namespace, name);

		CREATE INDEX IF NOT EXISTS idx_resources_kind_version_current_version 
			ON resources (kind, version, current_version);

		CREATE TABLE IF NOT EXISTS labels (
			id SERIAL PRIMARY KEY,
			resource_id INTEGER NOT NULL,
			name VARCHAR(255) NOT NULL,
			value VARCHAR(255) NOT NULL,
			CONSTRAINT fk_resource
				FOREIGN KEY(resource_id) 
				REFERENCES resources(id)
				ON DELETE CASCADE
		);

		CREATE INDEX IF NOT EXISTS idx_labels_resource_id 
			ON labels (resource_id);

		CREATE INDEX IF NOT EXISTS idx_labels_name_value 
			ON labels (name, value);
	`)
	return err
}

func downDropResourcesTables(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.Exec(`
		DROP TABLE IF EXISTS labels;
		DROP TABLE IF EXISTS resources;
	`)
	return err
}
