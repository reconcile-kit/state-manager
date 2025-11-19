package schema

import (
	"context"
	"database/sql"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upAddUniqueIndexLabels, downUniqueIndexLabels)
}

func upAddUniqueIndexLabels(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.Exec(`
		ALTER TABLE labels
		ADD CONSTRAINT labels_resource_id_name_key
		UNIQUE (resource_id, name);
	`)
	return err
}

func downUniqueIndexLabels(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.Exec(`
		ALTER TABLE labels
		DROP CONSTRAINT IF EXISTS labels_resource_id_name_key;
	`)
	return err
}
