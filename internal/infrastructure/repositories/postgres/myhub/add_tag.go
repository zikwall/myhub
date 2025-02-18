package myhub

import (
	"context"
	"fmt"

	"github.com/zikwall/myhub/internal/application/domain"
)

func (r *Repository) AddTag(ctx context.Context, tag domain.Tag) error {
	_, err := r.database.Builder().
		Insert("public.tags").
		Rows(tagDataTransferObject{
			ID:    tag.ID,
			Label: tag.Label,
		}).
		Executor().
		ExecContext(ctx)
	if err != nil {
		return fmt.Errorf("ExecContext: %w", err)
	}

	return nil
}
