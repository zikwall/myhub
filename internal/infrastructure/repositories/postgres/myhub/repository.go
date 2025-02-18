package myhub

import "github.com/zikwall/myhub/pkg/database"

type Repository struct {
	database database.Pool
}

func New(database database.Pool) *Repository {
	return &Repository{
		database: database,
	}
}
