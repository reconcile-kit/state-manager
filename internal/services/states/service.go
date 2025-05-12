package states

import (
	"context"
	"fmt"
	"github.com/dhnikolas/state-manager/internal/dto"
	"github.com/jackc/pgx/v5"
)

type ResourceRepository interface {
	BeginTx(ctx context.Context) (pgx.Tx, error)

	Create(ctx context.Context, tx pgx.Tx, opts dto.ResourceCreateOpts) (*dto.Resource, error)
	Update(ctx context.Context, tx pgx.Tx, opts dto.ResourceUpdateOpts) (*dto.Resource, error)
	Delete(ctx context.Context, tx pgx.Tx, id int64) error
	GetByResourceID(ctx context.Context, opts dto.ResourceID) (*dto.Resource, error)
	ListResources(ctx context.Context, listOpts dto.ListResourcesOpts) ([]*dto.Resource, error)
}

type StateService struct {
	repo ResourceRepository
}

func NewStateService(repo ResourceRepository) *StateService {
	return &StateService{repo: repo}
}

func (s *StateService) txWrapper(ctx context.Context, fn func(tx pgx.Tx) (*dto.Resource, error)) (*dto.Resource, error) {
	tx, err := s.repo.BeginTx(ctx)
	if err != nil {
		return nil, err
	}
	res, err := fn(tx)
	if err != nil {
		rollbackErr := tx.Rollback(ctx)
		if rollbackErr != nil {
			return nil, fmt.Errorf("rollback transaction error: %s", err)
		}
		return nil, err
	}
	if err = tx.Commit(ctx); err != nil {
		return nil, err
	}
	return res, nil
}
