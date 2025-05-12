package states

import (
	"context"
	"github.com/dhnikolas/state-manager/internal/dto"
	"github.com/jackc/pgx/v5"
)

func (s *StateService) Delete(ctx context.Context, id int64) error {
	_, err := s.txWrapper(ctx, func(tx pgx.Tx) (*dto.Resource, error) {
		return nil, s.repo.Delete(ctx, tx, id)
	})
	return err
}
