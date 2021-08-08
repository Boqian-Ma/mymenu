package auth

import (
	"github.com/gin-gonic/gin"

	"github.com/COMP3900-9900-Capstone-Project/capstoneproject-comp3900-w16a-jamar/pkg/errors"
)

func IsManagerOf(c *gin.Context, restaurantID string) error {
	restaurantIDs := c.GetStringSlice("restaurantIDs")

	for _, id := range restaurantIDs {

		if id == restaurantID {
			return nil
		}
	}

	return errors.Forbidden("You don't have permission to manage this restaurant")
}
