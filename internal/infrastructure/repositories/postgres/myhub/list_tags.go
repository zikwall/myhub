package myhub

import (
	"context"
	"fmt"

	builder "github.com/doug-martin/goqu/v9"
	"github.com/google/uuid"

	"github.com/zikwall/myhub/internal/application/domain"
	"github.com/zikwall/myhub/pkg/x"
)

type tagDataTransferObject struct {
	ID    uuid.UUID `db:"id"`
	Label string    `db:"label"`
}

func (r *Repository) ListTags(ctx context.Context) ([]domain.Tag, error) {
	var list []tagDataTransferObject

	err := r.database.Builder().
		Select(
			builder.C("id"),
			builder.C("label"),
		).
		From("tags").
		ScanStructsContext(ctx, &list)
	if err != nil {
		return nil, fmt.Errorf("ScanStructsContext: %w", err)
	}

	return x.Map(list, func(item tagDataTransferObject, index int) domain.Tag {
		return domain.Tag{
			ID:    item.ID,
			Label: item.Label,
		}
	}), nil
}
