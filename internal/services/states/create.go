package states

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/reconcile-kit/state-manager/internal/dto"
)

func (s *StateService) Create(ctx context.Context, opts *dto.ResourceCreateOpts) (*dto.Resource, error) {
	return s.repo.TxWrap(ctx, func(tx pgx.Tx) (*dto.Resource, error) {
		err := s.repo.Lock(ctx, tx, &opts.ResourceID)
		if err != nil {
			return nil, err
		}
		result, err := s.repo.Create(ctx, tx, opts)
		if err != nil {
			return nil, err
		}

		err = s.eventsRepo.Add(ctx, result.ShardID, "update", result.ResourceID)
		if err != nil {
			return nil, err
		}

		return result, nil
	})
}
