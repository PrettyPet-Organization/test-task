package inmemory

import (
	"sync"

	"github.com/KozlovNikolai/test-task/internal/model"
	"go.uber.org/zap"
)

// Repository is ...
type Repository struct {
	logger          *zap.Logger
	users           map[int]model.User
	orders          map[int]model.Order
	products        map[int]model.Product
	providers       map[int]model.Provider
	nextUsersID     int
	nextOrdersID    int
	nextProductsID  int
	nextProvidersID int
	mutex           sync.Mutex
}

// NewRepository is ...
func NewRepository(logger *zap.Logger) *Repository {
	logger = logger.Named("repo")
	return &Repository{
		logger:          logger,
		users:           make(map[int]model.User),
		orders:          make(map[int]model.Order),
		products:        make(map[int]model.Product),
		providers:       make(map[int]model.Provider),
		nextUsersID:     1,
		nextOrdersID:    1,
		nextProductsID:  1,
		nextProvidersID: 1,
	}
}
