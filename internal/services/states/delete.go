package states

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/reconcile-kit/state-manager/internal/dto"
	"time"
)

func (s *StateService) Delete(ctx context.Context, opts *dto.ResourceID) error {
	_, err := s.repo.TxWrap(ctx, func(tx pgx.Tx) (*dto.Resource, error) {
		err := s.repo.Lock(ctx, tx, opts)
		if err != nil {
			return nil, err
		}
		currentResource, err := s.repo.GetByResourceID(ctx, tx, opts)
		if err != nil {
			return nil, err
		}
		if len(currentResource.Finalizers) == 0 {
			err = s.repo.Delete(ctx, tx, int64(currentResource.ID))
			if err != nil {
				return nil, err
			}
			err = s.eventsRepo.Add(ctx, currentResource.ShardID, "delete", currentResource.ResourceID)
			if err != nil {
				return nil, err
			}
			return currentResource, nil
		}

		if currentResource.DeletionTimestamp == nil {
			timeNow := time.Now()
			currentResource.DeletionTimestamp = &timeNow

			err = s.repo.SetDeletionTimestamp(ctx, tx, opts, timeNow)
			if err != nil {
				return nil, err
			}
		}

		err = s.eventsRepo.Add(ctx, currentResource.ShardID, "update", currentResource.ResourceID)
		if err != nil {
			return nil, err
		}

		return currentResource, nil
	})
	return err
}
