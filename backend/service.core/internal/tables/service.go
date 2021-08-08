package tables

import (
	"time"

	"github.com/COMP3900-9900-Capstone-Project/capstoneproject-comp3900-w16a-jamar/pkg/auth"
	"github.com/COMP3900-9900-Capstone-Project/capstoneproject-comp3900-w16a-jamar/pkg/entity"
	"github.com/COMP3900-9900-Capstone-Project/capstoneproject-comp3900-w16a-jamar/pkg/errors"
	"github.com/COMP3900-9900-Capstone-Project/capstoneproject-comp3900-w16a-jamar/service.core/internal/users"
	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// Service encapsulates the usage logic for tables service
type Service interface {
	// Creates a new table attached to the current restaurant
	Create(c *gin.Context, restaurantID string, req CreateTableRequest) (*entity.Table, error)
	// List all the tables a restaurant has
	List(c *gin.Context, restaurantID string) ([]*entity.Table, error)
	// Returns a table's details
	Get(c *gin.Context, restaurantID string, tableNum int64) (*entity.Table, error)
	// Updates a table's details
	Update(c *gin.Context, restaurantID string, tableNum int64, req UpdateTableRequest) (*entity.Table, error)
	// Updates a table's status
	UpdateStatus(c *gin.Context, restaurantID string, tableNum int64, status entity.StatusType) (*entity.Table, error)
}

type CreateTableRequest struct {
	TableNum int64 `json:"table_num"`
	NumSeats int64 `json:"num_seats"`
}

type UpdateTableRequest struct {
	Status   entity.StatusType `json:"status"`
	NumSeats int64             `json:"num_seats"`
}

func (m CreateTableRequest) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.NumSeats, validation.Required),
	)
}

func (m UpdateTableRequest) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.NumSeats, validation.Required),
	)
}

type service struct {
	repo     Repository
	userRepo users.Repository
}

// create new table service
func NewService(repo Repository, userRepo users.Repository) service {
	return service{repo, userRepo}
}

// create new table
func (s service) Create(c *gin.Context, restaurantID string, req CreateTableRequest) (*entity.Table, error) {
	err := auth.IsManagerOf(c, restaurantID)

	if err != nil {
		return nil, err
	}
	err = req.Validate()

	if err != nil {
		return nil, errors.BadRequest(err.Error())
	}

	existingTable, err := s.repo.GetTableByNumResID(c, req.TableNum, restaurantID)
	if err != nil && err.Error() != "Not Found" {
		return nil, err
	}

	if existingTable != nil {
		return nil, errors.BadRequest("Table already exists")
	}

	now := time.Now()
	table := &entity.Table{
		TableNum:     req.TableNum,
		NumSeats:     req.NumSeats,
		RestaurantID: restaurantID,
		Status:       entity.Free,
		CreatedAt:    now,
		UpdatedAt:    now,
	}
	err = s.repo.Create(c, table)

	if err != nil {
		return nil, err
	}
	return table, nil
}

// List all the tables a restaurant has
func (s service) List(c *gin.Context, restaurantID string) ([]*entity.Table, error) {

	err := auth.IsManagerOf(c, restaurantID)

	if err != nil {
		return nil, err
	}

	return s.repo.List(c, restaurantID)
}

// Returns a table's details
func (s service) Get(c *gin.Context, restaurantID string, tableNum int64) (*entity.Table, error) {
	return s.repo.Get(c, restaurantID, tableNum)
}

// Updates a table's details
func (s service) Update(c *gin.Context, restaurantID string, tableNum int64, req UpdateTableRequest) (*entity.Table, error) {

	err := auth.IsManagerOf(c, restaurantID)

	if err != nil {
		return nil, err
	}

	err = req.Validate()
	if err != nil {
		return nil, errors.BadRequest(err.Error())
	}

	table, err := s.repo.Get(c, restaurantID, tableNum)

	if err != nil {
		return nil, err
	}

	table.NumSeats = req.NumSeats
	table.Status = req.Status
	table.UpdatedAt = time.Now()

	err = s.repo.Update(c, table)

	if err != nil {
		return nil, err
	}

	return table, nil
}

// Updates a table's status
func (s service) UpdateStatus(c *gin.Context, restaurantID string, tableNum int64, status entity.StatusType) (*entity.Table, error) {
	table, err := s.repo.Get(c, restaurantID, tableNum)

	if err != nil {
		return nil, err
	}

	table.Status = status
	table.UpdatedAt = time.Now()
	err = s.repo.Update(c, table)

	if err != nil {
		return nil, err
	}

	return table, nil
}
