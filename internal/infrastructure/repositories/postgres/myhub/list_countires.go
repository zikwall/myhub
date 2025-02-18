package myhub

import (
	"context"
	"fmt"

	builder "github.com/doug-martin/goqu/v9"
	"github.com/google/uuid"

	"github.com/zikwall/myhub/internal/application/domain"
	"github.com/zikwall/myhub/pkg/x"
)

type countryDataTransferObject struct {
	ID       uuid.UUID `db:"id"`
	Name     string    `db:"name"`
	Iso33661 string    `db:"iso3366_1"`
}

func (r *Repository) ListCountries(ctx context.Context) ([]domain.Country, error) {
	var list []countryDataTransferObject

	err := r.database.Builder().
		Select(
			builder.C("id"),
			builder.C("name"),
			builder.C("iso3366_1"),
		).
		From("countries").
		ScanStructsContext(ctx, &list)
	if err != nil {
		return nil, fmt.Errorf("ScanStructsContext: %w", err)
	}

	return x.Map(list, func(item countryDataTransferObject, index int) domain.Country {
		return domain.Country{
			ID:       item.ID,
			Name:     item.Name,
			Iso33661: item.Iso33661,
		}
	}), nil
}
