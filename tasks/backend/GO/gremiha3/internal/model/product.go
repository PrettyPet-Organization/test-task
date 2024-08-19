package model

import (
	"errors"
)

// Product is ...
type Product struct {
	ID         int     `json:"id" db:"id"`
	Name       string  `json:"name" db:"name"`
	ProviderID int     `json:"provider_id" db:"provider_id"`
	Price      float64 `json:"price" db:"price"`
	Stock      int     `json:"stock" db:"stock"`
}

// NewProduct is ...
func NewProduct() Product {
	return Product{}
}

// Validate проверяет правильность данных поставщика
func (u *Product) Validate() error {
	if u.Name == "" {
		return errors.New("name is required")
	}
	if u.ProviderID == 0 {
		return errors.New("provider is required")
	}
	if u.Price < 0 {
		return errors.New("price must be positive or null")
	}
	if u.Stock < 0 {
		return errors.New("quantity must be positive or null")
	}
	return nil
}
