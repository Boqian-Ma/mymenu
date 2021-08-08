package entity

import "time"

type StatusType string

const (
	Taken StatusType = "Taken"
	Free  StatusType = "Free"
)

// Represents a table internally
type Table struct {
	TableNum  int64     `json:"table_num" db:"table_num"`
	CreatedAt time.Time `json:"-"  db:"created_at"`
	UpdatedAt time.Time `json:"-"  db:"updated_at"`

	RestaurantID string     `json:"restaurant_id" db:"restaurant_id"`
	Status       StatusType `json:"status" db:"status"`
	NumSeats     int64      `json:"num_seats" db:"num_seats"`
}

func ValidStatus() []StatusType {
	return []StatusType{
		Taken,
		Free,
	}
}
