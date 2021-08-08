package users

import (
	"github.com/gin-gonic/gin"

	"github.com/COMP3900-9900-Capstone-Project/capstoneproject-comp3900-w16a-jamar/pkg/db"
	"github.com/COMP3900-9900-Capstone-Project/capstoneproject-comp3900-w16a-jamar/pkg/entity"
)

// Repository encapsulates the logic to access user data in the DB
type Repository interface {
	// Gets the given user by their user id
	GetUser(c *gin.Context, id string) (*entity.User, error)
	// Gets the given user by their email
	GetUserByEmail(c *gin.Context, email string) (*entity.User, error)
	// Creates a new user
	CreateUser(c *gin.Context, user *entity.User) error
	// Updates a user's details
	UpdateUser(c *gin.Context, user *entity.User) error

	// Creates a new user session
	CreateSession(c *gin.Context, session *entity.Session) error
	// Invalidate session
	InvalidateSession(c *gin.Context) error
}

type repository struct {
	db db.DB
}

// NewRepository returns a new repostory that can be used to access invoice data
func NewRepository(db db.DB) Repository {
	return repository{db}
}

func (r repository) GetUser(c *gin.Context, id string) (*entity.User, error) {
	var user entity.User

	if err := r.db.Select(c, &user, "select * from users where id=$1", id); err != nil {
		return nil, err
	}

	return &user, nil
}

func (r repository) GetUserByEmail(c *gin.Context, email string) (*entity.User, error) {
	var user entity.User

	if err := r.db.Select(c, &user, "select * from users where email=$1", email); err != nil {
		return nil, err
	}

	return &user, nil
}

func (r repository) CreateUser(c *gin.Context, user *entity.User) error {
	return r.db.Insert(c, "users", user)
}

func (r repository) UpdateUser(c *gin.Context, user *entity.User) error {
	return r.db.Update(c, user, "update users set ... where id=$1", user.ID)
}

func (r repository) CreateSession(c *gin.Context, session *entity.Session) error {
	return r.db.Insert(c, "sessions", session)
}

func (r repository) InvalidateSession(c *gin.Context) error {
	return r.db.Delete(c, "delete from sessions where user_id=$1", c.GetString("userID"))
}
