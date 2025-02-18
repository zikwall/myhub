// nolint:errcheck,stylecheck // because tests
package test_container

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"path"
	"testing"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/pressly/goose/v3"
	"github.com/stretchr/testify/require"

	"github.com/zikwall/myhub/pkg/database"
	"github.com/zikwall/myhub/pkg/database/postgres"
)

type TestDatabase struct {
	Pool database.Pool
}

func NewTestDatabase(t *testing.T) *TestDatabase {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	p, err := postgres.NewPoolSqlx(ctx, &database.Opt{
		Host:               "localhost",
		User:               "user",
		Password:           "pass",
		Port:               "5499",
		Name:               "my_hub",
		Dialect:            "postgres",
		Debug:              true,
		MaxIdleConns:       2,
		MaxOpenConns:       2,
		MaxConnMaxLifetime: time.Minute * 5,
	})
	require.NoError(t, err)

	test := &TestDatabase{Pool: p}

	err = test.up(ctx)
	require.NoError(t, err)

	return test
}

func (t *TestDatabase) Drop() error {
	if err := t.reset(context.Background()); err != nil {
		log.Printf("failed to reset database container: %s\n", err)
	}
	return nil
}

func (t *TestDatabase) up(ctx context.Context) error {
	wd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get current directory: %w", err)
	}
	dir := path.Join(wd, "../../../../../migrations")
	if err = goose.UpContext(ctx, t.db(), dir); err != nil {
		return fmt.Errorf("failed to run migrations: %w", err)
	}
	return nil
}

func (t *TestDatabase) reset(ctx context.Context) error {
	wd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get current directory: %w", err)
	}
	dir := path.Join(wd, "../../../../../migrations")

	if err = goose.ResetContext(ctx, t.db(), dir); err != nil {
		return fmt.Errorf("failed to reset migrations: %w", err)
	}
	return nil
}

func (t *TestDatabase) db() *sql.DB {
	mySqlxDb := t.Pool.Builder().Db.(*sqlx.DB)
	return mySqlxDb.DB
}
