package users

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/google/uuid"

	"github.com/COMP3900-9900-Capstone-Project/capstoneproject-comp3900-w16a-jamar/pkg/entity"
	"github.com/COMP3900-9900-Capstone-Project/capstoneproject-comp3900-w16a-jamar/pkg/errors"
	"github.com/COMP3900-9900-Capstone-Project/capstoneproject-comp3900-w16a-jamar/pkg/utils"
)

// Service encapsulates the usage logic for the users service
type Service interface {
	// Registers a new user account
	Register(c *gin.Context, req RegisterUserRequest) (string, error)
	// Logs in a user and retrives a new session token
	Login(c *gin.Context, req LoginRequest) (string, error)
	// Logs in a guest user
	LoginGuest(c *gin.Context) (string, error)
	// Logs out the user and invalidates the session token
	Logout(c *gin.Context) error

	// Returns a user's details
	Get(c *gin.Context, userID string) (*entity.User, error)
	// Updates a user's details
	Update(c *gin.Context, req UpdateUserRequest) (*entity.User, error)

	// Reset a user's password
	ResetPassword(c *gin.Context, req ResetPasswordRequest) (*entity.User, error)
}

type RegisterUserRequest struct {
	AccountType entity.AccountType       `json:"accountType"`
	Details     UpdateUserDetailsRequest `json:"details"`
	Password    string                   `json:"password"`
}

type UpdateUserRequest struct {
	Details     UpdateUserDetailsRequest `json:"details"`
	NewPassword string                   `json:"newPassword"`
}

type UpdateUserDetailsRequest struct {
	Email   string `json:"email"`
	Name    string `json:"name"`
	Allergy string `json:"allergy"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ResetPasswordRequest struct {
	Email       string `json:"email"`
	NewPassword string `json:"newPassword"`
}

func (m RegisterUserRequest) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Password, validation.Required, validation.Length(8, 0)),
		validation.Field(&m.AccountType, validation.Required, validation.In(entity.Manager, entity.Customer)),
		validation.Field(&m.Details),
	)
}

// TODO add allergy check
func (m UpdateUserDetailsRequest) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Email, validation.Required, is.EmailFormat),
		validation.Field(&m.Name, validation.Required),
	)
}

type service struct {
	repo Repository
}

// NewService creates a new users service
func NewService(repo Repository) Service {
	return service{repo}
}

func (s service) Register(c *gin.Context, req RegisterUserRequest) (string, error) {
	if err := req.Validate(); err != nil {
		return "", errors.BadRequest(err.Error())
	}

	existingUser, err := s.repo.GetUserByEmail(c, req.Details.Email)
	if err != nil && err.Error() != "Not Found" {
		return "", err
	}

	if existingUser != nil {
		return "", errors.Conflict("An account with this email already exists")
	}

	user := s.convertRegisterRequestToUser(req)
	if err := s.repo.CreateUser(c, user); err != nil {
		return "", err
	}

	session := &entity.Session{
		Token:  uuid.New().String(),
		UserID: user.ID,
	}

	if err := s.repo.CreateSession(c, session); err != nil {
		return "", err
	}

	return session.Token, nil
}

func (s service) Login(c *gin.Context, req LoginRequest) (string, error) {
	user, err := s.repo.GetUserByEmail(c, req.Email)
	if err != nil {
		if err.Error() == "Not Found" {
			return "", errors.BadRequest("No user with that email address could be found")
		}

		return "", err
	}

	if s.hashAndSaltPassword(req.Password, user.Login.Salt) != user.Login.HashedPassword {
		return "", errors.BadRequest("Incorrect password")
	}

	session := &entity.Session{
		Token:  uuid.New().String(),
		UserID: user.ID,
	}

	if err := s.repo.CreateSession(c, session); err != nil {
		return "", err
	}

	return session.Token, nil
}

func (s service) LoginGuest(c *gin.Context) (string, error) {
	user := &entity.User{
		ID:          entity.GenerateUserID(),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		AccountType: entity.Customer,
		Details: entity.UserDetails{
			Name: "Jane Doe",
		},
		Login: entity.UserLogin{
			Salt: uuid.New().String(),
		},
	}

	if err := s.repo.CreateUser(c, user); err != nil {
		return "", err
	}

	session := &entity.Session{
		Token:  uuid.New().String(),
		UserID: user.ID,
	}

	if err := s.repo.CreateSession(c, session); err != nil {
		return "", err
	}

	return session.Token, nil
}

func (s service) Logout(c *gin.Context) error {
	if err := s.repo.InvalidateSession(c); err != nil {
		return err
	}

	return nil
}

func (s service) Update(c *gin.Context, req UpdateUserRequest) (*entity.User, error) {
	if err := req.Details.Validate(); err != nil {
		return nil, errors.BadRequest(err.Error())
	}

	user, err := s.repo.GetUser(c, c.GetString("userID"))
	if err != nil {
		return nil, err
	}

	user.Details.Email = utils.String(req.Details.Email)
	user.Details.Name = req.Details.Name
	if len(req.NewPassword) > 0 {
		user.Login.HashedPassword = s.hashAndSaltPassword(req.NewPassword, user.Login.Salt)
	}
	if err := s.repo.UpdateUser(c, user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s service) ResetPassword(c *gin.Context, req ResetPasswordRequest) (*entity.User, error) {
	user, err := s.repo.GetUserByEmail(c, req.Email)
	if err != nil {
		if err.Error() == "Not Found" {
			return nil, errors.BadRequest("No user with that email address could be found")
		}

		return nil, err
	}

	if s.hashAndSaltPassword(req.NewPassword, user.Login.Salt) == user.Login.HashedPassword {
		return nil, errors.BadRequest("New password cannot be the same as your old password")
	}

	user.Login.HashedPassword = s.hashAndSaltPassword(req.NewPassword, user.Login.Salt)
	if err := s.repo.UpdateUser(c, user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s service) Get(c *gin.Context, userID string) (*entity.User, error) {
	return s.repo.GetUser(c, userID)
}

func (s service) convertRegisterRequestToUser(req RegisterUserRequest) *entity.User {
	salt := uuid.New().String()
	return &entity.User{
		ID:          entity.GenerateUserID(),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		AccountType: req.AccountType,

		Details: entity.UserDetails{
			Email:   utils.String(req.Details.Email),
			Name:    req.Details.Name,
			Allergy: req.Details.Allergy,
		},
		Login: entity.UserLogin{
			HashedPassword: s.hashAndSaltPassword(req.Password, salt),
			Salt:           salt,
		},
	}
}

func (s service) hashAndSaltPassword(password, salt string) string {
	combined := fmt.Sprintf("%s%s", password, salt)

	h := sha256.New()
	if _, err := h.Write([]byte(combined)); err != nil {
		panic(err)
	}

	sha256Hash := hex.EncodeToString(h.Sum(nil))
	return sha256Hash
}
