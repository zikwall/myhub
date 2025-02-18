package myhub

import (
	"context"
	"fmt"

	"github.com/zikwall/myhub/internal/application/domain"
)

func (r *Repository) AddTimezone(ctx context.Context, timezone domain.Timezone) error {
	_, err := r.database.Builder().
		Insert("public.timezones").
		Rows(timezoneDataTransferObject{
			ID:      timezone.ID,
			Zone:    timezone.Zone,
			ZoneRFC: timezone.ZoneRFC,
			Name:    timezone.Name,
		}).
		Executor().
		ExecContext(ctx)
	if err != nil {
		return fmt.Errorf("ExecContext: %w", err)
	}

	return nil
}
