package states

import (
	"context"
	"github.com/dhnikolas/state-manager/internal/dto"
	"github.com/jackc/pgx/v5"
	"time"
)

func (s *StateService) Delete(ctx context.Context, opts *dto.ResourceID) error {
	_, err := s.repo.TxWrap(ctx, func(tx pgx.Tx) (*dto.Resource, error) {
		currentResource, err := s.repo.GetByResourceID(ctx, opts)
		if err != nil {
			return nil, err
		}
		if len(currentResource.Finalizers) == 0 {
			err = s.repo.Delete(ctx, tx, int64(currentResource.ID))
			if err != nil {
				return nil, err
			}
			return currentResource, nil
		}
		if currentResource.DeletionTimestamp != nil {
			return currentResource, nil
		}
		timeNow := time.Now()
		currentResource.DeletionTimestamp = &timeNow

		return currentResource, s.repo.SetDeletionTimestamp(ctx, tx, opts, timeNow)
	})
	return err
}
