package entity

import "time"

// Represents a restaurant internally
type Restaurant struct {
	ID        string    `json:"id" db:"id"            example:"res_abc123d"`
	CreatedAt time.Time `json:"-"  db:"created_at"`
	UpdatedAt time.Time `json:"-"  db:"updated_at"`

	Name          string `json:"name" db:"name"`
	Type          string `json:"type" db:"type"`
	Cuisine       string `json:"cuisine" db:"cuisine"`
	Location      string `json:"location" db:"location"`
	Email         string `json:"email" db:"email"`
	Phone         string `json:"phone" db:"phone"`
	Website       string `json:"website" db:"website"`
	BusinessHours string `json:"businessHours" db:"business_hours"`
	File          string `json:"file" db:"file"`
}

type Cuisine struct {
	Cuisine string `db:"cuisine"`
}
