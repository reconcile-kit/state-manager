package states

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/reconcile-kit/state-manager/internal/dto"
)

func (s *StateService) GetByResourceID(ctx context.Context, opts *dto.ResourceID) (*dto.Resource, error) {
	result, err := s.repo.TxWrap(ctx, func(tx pgx.Tx) (*dto.Resource, error) {
		err := s.repo.Lock(ctx, tx, opts)
		if err != nil {
			return nil, err
		}
		return s.repo.GetByResourceID(ctx, tx, opts)
	})
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *StateService) ListPending(ctx context.Context, opts *dto.ListResourcesOpts) ([]*dto.Resource, error) {
	return s.repo.ListResources(ctx, opts)
}
