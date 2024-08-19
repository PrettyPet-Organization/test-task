package store

import (
	"context"

	"github.com/KozlovNikolai/test-task/internal/model"
)

// IOrderRepository is ...
type IOrderRepository interface {
	Create(context.Context, model.Order) (int, error)
	GetAll(ctx context.Context) ([]model.Order, error)
	Get(context.Context, int) (model.Order, error)
	Update(context.Context, int) error
	Delete(context.Context, int) error
}
