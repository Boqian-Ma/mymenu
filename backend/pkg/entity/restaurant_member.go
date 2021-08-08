package entity

// Represents a relationship between a manager and a restaurant
type RestaurantMember struct {
	UserID       string `json:"-" db:"user_id"`
	RestaurantID string `json:"-" db:"restaurant_id"`
}
