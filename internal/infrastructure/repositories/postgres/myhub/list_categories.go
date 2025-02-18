package myhub

import (
	"context"
	"fmt"

	builder "github.com/doug-martin/goqu/v9"
	"github.com/google/uuid"

	"github.com/zikwall/myhub/internal/application/domain"
	"github.com/zikwall/myhub/pkg/x"
)

type categoryDataTransferObject struct {
	ID   uuid.UUID `db:"id"`
	Name string    `db:"name"`
}

func (r *Repository) ListCategories(ctx context.Context) ([]domain.Category, error) {
	var list []categoryDataTransferObject

	err := r.database.Builder().
		Select(
			builder.C("id"),
			builder.C("name"),
		).
		From("categories").
		ScanStructsContext(ctx, &list)
	if err != nil {
		return nil, fmt.Errorf("ScanStructsContext: %w", err)
	}

	return x.Map(list, func(item categoryDataTransferObject, index int) domain.Category {
		return domain.Category{
			ID:   item.ID,
			Name: item.Name,
		}
	}), nil
}
