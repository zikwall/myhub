package myhub

import (
	"context"
	"fmt"

	"github.com/zikwall/myhub/internal/application/domain"
)

func (r *Repository) AddChannel(ctx context.Context, channel domain.Channel) error {
	_, err := r.database.Builder().
		Insert("public.channels").
		Rows(channelDataTransferObject{
			ID:          channel.ID,
			Name:        channel.Name,
			Description: channel.Description,
		}).
		Executor().
		ExecContext(ctx)
	if err != nil {
		return fmt.Errorf("ExecContext: %w", err)
	}

	return nil
}
