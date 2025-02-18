package myhub

import (
	"context"
	"fmt"

	builder "github.com/doug-martin/goqu/v9"
	"github.com/google/uuid"

	"github.com/zikwall/myhub/internal/application/domain"
	"github.com/zikwall/myhub/pkg/x"
)

type userDataTransferObject struct {
	ID        uuid.UUID `db:"id"`
	UserName  string    `db:"user_name"`
	FirstName string    `db:"first_name"`
	LastName  string    `db:"last_name"`
}

func (r *Repository) ListUsers(ctx context.Context) ([]domain.User, error) {
	var list []userDataTransferObject

	err := r.database.Builder().
		Select(
			builder.C("id"),
			builder.C("user_name"),
			builder.C("first_name"),
			builder.C("last_name"),
		).
		From("users").
		ScanStructsContext(ctx, &list)
	if err != nil {
		return nil, fmt.Errorf("ScanStructsContext: %w", err)
	}

	return x.Map(list, func(item userDataTransferObject, index int) domain.User {
		return domain.User{
			ID:        item.ID,
			UserName:  item.UserName,
			FirstName: item.FirstName,
			LastName:  item.LastName,
		}
	}), nil
}
