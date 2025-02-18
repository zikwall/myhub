package myhub

import (
	"context"
	"fmt"

	"github.com/zikwall/myhub/internal/application/domain"
)

func (r *Repository) AddCategory(ctx context.Context, category domain.Category) error {
	_, err := r.database.Builder().
		Insert("public.categories").
		Rows(categoryDataTransferObject{
			ID:   category.ID,
			Name: category.Name,
		}).
		Executor().
		ExecContext(ctx)
	if err != nil {
		return fmt.Errorf("ExecContext: %w", err)
	}

	return nil
}
