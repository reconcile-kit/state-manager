package states

import (
	"context"
	"github.com/dhnikolas/state-manager/internal/dto"
	"github.com/jackc/pgx/v5"
)

type ResourceRepository interface {
	TxWrap(ctx context.Context, fn func(tx pgx.Tx) (*dto.Resource, error)) (*dto.Resource, error)

	Create(ctx context.Context, tx pgx.Tx, opts *dto.ResourceCreateOpts) (*dto.Resource, error)
	Update(ctx context.Context, tx pgx.Tx, opts *dto.ResourceUpdateOpts) (*dto.Resource, error)
	UpdateStatus(ctx context.Context, tx pgx.Tx, opts *dto.ResourceUpdateStatusOpts) (*dto.Resource, error)
	Delete(ctx context.Context, tx pgx.Tx, id int64) error
	GetByResourceID(ctx context.Context, opts *dto.ResourceID) (*dto.Resource, error)
	ListResources(ctx context.Context, listOpts *dto.ListResourcesOpts) ([]*dto.Resource, error)
}

type StateService struct {
	repo ResourceRepository
}

func NewStateService(repo ResourceRepository) *StateService {
	return &StateService{repo: repo}
}
