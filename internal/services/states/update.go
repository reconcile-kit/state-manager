package states

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/reconcile-kit/state-manager/internal/dto"
)

func (s *StateService) Update(ctx context.Context, opts *dto.ResourceUpdateOpts) (*dto.Resource, error) {
	result, err := s.repo.TxWrap(ctx, func(tx pgx.Tx) (*dto.Resource, error) {

		err := s.repo.Lock(ctx, tx, &opts.ResourceID)
		if err != nil {
			return nil, err
		}
		currentResource, err := s.repo.GetByResourceID(ctx, tx, &opts.ResourceID)
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
		result, err := s.repo.Update(ctx, tx, opts)
		if err != nil {
			return nil, err
		}

		toUpdate, toDelete := DiffLabels(currentResource.Labels, opts.Labels)
		err = s.repo.UpdateLabels(ctx, tx, currentResource.ID, toUpdate, toDelete)
		if err != nil {
			return nil, fmt.Errorf("failed to update labels: %w", err)
		}

		err = s.eventsRepo.Add(ctx, result.ShardID, "update", result.ResourceID)
		if err != nil {
			return nil, err
		}

		return result, nil
	})
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *StateService) UpdateStatus(ctx context.Context, opts *dto.ResourceUpdateStatusOpts) (*dto.Resource, error) {
	return s.repo.TxWrap(ctx, func(tx pgx.Tx) (*dto.Resource, error) {
		err := s.repo.Lock(ctx, tx, &opts.ResourceID)
		if err != nil {
			return nil, err
		}
		currentResource, err := s.repo.GetByResourceID(ctx, tx, &opts.ResourceID)
		if err != nil {
			return nil, err
		}
		switch {
		case currentResource.Version != opts.Version:
			return nil, dto.ConflictError
		case opts.CurrentVersion > currentResource.Version:
			return nil, dto.UnavailableVersionError
		}

		if currentResource.DeletionTimestamp != nil && len(opts.Finalizers) == 0 {
			err = s.repo.Delete(ctx, tx, int64(currentResource.ID))
			if err != nil {
				return nil, err
			}
			return currentResource, nil
		}

		result, err := s.repo.UpdateStatus(ctx, tx, opts)
		if err != nil {
			return nil, err
		}

		toUpdate, toDelete := DiffLabels(currentResource.Labels, opts.Labels)
		err = s.repo.UpdateLabels(ctx, tx, currentResource.ID, toUpdate, toDelete)
		if err != nil {
			return nil, fmt.Errorf("failed to update labels: %w", err)
		}

		return result, nil
	})
}
