package categories

import (
	"time"

	"github.com/COMP3900-9900-Capstone-Project/capstoneproject-comp3900-w16a-jamar/pkg/auth"
	"github.com/COMP3900-9900-Capstone-Project/capstoneproject-comp3900-w16a-jamar/pkg/entity"
	"github.com/COMP3900-9900-Capstone-Project/capstoneproject-comp3900-w16a-jamar/pkg/errors"
	"github.com/COMP3900-9900-Capstone-Project/capstoneproject-comp3900-w16a-jamar/service.core/internal/users"
	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// Service encapsulates the usage logic for the categories service

type Service interface {
	// Creates a new menu item attached to the current restaurant
	Create(c *gin.Context, restaurantID string, req CreateCategoryRequest) (*entity.MenuItemCategory, error)
	// List all the menu items a restaurant has
	List(c *gin.Context, restaurantID string) ([]*entity.MenuItemCategory, error)
	// Get a specific menu item category
	Get(c *gin.Context, restaurantID string, categoryID string) (*entity.MenuItemCategory, error)
}

type CreateCategoryRequest struct {
	Name string `json:"name"`
}

func (m CreateCategoryRequest) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Name, validation.Required),
	)
}

type service struct {
	repo     Repository
	userRepo users.Repository
}

// NewService creates a new menu item servi ce
func NewService(repo Repository, userRepo users.Repository) service {
	return service{repo, userRepo}
}

func (s service) Create(c *gin.Context, restaurantID string, req CreateCategoryRequest) (*entity.MenuItemCategory, error) {

	err := auth.IsManagerOf(c, restaurantID)

	if err != nil {
		return nil, err
	}
	err = req.Validate()

	if err != nil {
		return nil, errors.BadRequest(err.Error())
	}

	existingCategory, err := s.repo.GetCategoryByNameResID(c, req.Name, restaurantID)
	if err != nil && err.Error() != "Not Found" {
		return nil, err
	}

	if existingCategory != nil {
		return nil, errors.BadRequest("Category already exists")
	}

	now := time.Now()

	category := &entity.MenuItemCategory{
		ID:           entity.GenerateCategoryID(),
		Name:         req.Name,
		CreatedAt:    now,
		UpdatedAt:    now,
		RestaurantID: restaurantID,
	}

	err = s.repo.Create(c, category)

	if err != nil {
		return nil, err
	}
	return category, nil
}

// Returns a menu item category's details
func (s service) Get(c *gin.Context, restaurantID string, categoryID string) (*entity.MenuItemCategory, error) {
	return s.repo.Get(c, restaurantID, categoryID)
}

// List all the menu item categories a restaurant has
func (s service) List(c *gin.Context, restaurantID string) ([]*entity.MenuItemCategory, error) {
	return s.repo.List(c, restaurantID)
}
