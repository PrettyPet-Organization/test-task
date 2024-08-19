package model

import (
	"time"
)

// Order is ...
type Order struct {
	ID          int       `json:"id" db:"id"`
	UserID      int       `json:"user_id" db:"user_id"`
	StateID     int       `json:"state_id" db:"state_id"`
	Items       []int     `json:"items" db:"items"`
	TotalAmount float64   `json:"total_amount" db:"total_amount"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
}

// Item is ...
type Item struct {
	ID         int     `json:"id" db:"id"`
	ProductID  int     `json:"product_id" db:"product_id"`
	Quantity   int     `json:"quantity" db:"quantity"`
	TotalPrice float64 `json:"total_price" db:"total_price"`
}

// OrderState is ...
type OrderState struct {
	ID   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}

// NewOrder is ...
func NewOrder() Order {
	return Order{}
}

// // Validate проверяет правильность данных поставщика
// func (u *Order) Validate() error {
// 	if u.Name == "" {
// 		return errors.New("name is required")
// 	}
// 	if u.ProviderID == 0 {
// 		return errors.New("provider is required")
// 	}
// 	if u.Price < 0 {
// 		return errors.New("price must be positive or null")
// 	}
// 	if u.Quantity < 0 {
// 		return errors.New("quantity must be positive or null")
// 	}
// 	return nil
// }
