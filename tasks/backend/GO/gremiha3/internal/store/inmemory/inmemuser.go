package inmemory

import (
	"context"
	"errors"
	"sync"

	"github.com/KozlovNikolai/test-task/internal/model"
	"go.uber.org/zap"
)

// UserRepository is ...
type UserRepository struct {
	logger *zap.Logger
	users  map[int]model.User
	nextID int
	mutex  sync.Mutex
}

// NewUserRepository is ...
func NewUserRepository(logger *zap.Logger) *UserRepository {
	logger = logger.Named("repo")
	return &UserRepository{
		logger: logger,
		users:  make(map[int]model.User),
		nextID: 1,
	}
}

// Create implements store.IUserRepository.
func (i *UserRepository) Create(ctx context.Context, user model.User) (int, error) {

	_, err := i.GetByLogin(ctx, user.Login)
	if err == nil {
		return 0, errors.New("the login already exists")
	}
	i.mutex.Lock()
	defer i.mutex.Unlock()
	// i.logger.Info("INFO DEBUG", zap.String("Create Inmemory", "Come into handle"))

	user.ID = i.nextID
	i.nextID++
	i.users[user.ID] = user
	// msg := fmt.Sprintf("Created user : %v\n", user)
	// i.logger.Debug("CreateUser",
	// 	zap.String("info", msg),
	// )
	return user.ID, nil
}

// Delete implements store.IUserRepository.
func (i *UserRepository) Delete(context.Context, int) error {
	panic("unimplemented")
}

// GetAll implements store.IUserRepository.
func (i *UserRepository) GetAll(_ context.Context) ([]model.User, error) {
	i.mutex.Lock()
	defer i.mutex.Unlock()
	var users []model.User
	for _, user := range i.users {
		users = append(users, user)
	}
	return users, nil
}

// GetByID implements store.IUserRepository.
func (i *UserRepository) GetByID(_ context.Context, userID int) (model.User, error) {
	i.mutex.Lock()
	defer i.mutex.Unlock()
	user, exists := i.users[userID]
	if !exists {
		return model.User{}, errors.New("user not found")
	}
	return user, nil
}

// GetByLogin implements store.IUserRepository.
func (i *UserRepository) GetByLogin(_ context.Context, login string) (model.User, error) {
	i.mutex.Lock()
	defer i.mutex.Unlock()
	for _, user := range i.users {
		if user.Login == login {
			return user, nil
		}
	}
	return model.User{}, errors.New("login not found")
}

// Update implements store.IUserRepository.
func (i *UserRepository) Update(context.Context, int) error {
	panic("unimplemented")
}
