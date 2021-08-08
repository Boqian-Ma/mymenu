package entity

import (
	"time"
)

// Represents an menu item
type MenuItem struct {
	ID           string    `json:"id" db:"id"`
	Name         string    `json:"name" db:"name"`
	Description  string    `json:"description" db:"description"`
	Price        float64   `json:"price" db:"price"`
	IsSpecial    bool      `json:"is_special" db:"is_special"`
	IsMenu       bool      `json:"is_menu" db:"is_menu"`
	Allergy      string    `json:"allergy" db:"allergy"`
	CategoryID   string    `json:"category_id" db:"category_id"`
	RestaurantID string    `json:"-" db:"restaurant_id"`
	CreatedAt    time.Time `json:"-"  db:"created_at"`
	UpdatedAt    time.Time `json:"-"  db:"updated_at"`
	File         string    `json:"file"  db:"file"`
}

type MenuItemResult struct {
	ID           string    `json:"id" db:"id"`
	Name         string    `json:"name" db:"name"`
	Description  string    `json:"description" db:"description"`
	Price        float64   `json:"price" db:"price"`
	IsSpecial    bool      `json:"is_special" db:"is_special"`
	IsMenu       bool      `json:"is_menu" db:"is_menu"`
	Allergy      string    `json:"allergy" db:"allergy"`
	CategoryID   string    `json:"category_id" db:"category_id"`
	CategoryName string    `json:"category_name" db:"category_name"`
	RestaurantID string    `json:"-" db:"restaurant_id"`
	CreatedAt    time.Time `json:"-"  db:"created_at"`
	UpdatedAt    time.Time `json:"-"  db:"updated_at"`
	File         string    `json:"file"  db:"file"`
}
