package store

import (
	"context"

	"github.com/KozlovNikolai/test-task/internal/model"
)

// IProviderRepository is ...
type IProviderRepository interface {
	Create(context.Context, model.Provider) (int, error)
	GetAll(ctx context.Context) ([]model.Provider, error)
	Get(context.Context, int) (model.Provider, error)
	Update(context.Context, int) error
	Delete(context.Context, int) error
}
