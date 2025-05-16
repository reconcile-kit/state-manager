package states

import (
	"context"
	"errors"
	"github.com/dhnikolas/state-manager/internal/dto"
	"github.com/jackc/pgx/v5"
)

var ConflictError = errors.New("resource version not match")
var UnavailableVersion = errors.New("current version unavailable: current_version > version")

func (s *StateService) Update(ctx context.Context, opts *dto.ResourceUpdateOpts) (*dto.Resource, error) {
	return s.repo.TxWrap(ctx, func(tx pgx.Tx) (*dto.Resource, error) {
		currentResource, err := s.repo.GetByResourceID(ctx, &opts.ResourceID)
		if err != nil {
			return nil, err
		}
		if currentResource.DeletionTimestamp != nil && len(opts.Finalizers) == 0 {
			err = s.repo.Delete(ctx, tx, int64(currentResource.ID))
			if err != nil {
				return nil, err
			}
			return currentResource, nil
		}
		return s.repo.Update(ctx, tx, opts)
	})
}

func (s *StateService) UpdateStatus(ctx context.Context, opts *dto.ResourceUpdateStatusOpts) (*dto.Resource, error) {
	return s.repo.TxWrap(ctx, func(tx pgx.Tx) (*dto.Resource, error) {
		currentResource, err := s.repo.GetByResourceID(ctx, &opts.ResourceID)
		if err != nil {
			return nil, err
		}
		switch {
		case currentResource.Version != opts.Version:
			return nil, ConflictError
		case opts.CurrentVersion > currentResource.Version:
			return nil, UnavailableVersion
		}

		if currentResource.DeletionTimestamp != nil && len(opts.Finalizers) == 0 {
			err = s.repo.Delete(ctx, tx, int64(currentResource.ID))
			if err != nil {
				return nil, err
			}
			return currentResource, nil
		}

		return s.repo.UpdateStatus(ctx, tx, opts)
	})
}
