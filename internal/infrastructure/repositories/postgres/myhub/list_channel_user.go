package myhub

import (
	"context"
	"fmt"

	builder "github.com/doug-martin/goqu/v9"
	"github.com/google/uuid"

	"github.com/zikwall/myhub/internal/application/domain"
	"github.com/zikwall/myhub/pkg/x"
)

type channelUserDataTransferObject struct {
	ID         uuid.UUID   `db:"id"`
	ChannelID  uuid.UUID   `db:"channel_id"`
	ZoneID     uuid.UUID   `db:"zone_id"`
	CategoryID uuid.UUID   `db:"category_id"`
	UserID     uuid.UUID   `db:"user_id"`
	Name       string      `db:"name"`
	URL        string      `db:"url"`
	Countries  []uuid.UUID `db:"countries"`
	Tags       []uuid.UUID `db:"tags"`
}

func (r *Repository) ListChannelUsers(ctx context.Context) ([]domain.ChannelUser, error) {
	var list []channelUserDataTransferObject

	err := r.database.Builder().
		Select(
			builder.C("id"),
			builder.C("channel_id"),
			builder.C("zone_id"),
			builder.C("category_id"),
			builder.C("user_id"),
			builder.C("name"),
			builder.C("url"),
			builder.C("countries"),
			builder.C("tags"),
		).
		From("channel_user").
		ScanStructsContext(ctx, &list)
	if err != nil {
		return nil, fmt.Errorf("ScanStructsContext: %w", err)
	}

	return x.Map(list, func(item channelUserDataTransferObject, index int) domain.ChannelUser {
		return domain.ChannelUser{
			ID:         item.ID,
			ChannelID:  item.ChannelID,
			ZoneID:     item.ZoneID,
			CategoryID: item.CategoryID,
			UserID:     item.UserID,
			Name:       item.Name,
			URL:        item.URL,
			Countries:  item.Countries,
			Tags:       item.Tags,
		}
	}), nil
}
