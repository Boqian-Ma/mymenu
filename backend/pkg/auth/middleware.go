package auth

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/COMP3900-9900-Capstone-Project/capstoneproject-comp3900-w16a-jamar/pkg/entity"
	"github.com/COMP3900-9900-Capstone-Project/capstoneproject-comp3900-w16a-jamar/pkg/errors"
)

// UserHandler ensures the current user is properly authenticated
func UserHandler(repo Repository) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				errors.Abort(c, errors.InternalServerError(""))
				return
			}
		}()

		sessionToken := c.GetHeader("Authorization")
		if sessionToken == "" {
			errors.Abort(c, errors.Unauthorized(""))
			return
		}

		sessionToken = strings.TrimPrefix(sessionToken, "Bearer ")

		fmt.Println("info: verifying authorization token")

		user, err := repo.GetUserBySessionToken(c, sessionToken)
		if err != nil {
			if err.Error() == "Not Found" {
				errors.Abort(c, errors.Unauthorized(""))
				return
			} else {
				errors.Abort(c, err)
				return
			}
		}

		c.Set("userID", user.ID)
		c.Set("accountType", string(user.AccountType))

		if user.AccountType == entity.Manager {
			restaurantIDs, err := repo.ListRestaurantsForUser(c)
			if err != nil {
				errors.Abort(c, err)
				return
			}

			c.Set("restaurantIDs", restaurantIDs)
		}

		fmt.Println("info: user successfully authenticated " + user.ID)
	}
}

// ManagerHandler ensures the current user is a manager, should be used after UserHandler
func ManagerHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				errors.Abort(c, errors.InternalServerError(""))
				return
			}
		}()

		if c.GetString("accountType") != string(entity.Manager) {
			errors.Abort(c, errors.Forbidden("Only managers can perform this action"))
			return
		}
	}
}
