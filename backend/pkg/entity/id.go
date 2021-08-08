package entity

import (
	"math/rand"

	"github.com/COMP3900-9900-Capstone-Project/capstoneproject-comp3900-w16a-jamar/pkg/constants"
)

const idValues = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// GenerateUserID generates a new id for a user row
func GenerateUserID() string {
	return GenerateID("usr", constants.IDLength)
}

// GenerateRestaurantID generates a new id for a restaurant row
func GenerateRestaurantID() string {
	return GenerateID("res", constants.IDLength)
}

// GenerateMenuItemID generates a new id for a menu item
func GenerateMenuItemID() string {
	return GenerateID("itm", constants.IDLength)
}

// GenerateOrderID generates a new id for an order
func GenerateOrderID() string {
	return GenerateID("ord", constants.IDLength)
}

// GenerateOrderID generates a new id for an order
func GenerateCategoryID() string {
	return GenerateID("cat", constants.IDLength)
}

// TODO check for duplicate?
func GenerateID(prefix string, length int) string {
	b := make([]byte, length)

	for i := range b {
		b[i] = idValues[rand.Intn(len(idValues))]
	}

	if prefix != "" {
		prefix += "_"
	}

	return prefix + string(b)
}
