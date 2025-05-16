package states

import (
	"context"
	"github.com/dhnikolas/state-manager/internal/dto"
)

func (s *StateService) GetByKey(ctx context.Context, opts *dto.ResourceID) (*dto.Resource, error) {
	return s.repo.GetByResourceID(ctx, opts)
}

func (s *StateService) ListPending(ctx context.Context, opts *dto.ListResourcesOpts) ([]*dto.Resource, error) {
	return s.repo.ListResources(ctx, opts)
}
