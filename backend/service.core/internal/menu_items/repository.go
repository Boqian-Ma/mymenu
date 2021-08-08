package menu_items

import (
	"github.com/gin-gonic/gin"

	"github.com/COMP3900-9900-Capstone-Project/capstoneproject-comp3900-w16a-jamar/pkg/db"
	"github.com/COMP3900-9900-Capstone-Project/capstoneproject-comp3900-w16a-jamar/pkg/entity"
)

// Repository encapsulates the logic to access menu item data in the DB
type Repository interface {
	// Gets a list of menu items that are currently on the menu
	ListVisible(c *gin.Context, restaurantID string) ([]*entity.MenuItemResult, error)
	// Creates a new menu item
	Create(c *gin.Context, menuItem *entity.MenuItem) error
	// Get a menu item's details
	GetResult(c *gin.Context, restaurantID string, menuItemID string) (*entity.MenuItemResult, error)
	// Get a menu item's details for the service.update method
	Get(c *gin.Context, restaurantID string, menuItemID string) (*entity.MenuItem, error)
	// Returns a list of all menu item associated to current restaurant
	List(c *gin.Context, restaurantID string) ([]*entity.MenuItemResult, error)
	// Updates a menu item's details
	Update(c *gin.Context, menuItem *entity.MenuItem) error
	// Delets a manu item
	Delete(c *gin.Context, restaurantID string, menuItemID string) error
}

type repository struct {
	db db.DB
}

// NewRepository returns a new repostory that can be used to access menu item data
func NewRepository(db db.DB) Repository {
	return repository{db}
}

func (r repository) ListVisible(c *gin.Context, restaurantID string) ([]*entity.MenuItemResult, error) {
	var menuItems []*entity.MenuItemResult
	if err := r.db.Select(c, &menuItems, "select m.*, c.name as category_name from menu_items m join categories c on m.category_id=c.id where m.restaurant_id=$1 and m.is_menu=true", restaurantID); err != nil {
		return nil, err
	}
	return menuItems, nil
}

func (r repository) Create(c *gin.Context, menuItem *entity.MenuItem) error {
	return r.db.Insert(c, "menu_items", menuItem)
}

func (r repository) GetResult(c *gin.Context, restaurantID string, menuItemID string) (*entity.MenuItemResult, error) {
	var menuItem *entity.MenuItemResult
	if err := r.db.Select(c, &menuItem, "select m.*, c.name as category_name from menu_items m join categories c on m.category_id=c.id where m.restaurant_id=$1 and m.id=$2", restaurantID, menuItemID); err != nil {
		return nil, err
	}
	return menuItem, nil
}

func (r repository) Get(c *gin.Context, restaurantID string, menuItemID string) (*entity.MenuItem, error) {
	var menuItem entity.MenuItem

	if err := r.db.Select(c, &menuItem, "select * from menu_items where restaurant_id=$1 and id=$2", restaurantID, menuItemID); err != nil {
		return nil, err
	}

	return &menuItem, nil
}

// TODO pagination &/ filtering &/ sorting as req for FE
func (r repository) List(c *gin.Context, restaurantID string) ([]*entity.MenuItemResult, error) {
	var menuItems []*entity.MenuItemResult

	if err := r.db.Select(c, &menuItems, "select m.*, c.name as category_name from menu_items m join categories c on m.category_id=c.id where m.restaurant_id=$1", restaurantID); err != nil {
		return nil, err
	}
	return menuItems, nil
}

func (r repository) Update(c *gin.Context, menuItem *entity.MenuItem) error {
	return r.db.Update(c, menuItem, "update menu_items set ... where id=$1", menuItem.ID)
}

func (r repository) Delete(c *gin.Context, restaurantID string, menuItemID string) error {
	return r.db.Delete(c, "delete from menu_items where restaurant_id=$1 and id=$2", restaurantID, menuItemID)
}
