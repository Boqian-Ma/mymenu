package entity

import "time"

// Represents all possible user account types
type AccountType string

const (
	Manager  AccountType = "manager"
	Customer AccountType = "customer"
)

// Represents a user internally
type User struct {
	ID          string      `json:"id" db:"id"            example:"usr_abc123d"`
	CreatedAt   time.Time   `json:"-"  db:"created_at"`
	UpdatedAt   time.Time   `json:"-"  db:"updated_at"`
	AccountType AccountType `json:"accountType" db:"account_type"`
	Details     UserDetails `json:"details"`
	Login       UserLogin   `json:"-"`
}

type UserDetails struct {
	Email   *string `json:"email" db:"email"`
	Name    string  `json:"name" db:"name"`
	Allergy string  `json:"allergy" db:"allergy"`
}

type UserLogin struct {
	HashedPassword string `json:"-" db:"hashed_password"`
	Salt           string `json:"-" db:"salt"`
}
