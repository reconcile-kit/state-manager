package states

import (
	"context"

	"github.com/reconcile-kit/state-manager/internal/auth"
	"github.com/reconcile-kit/state-manager/internal/dto"
)

func (s *StateService) authorize(ctx context.Context, verb auth.Verb, id *dto.ResourceID, shardID string) error {
	return s.authorizer.Authorize(ctx, verb, &auth.Attributes{
		ResourceGroup: id.ResourceGroup,
		Namespace:     id.Namespace,
		Kind:          id.Kind,
		Name:          id.Name,
		ShardID:       shardID,
	})
}

// authorizeChange checks the stored resource and, when the request moves it
// to another shard, the target shard as well.
func (s *StateService) authorizeChange(ctx context.Context, verb auth.Verb, current *dto.Resource, newShardID string) error {
	if err := s.authorize(ctx, verb, &current.ResourceID, current.ShardID); err != nil {
		return err
	}
	if newShardID != current.ShardID {
		return s.authorize(ctx, verb, &current.ResourceID, newShardID)
	}
	return nil
}
