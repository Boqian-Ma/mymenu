package entity

import "time"

type OrderStatus string

const (
	Ordered   OrderStatus = "ordered"
	Served    OrderStatus = "served"
	Completed OrderStatus = "completed"
	Canceled  OrderStatus = "canceled"
)

// Represents an order internally
type Order struct {
	ID           string       `json:"id" db:"id"`
	RestaurantID string       `json:"restaurant_id" db:"restaurant_id"`
	UserID       string       `json:"user_id" db:"user_id"`
	TotalCost    float64      `json:"total_cost" db:"total_cost"`
	Status       OrderStatus  `json:"status" db:"status"`
	CreatedAt    time.Time    `json:"-"  db:"created_at"`
	UpdatedAt    time.Time    `json:"-"  db:"updated_at"`
	TableNumber  int64        `json:"table_num" db:"table_num"`
	OrderItems   []*OrderItem `json:"items"`
}

type OrderResult struct {
	ID           string             `json:"id" db:"id"`
	RestaurantID string             `json:"restaurant_id" db:"restaurant_id"`
	UserID       string             `json:"user_id" db:"user_id"`
	TotalCost    float64            `json:"total_cost" db:"total_cost"`
	Status       OrderStatus        `json:"status" db:"status"`
	CreatedAt    time.Time          `json:"created_at"  db:"created_at"`
	UpdatedAt    time.Time          `json:"-"  db:"updated_at"`
	TableNumber  int64              `json:"table_num" db:"table_num"`
	OrderItems   []*OrderItemResult `json:"items"`
}

func GetOrderStatus() []OrderStatus {
	return []OrderStatus{
		Ordered,
		Served,
		Completed,
		Canceled,
	}
}
