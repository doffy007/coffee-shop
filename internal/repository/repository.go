package repository

import (
	"context"

	"github.com/doffy007/coffee-shop/config"
	"github.com/doffy007/coffee-shop/database"
	"github.com/jmoiron/sqlx"
)

type repository struct {
	ctx    context.Context
	config *config.Config
	db     *sqlx.DB
}

func NewRepository(ctx context.Context, config *config.Config) Repository {
	return repository{
		ctx:    context.Background(),
		config: config,
		db:     database.Mysql(),
	}
}
