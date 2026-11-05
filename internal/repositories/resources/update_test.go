package resources

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/jackc/pgx/v5"
	"github.com/reconcile-kit/state-manager/internal/dto"
)

func TestUpdateMissingResource(t *testing.T) {
	t.Parallel()

	databaseErr := errors.New("database unavailable")
	tests := []struct {
		name    string
		scanErr error
		wantErr error
	}{
		{name: "no rows", scanErr: pgx.ErrNoRows, wantErr: dto.NotFoundError},
		{name: "wrapped no rows", scanErr: fmt.Errorf("scan: %w", pgx.ErrNoRows), wantErr: dto.NotFoundError},
		{name: "database failure", scanErr: databaseErr, wantErr: databaseErr},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			repo := NewResourceRepository(nil)
			tx := failingScanTx{scanErr: tt.scanErr}
			fields := dto.ResourceFields{
				ResourceID: dto.ResourceID{
					ResourceGroup: "group1",
					Namespace:     "ns1",
					Kind:          "type1",
					Name:          "resource1",
				},
			}

			t.Run("update", func(t *testing.T) {
				result, err := repo.Update(context.Background(), tx, &dto.ResourceUpdateOpts{ResourceFields: fields})

				if result != nil || !errors.Is(err, tt.wantErr) {
					t.Fatalf("Update = (%v, %v), want (nil, %v)", result, err, tt.wantErr)
				}
			})

			t.Run("update status", func(t *testing.T) {
				result, err := repo.UpdateStatus(context.Background(), tx, &dto.ResourceUpdateStatusOpts{ResourceFields: fields})

				if result != nil || !errors.Is(err, tt.wantErr) {
					t.Fatalf("UpdateStatus = (%v, %v), want (nil, %v)", result, err, tt.wantErr)
				}
			})
		})
	}
}

type failingScanTx struct {
	pgx.Tx
	scanErr error
}

func (tx failingScanTx) QueryRow(context.Context, string, ...any) pgx.Row {
	return failingRow{err: tx.scanErr}
}

type failingRow struct {
	err error
}

func (r failingRow) Scan(...any) error {
	return r.err
}
