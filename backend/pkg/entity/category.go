package entity

import (
	"time"
)

// Represents a category
type MenuItemCategory struct {
	ID           string    `json:"id" db:"id"`
	Name         string    `json:"name" db:"name"`
	RestaurantID string    `json:"-" db:"restaurant_id"`
	CreatedAt    time.Time `json:"-"  db:"created_at"`
	UpdatedAt    time.Time `json:"-"  db:"updated_at"`
}
