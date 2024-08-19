package model

import (
	"errors"
)

// Provider is ...
type Provider struct {
	ID     int    `json:"id" db:"id"`
	Name   string `json:"name" db:"name"`
	Origin string `json:"origin" db:"origin"`
}

// NewProvider is ...
func NewProvider() Provider {
	return Provider{}
}

// Validate проверяет правильность данных поставщика
func (u *Provider) Validate() error {
	if u.Name == "" {
		return errors.New("name is required")
	}
	if u.Origin == "" {
		return errors.New("origin is required")
	}
	return nil
}
