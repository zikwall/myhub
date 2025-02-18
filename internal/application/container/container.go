package container

import (
	"context"

	"github.com/zikwall/myhub/config"
	"github.com/zikwall/myhub/pkg/database"
	"github.com/zikwall/myhub/pkg/database/postgres"
	"github.com/zikwall/myhub/pkg/drop"
)

type Container struct {
	*drop.Impl

	Pool database.Pool

	configuration *config.Config
}

func New(ctx context.Context, configuration *config.Config) (*Container, error) {
	c := &Container{
		Impl:          drop.NewContext(ctx),
		configuration: configuration,
	}

	var err error

	if configuration.Database.Host != "" {
		c.Pool, err = postgres.NewPoolSqlx(c.Context(), &configuration.Database)
		if err != nil {
			return nil, err
		}
		c.AddDropper(c.Pool.(*postgres.Pool))
	}

	return c, nil
}
