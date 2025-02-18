package myhub

import (
	"context"
	"fmt"

	"github.com/zikwall/myhub/internal/application/domain"
)

func (r *Repository) AddUser(ctx context.Context, user domain.User) error {
	_, err := r.database.Builder().
		Insert("public.users").
		Rows(userDataTransferObject{
			ID:        user.ID,
			UserName:  user.UserName,
			FirstName: user.FirstName,
			LastName:  user.LastName,
		}).
		Executor().
		ExecContext(ctx)
	if err != nil {
		return fmt.Errorf("ExecContext: %w", err)
	}

	return nil
}
