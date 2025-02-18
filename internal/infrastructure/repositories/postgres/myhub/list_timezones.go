package myhub

import (
	"context"
	"fmt"

	builder "github.com/doug-martin/goqu/v9"
	"github.com/google/uuid"

	"github.com/zikwall/myhub/internal/application/domain"
	"github.com/zikwall/myhub/pkg/x"
)

type timezoneDataTransferObject struct {
	ID      uuid.UUID `db:"id"`
	Zone    int16     `db:"zone"`
	ZoneRFC string    `db:"zone_rfc"`
	Name    string    `db:"name"`
}

func (r *Repository) ListTimezones(ctx context.Context) ([]domain.Timezone, error) {
	var list []timezoneDataTransferObject

	err := r.database.Builder().
		Select(
			builder.C("id"),
			builder.C("zone"),
			builder.C("zone_rfc"),
			builder.C("name"),
		).
		From("timezones").
		ScanStructsContext(ctx, &list)
	if err != nil {
		return nil, fmt.Errorf("ScanStructsContext: %w", err)
	}

	return x.Map(list, func(item timezoneDataTransferObject, index int) domain.Timezone {
		return domain.Timezone{
			ID:      item.ID,
			Zone:    item.Zone,
			ZoneRFC: item.ZoneRFC,
			Name:    item.Name,
		}
	}), nil
}
