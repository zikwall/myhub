package myhub

import (
	"context"
	"fmt"

	builder "github.com/doug-martin/goqu/v9"
	"github.com/google/uuid"

	"github.com/zikwall/myhub/internal/application/domain"
	"github.com/zikwall/myhub/pkg/x"
)

type channelDataTransferObject struct {
	ID          uuid.UUID `db:"id"`
	Name        string    `db:"name"`
	Description string    `db:"description"`
}

func (r *Repository) ListChannels(ctx context.Context) ([]domain.Channel, error) {
	var list []channelDataTransferObject

	err := r.database.Builder().
		Select(
			builder.C("id"),
			builder.C("name"),
			builder.C("description"),
		).
		From("channels").
		ScanStructsContext(ctx, &list)
	if err != nil {
		return nil, fmt.Errorf("ScanStructsContext: %w", err)
	}

	return x.Map(list, func(item channelDataTransferObject, index int) domain.Channel {
		return domain.Channel{
			ID:          item.ID,
			Name:        item.Name,
			Description: item.Description,
		}
	}), nil
}
