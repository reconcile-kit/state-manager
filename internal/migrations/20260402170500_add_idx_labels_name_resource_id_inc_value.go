package schema

import (
	"context"
	"database/sql"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationNoTxContext(upAddIdxLabelsNameResourceIDIncValue, downDropIdxLabelsNameResourceIDIncValue)
}

func upAddIdxLabelsNameResourceIDIncValue(ctx context.Context, db *sql.DB) error {
	if _, err := db.ExecContext(ctx, `
		CREATE INDEX CONCURRENTLY IF NOT EXISTS idx_labels_name_resource_id_inc_value
		ON labels (name, resource_id) INCLUDE (value);
	`); err != nil {
		return err
	}

	if _, err := db.ExecContext(ctx, `
		CREATE INDEX CONCURRENTLY IF NOT EXISTS idx_labels_name_value_resource_id 
		ON labels USING btree (name, value, resource_id);
	`); err != nil {
		return err
	}

	if _, err := db.ExecContext(ctx, `
		DROP INDEX CONCURRENTLY IF EXISTS idx_labels_name_value;
	`); err != nil {
		return err
	}

	_, err := db.ExecContext(ctx, `ANALYZE labels;`)
	return err
}

func downDropIdxLabelsNameResourceIDIncValue(ctx context.Context, db *sql.DB) error {
	_, err := db.ExecContext(ctx, `
		DROP INDEX CONCURRENTLY IF EXISTS idx_labels_name_resource_id_inc_value;
	`)
	if err != nil {
		return err
	}

	if _, err := db.ExecContext(ctx, `
		DROP INDEX CONCURRENTLY IF EXISTS idx_labels_name_value_resource_id;
	`); err != nil {
		return err
	}

	if _, err := db.ExecContext(ctx, `
		CREATE INDEX CONCURRENTLY IF NOT EXISTS idx_labels_name_value 
			ON labels (name, value);
	`); err != nil {
		return err
	}

	_, err = db.ExecContext(ctx, `ANALYZE labels;`)
	return err
}
