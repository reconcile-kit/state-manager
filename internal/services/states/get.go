package states

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/reconcile-kit/state-manager/internal/auth"
	"github.com/reconcile-kit/state-manager/internal/dto"
)

func (s *StateService) GetByResourceID(ctx context.Context, opts *dto.ResourceID) (*dto.Resource, error) {
	result, err := s.repo.TxWrap(ctx, func(tx pgx.Tx) (*dto.Resource, error) {
		err := s.repo.Lock(ctx, tx, opts)
		if err != nil {
			return nil, err
		}
		resource, err := s.repo.GetByResourceID(ctx, tx, opts)
		if err != nil {
			return nil, err
		}
		// shard_id is known only after loading the resource.
		if err = s.authorize(ctx, auth.VerbGet, &resource.ResourceID, resource.ShardID); err != nil {
			return nil, err
		}
		return resource, nil
	})
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ListResources requires the whole filter to be covered by a single rule,
// so the result is never post-filtered and pagination stays correct.
func (s *StateService) ListResources(ctx context.Context, opts *dto.ListResourcesOpts) ([]*dto.Resource, error) {
	if err := s.authorize(ctx, auth.VerbList, &opts.ResourceID, opts.ShardID); err != nil {
		return nil, err
	}
	return s.repo.ListResources(ctx, opts)
}
