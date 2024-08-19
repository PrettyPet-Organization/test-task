package pgstore

import (
	"context"

	"github.com/KozlovNikolai/test-task/internal/model"
	"github.com/jackc/pgx/v4/pgxpool"
	"go.uber.org/zap"
)

// UserRepository is ...
type UserRepository struct {
	logger *zap.Logger
	db     *pgxpool.Pool
}

// NewUserRepository is ...
func NewUserRepository(db *pgxpool.Pool, logger *zap.Logger) *UserRepository {
	logger = logger.Named("repo")
	return &UserRepository{
		logger: logger,
		db:     db,
	}
}

// Create implements store.IUserRepository.
func (p *UserRepository) Create(context.Context, model.User) (int, error) {
	panic("unimplemented")
}

// Delete implements store.IUserRepository.
func (p *UserRepository) Delete(context.Context, int) error {
	panic("unimplemented")
}

// GetByID implements store.IUserRepository.
func (p *UserRepository) GetByID(context.Context, int) (model.User, error) {
	panic("unimplemented")
}

// GetByLogin implements store.IUserRepository.
func (p *UserRepository) GetByLogin(context.Context, string) (model.User, error) {
	panic("unimplemented")
}

// GetAll implements store.IUserRepository.
func (p *UserRepository) GetAll(context.Context) ([]model.User, error) {
	panic("unimplemented")
}

// Update implements store.IUserRepository.
func (p *UserRepository) Update(context.Context, int) error {
	panic("unimplemented")
}
