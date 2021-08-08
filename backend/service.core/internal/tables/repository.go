package tables

import (
	"github.com/gin-gonic/gin"

	"github.com/COMP3900-9900-Capstone-Project/capstoneproject-comp3900-w16a-jamar/pkg/db"
	"github.com/COMP3900-9900-Capstone-Project/capstoneproject-comp3900-w16a-jamar/pkg/entity"
)

// Repository encapsulates the logic to access table data in the DB
type Repository interface {
	// Creates a new table
	Create(c *gin.Context, table *entity.Table) error
	// Get a table's details
	Get(c *gin.Context, restaurantID string, tableNum int64) (*entity.Table, error)
	// Get a table by number and restaurant id
	GetTableByNumResID(c *gin.Context, tableNum int64, restaurantID string) (*entity.Table, error)
	// Returns a list of all table associated to current restaurant
	List(c *gin.Context, restaurantID string) ([]*entity.Table, error)
	// Updates a table's details
	Update(c *gin.Context, table *entity.Table) error
}

type repository struct {
	db db.DB
}

// NewRepository returns a new repostory that can be used to access menu item data
func NewRepository(db db.DB) Repository {
	return repository{db}
}

func (r repository) Create(c *gin.Context, table *entity.Table) error {
	return r.db.Insert(c, "tables", table)
}

func (r repository) Get(c *gin.Context, restaurantID string, tableNum int64) (*entity.Table, error) {
	var table entity.Table

	if err := r.db.Select(c, &table, "select * from tables where restaurant_id=$1 and table_num=$2", restaurantID, tableNum); err != nil {
		return nil, err
	}

	return &table, nil
}

func (r repository) GetTableByNumResID(c *gin.Context, tableNum int64, restaurantID string) (*entity.Table, error) {
	var table entity.Table
	if err := r.db.Select(c, &table, "select * from tables where table_num=$1 and restaurant_id=$2", tableNum, restaurantID); err != nil {
		return nil, err
	}

	return &table, nil
}

// TODO pagination &/ filtering &/ sorting as req for FE
func (r repository) List(c *gin.Context, restaurantID string) ([]*entity.Table, error) {
	var tables []*entity.Table

	if err := r.db.Select(c, &tables, "select * from tables where restaurant_id=$1 order by table_num asc", restaurantID); err != nil {
		return nil, err
	}

	return tables, nil
}

func (r repository) Update(c *gin.Context, table *entity.Table) error {
	return r.db.Update(c, table, "update tables set ... where restaurant_id=$1 and table_num=$2", table.RestaurantID, table.TableNum)
}
