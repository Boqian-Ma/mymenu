package auth

import (
	"github.com/gin-gonic/gin"

	"github.com/COMP3900-9900-Capstone-Project/capstoneproject-comp3900-w16a-jamar/pkg/db"
	"github.com/COMP3900-9900-Capstone-Project/capstoneproject-comp3900-w16a-jamar/pkg/entity"
)

// Repository encapsulates the logic to access user data in the DB
type Repository interface {
	// Retrieves a user from their session token
	GetUserBySessionToken(c *gin.Context, sessionToken string) (*entity.User, error)
	// Retrieves the list of restaurants this user has access to
	ListRestaurantsForUser(c *gin.Context) ([]string, error)
}

type repository struct {
	db db.DB
}

// NewRepository returns a new repostory that can be used to access invoice data
func NewRepository(db db.DB) Repository {
	return repository{db}
}

func (r repository) GetUserBySessionToken(c *gin.Context, sessionToken string) (*entity.User, error) {
	var session entity.Session
	if err := r.db.Select(c, &session, "select * from sessions where token=$1", sessionToken); err != nil {
		return nil, err
	}

	var user entity.User
	if err := r.db.Select(c, &user, "select * from users where id=$1", session.UserID); err != nil {
		return nil, err
	}

	return &user, nil
}

func (r repository) ListRestaurantsForUser(c *gin.Context) ([]string, error) {
	var resMembers []*entity.RestaurantMember
	if err := r.db.Select(c, &resMembers, "select * from restaurant_members where user_id=$1", c.GetString("userID")); err != nil {
		return nil, err
	}

	var restaurantIDs []string
	for _, member := range resMembers {
		restaurantIDs = append(restaurantIDs, member.RestaurantID)
	}

	return restaurantIDs, nil
}
