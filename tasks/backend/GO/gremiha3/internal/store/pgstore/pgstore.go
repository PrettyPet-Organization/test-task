package pgstore

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
	"go.uber.org/zap"
)

// Repository is ...
type Repository struct {
	logger *zap.Logger
	db     *pgxpool.Pool
}

// UpdateUser implements store.IRepository.
func (repo *Repository) UpdateUser(context.Context, int) error {
	panic("unimplemented")
}

// NewRepository is ...
func NewRepository(db *pgxpool.Pool, logger *zap.Logger) *Repository {
	logger = logger.Named("repo")
	return &Repository{
		logger: logger,
		db:     db,
	}
}
