package service

import (
	"small_warehouse/internal/model"
)

type GoodsRepo interface {
	Get(string) (model.Good, error)
	GetAll() []model.Good
	Put(model.Good)
	Delete(key string)
	Count() int
}

type GoodsService struct {
	repo GoodsRepo
}

func NewGoods(repo GoodsRepo) *GoodsService {
	return &GoodsService{
		repo: repo,
	}
}

func (s *GoodsService) GetGoods() []model.Good {
	return s.repo.GetAll()
}

func (s *GoodsService) GetGood(key string) (model.Good, error) {
	return s.repo.Get(key)
}

func (s *GoodsService) CreateOrUpdateGood(t model.Good) error {
	s.repo.Put(t)
	return nil
}

func (s *GoodsService) CountGoods() int {
	return s.repo.Count()
}
