package states

import (
	"context"
	"github.com/dhnikolas/state-manager/internal/dto"
	"github.com/jackc/pgx/v5"
)

func (s *StateService) Create(ctx context.Context, opts dto.ResourceCreateOpts) (*dto.Resource, error) {
	return s.txWrapper(ctx, func(tx pgx.Tx) (*dto.Resource, error) {

		return s.repo.Create(ctx, tx, opts)
	})
}
