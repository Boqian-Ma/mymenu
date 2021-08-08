package utils

import (
	"time"
)

// String returns a pointer to the given string
func String(s string) *string {
	return &s
}

// Float64 returns a pointer to the given float64 value
func Float64(v float64) *float64 {
	return &v
}

// Bool returns a pointer to the given boolean value
func Bool(b bool) *bool {
	return &b
}

// Int returns a pointer to the given integer value
func Int(i int) *int {
	return &i
}

// Time returns a pointer to the given time value
func Time(t time.Time) *time.Time {
	return &t
}
