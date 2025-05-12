package states

import (
	"context"
	"github.com/dhnikolas/state-manager/internal/dto"
	"github.com/jackc/pgx/v5"
)

func (s *StateService) Update(ctx context.Context, opts dto.ResourceUpdateOpts) (*dto.Resource, error) {
	return s.txWrapper(ctx, func(tx pgx.Tx) (*dto.Resource, error) {
		return s.repo.Update(ctx, tx, opts)
	})
}
