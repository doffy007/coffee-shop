package repository

import (
	"context"

	"github.com/doffy007/coffee-shop/config"
	"github.com/jmoiron/sqlx"
)

type UserRepository interface {
}

type userRepository struct {
	ctx    context.Context
	config *config.Config
	db     *sqlx.DB
}

// UserRepository implements Repository.
func (r repository) UserRepository() UserRepository {
	return &userRepository{
		ctx:    r.ctx,
		config: r.config,
		db:     r.db,
	}
}
