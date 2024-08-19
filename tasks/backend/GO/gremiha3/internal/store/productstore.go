package store

import (
	"context"

	"github.com/KozlovNikolai/test-task/internal/model"
)

// IProductRepository is ...
type IProductRepository interface {
	Create(context.Context, model.Product) (int, error)
	GetAll(ctx context.Context) ([]model.Product, error)
	Get(context.Context, int) (model.Product, error)
	Update(context.Context, int) error
	Delete(context.Context, int) error
}
