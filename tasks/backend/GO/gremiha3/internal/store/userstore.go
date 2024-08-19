package store

import (
	"context"

	"github.com/KozlovNikolai/test-task/internal/model"
)

// IUserRepository is ...
type IUserRepository interface {
	Create(context.Context, model.User) (int, error)
	GetAll(ctx context.Context) ([]model.User, error)
	GetByID(context.Context, int) (model.User, error)
	GetByLogin(context.Context, string) (model.User, error)
	Update(context.Context, int) error
	Delete(context.Context, int) error
}
