package myhub

import (
	"context"
	"fmt"

	"github.com/zikwall/myhub/internal/application/domain"
)

func (r *Repository) AddCountry(ctx context.Context, country domain.Country) error {
	_, err := r.database.Builder().
		Insert("public.countries").
		Rows(countryDataTransferObject{
			ID:       country.ID,
			Iso33661: country.Iso33661,
			Name:     country.Name,
		}).
		Executor().
		ExecContext(ctx)
	if err != nil {
		return fmt.Errorf("ExecContext: %w", err)
	}

	return nil
}
