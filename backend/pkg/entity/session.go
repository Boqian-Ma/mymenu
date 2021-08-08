package entity

type Session struct {
	Token  string `json:"-" db:"token"`
	UserID string `json:"-" db:"user_id"`
}
