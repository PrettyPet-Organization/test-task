package repository

import (
	"errors"
	"small_warehouse/internal/model"
	"sync"
)

var ErrNotExists = errors.New("Entity with specified key doesn't exists")

type GoodsStore struct {
	m  map[string]model.Good
	mu *sync.RWMutex
}

func NewGoods() *GoodsStore {
	return &GoodsStore{
		m:  make(map[string]model.Good),
		mu: &sync.RWMutex{},
	}
}

func (s *GoodsStore) Get(key string) (model.Good, error) {
	s.mu.RLock()
	value, exists := s.m[key]
	s.mu.RUnlock()

	if !exists {
		return value, ErrNotExists
	}
	return value, nil
}

func (s *GoodsStore) GetAll() []model.Good {
	m := make([]model.Good, len(s.m))
	for _, v := range s.m {
		m = append(m, v)
	}

	return m
}

func (s *GoodsStore) Put(value model.Good) {
	s.mu.Lock()
	s.m[value.GetKey()] = value
	s.mu.Unlock()
}

func (s *GoodsStore) Delete(key string) {
	s.mu.Lock()
	delete(s.m, key)
	s.mu.Unlock()
}

func (s *GoodsStore) Count() int {
	s.mu.RLock()
	count := len(s.m)
	s.mu.RUnlock()
	return count
}
