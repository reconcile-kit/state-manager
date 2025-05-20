package states

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/reconcile-kit/state-manager/internal/dto"
	"time"
)

type ResourceRepository interface {
	TxWrap(ctx context.Context, fn func(tx pgx.Tx) (*dto.Resource, error)) (*dto.Resource, error)

	Create(ctx context.Context, tx pgx.Tx, opts *dto.ResourceCreateOpts) (*dto.Resource, error)
	Update(ctx context.Context, tx pgx.Tx, opts *dto.ResourceUpdateOpts) (*dto.Resource, error)
	UpdateStatus(ctx context.Context, tx pgx.Tx, opts *dto.ResourceUpdateStatusOpts) (*dto.Resource, error)
	SetDeletionTimestamp(ctx context.Context, tx pgx.Tx, opts *dto.ResourceID, deletionTimestamp time.Time) error
	Delete(ctx context.Context, tx pgx.Tx, id int64) error
	GetByResourceID(ctx context.Context, opts *dto.ResourceID) (*dto.Resource, error)
	ListResources(ctx context.Context, listOpts *dto.ListResourcesOpts) ([]*dto.Resource, error)
}

type EventsRepository interface {
	Add(ctx context.Context, shardID, messageType string, resource dto.ResourceID) error
}

type StateService struct {
	repo       ResourceRepository
	eventsRepo EventsRepository
}

func NewStateService(repo ResourceRepository, eventsRepo EventsRepository) *StateService {
	return &StateService{repo: repo, eventsRepo: eventsRepo}
}
