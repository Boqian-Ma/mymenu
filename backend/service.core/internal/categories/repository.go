package categories

import (
	"github.com/gin-gonic/gin"

	"github.com/COMP3900-9900-Capstone-Project/capstoneproject-comp3900-w16a-jamar/pkg/db"
	"github.com/COMP3900-9900-Capstone-Project/capstoneproject-comp3900-w16a-jamar/pkg/entity"
)

// Repository encapsulates the logic to access menu item data in the DB
type Repository interface {
	// Creates a new menu item
	Create(c *gin.Context, menuItemCategory *entity.MenuItemCategory) error
	// Returns a list of all menu item associated to current restaurant
	List(c *gin.Context, restaurantID string) ([]*entity.MenuItemCategory, error)
	// Return the category based on restaurant id and category id
	Get(c *gin.Context, restaurantID string, categoryID string) (*entity.MenuItemCategory, error)
	// Returns the category based off name and restaurant id
	GetCategoryByNameResID(c *gin.Context, name string, restaurantID string) (*entity.MenuItemCategory, error)
}

type repository struct {
	db db.DB
}

// NewRepository returns a new repostory that can be used to access menu item data
func NewRepository(db db.DB) Repository {
	return repository{db}
}

func (r repository) Create(c *gin.Context, menuItemCategory *entity.MenuItemCategory) error {
	return r.db.Insert(c, "categories", menuItemCategory)
}

func (r repository) Get(c *gin.Context, restaurantID string, categoryID string) (*entity.MenuItemCategory, error) {
	var menuItemCategory entity.MenuItemCategory

	if err := r.db.Select(c, &menuItemCategory, "select * from categories where restaurant_id=$1 and id=$2", restaurantID, categoryID); err != nil {
		return nil, err
	}

	return &menuItemCategory, nil
}

func (r repository) GetCategoryByNameResID(c *gin.Context, name string, restaurantID string) (*entity.MenuItemCategory, error) {
	var menuItemCategory entity.MenuItemCategory
	if err := r.db.Select(c, &menuItemCategory, "select * from categories where name=$1 and restaurant_id=$2", name, restaurantID); err != nil {
		return nil, err
	}

	return &menuItemCategory, nil
}

func (r repository) List(c *gin.Context, restaurantID string) ([]*entity.MenuItemCategory, error) {
	var menuItemCategories []*entity.MenuItemCategory
	if err := r.db.Select(c, &menuItemCategories, "select * from categories where restaurant_id=$1", restaurantID); err != nil {
		return nil, err
	}

	return menuItemCategories, nil
}
