package tables

import (
	"strconv"

	"github.com/COMP3900-9900-Capstone-Project/capstoneproject-comp3900-w16a-jamar/pkg/entity"
	"github.com/COMP3900-9900-Capstone-Project/capstoneproject-comp3900-w16a-jamar/pkg/errors"
	"github.com/gin-gonic/gin"
)

// RegisterHandlers sets up the api routes for the core tables endpoints
func RegisterHandlers(r *gin.RouterGroup, service Service, userHandler, managerHandler gin.HandlerFunc) {
	res := resource{service}

	userGroup := r.Group("")
	userGroup.Use(userHandler)

	userGroup.GET("/restaurants/:res_id/tables", res.list)
	userGroup.GET("/restaurants/:res_id/tables/:tbl_num", res.get)
	userGroup.POST("/restaurants/:res_id/tables/:tbl_num/occupy", res.occupy)
	userGroup.POST("/restaurants/:res_id/tables/:tbl_num/free", res.free)

	managerGroup := r.Group("")
	managerGroup.Use(userHandler)
	managerGroup.Use(managerHandler)

	managerGroup.POST("/restaurants/:res_id/tables", res.create)
	managerGroup.PUT("/restaurants/:res_id/tables/:tbl_num", res.update)
}

type resource struct {
	service Service
}

type ItemResponse struct {
	Item *entity.Table `json:"item"`
}

type ListResponse struct {
	Data []*entity.Table `json:"data"`
}

// @Router /restaurants/{res_id}/tables [get]
// @Tags Tables
// @Summary Gets all tables associated to a restaurant
// @Param res_id path string true "The id of the restaurant"
// @Success 200 {object} ItemResponse "success"
// @Failure 400 {object} errors.ErrorResponse "Bad Request"
// @Failure 401 {object} errors.ErrorResponse "Unauthorized"
// @Failure 404 {object} errors.ErrorResponse "Not Found"
// @Failure 500 {object} errors.ErrorResponse "Internal Server Error"
func (r resource) list(c *gin.Context) {
	tables, err := r.service.List(c, c.Param("res_id"))

	if err != nil {
		errors.Abort(c, err)
		return
	}

	c.JSON(200,
		ListResponse{
			Data: tables,
		},
	)
}

// @Router /restaurants/{res_id}/tables [post]
// @Tags Tables
// @Summary Creates a new table for a restaurant
// @Param res_id path string true "The id of the restaurant"
// @Param request body tables.CreateTableRequest true "The table's details"
// @Success 200 {object} ItemResponse "success"
// @Failure 400 {object} errors.ErrorResponse "Bad Request"
// @Failure 401 {object} errors.ErrorResponse "Unauthorized"
// @Failure 404 {object} errors.ErrorResponse "Not Found"
// @Failure 500 {object} errors.ErrorResponse "Internal Server Error"
func (r resource) create(c *gin.Context) {
	var req CreateTableRequest

	err := c.BindJSON(&req)

	if err != nil {
		errors.Abort(c, errors.BadRequest(err.Error()))
		return
	}

	table, err := r.service.Create(c, c.Param("res_id"), req)

	if err != nil {
		errors.Abort(c, err)
		return
	}

	c.JSON(200, ItemResponse{
		Item: table,
	})
}

// @Router /restaurants/{res_id}/tables/{tbl_num} [put]
// @Tags Tables
// @Summary Updates a table for a restaurant
// @Param res_id path string true "The id of the restaurant"
// @Param tbl_num path string true "The id of the table"
// @Param request body tables.UpdateTableRequest true "The table's updated details"
// @Success 200 {object} ItemResponse "success"
// @Failure 400 {object} errors.ErrorResponse "Bad Request"
// @Failure 401 {object} errors.ErrorResponse "Unauthorized"
// @Failure 404 {object} errors.ErrorResponse "Not Found"
// @Failure 500 {object} errors.ErrorResponse "Internal Server Error"
func (r resource) update(c *gin.Context) {
	var req UpdateTableRequest

	err := c.BindJSON(&req)

	if err != nil {
		errors.Abort(c, errors.BadRequest(err.Error()))
		return
	}
	tableNum, _ := strconv.ParseInt(c.Param("tbl_num"), 10, 64)
	table, err := r.service.Update(c, c.Param("res_id"), tableNum, req)

	if err != nil {
		errors.Abort(c, err)
		return
	}

	c.JSON(200, ItemResponse{
		Item: table,
	})
}

// @Router /restaurants/{res_id}/tables/{tbl_num}/occupy [post]
// @Tags Tables
// @Summary Occupies a table
// @Param res_id path string true "The id of the restaurant"
// @Param tbl_num path string true "The id of the table"
// @Success 200 {object} ItemResponse "success"
// @Failure 400 {object} errors.ErrorResponse "Bad Request"
// @Failure 401 {object} errors.ErrorResponse "Unauthorized"
// @Failure 404 {object} errors.ErrorResponse "Not Found"
// @Failure 500 {object} errors.ErrorResponse "Internal Server Error"
func (r resource) occupy(c *gin.Context) {
	tableNum, _ := strconv.ParseInt(c.Param("tbl_num"), 10, 64)
	table, err := r.service.UpdateStatus(c, c.Param("res_id"), tableNum, "Taken")

	if err != nil {
		errors.Abort(c, err)
		return
	}

	c.JSON(200, ItemResponse{
		Item: table,
	})
}

// @Router /restaurants/{res_id}/tables/{tbl_num}/free [post]
// @Tags Tables
// @Summary Frees a table
// @Param res_id path string true "The id of the restaurant"
// @Param tbl_num path string true "The id of the table"
// @Success 200 {object} ItemResponse "success"
// @Failure 400 {object} errors.ErrorResponse "Bad Request"
// @Failure 401 {object} errors.ErrorResponse "Unauthorized"
// @Failure 404 {object} errors.ErrorResponse "Not Found"
// @Failure 500 {object} errors.ErrorResponse "Internal Server Error"
func (r resource) free(c *gin.Context) {
	tableNum, _ := strconv.ParseInt(c.Param("tbl_num"), 10, 64)
	table, err := r.service.UpdateStatus(c, c.Param("res_id"), tableNum, "Free")

	if err != nil {
		errors.Abort(c, err)
		return
	}

	c.JSON(200, ItemResponse{
		Item: table,
	})
}

// @Router /restaurants/{res_id}/tables/{tbl_num} [get]
// @Tags Tables
// @Summary Get a table's details
// @Param res_id path string true "The id of the restaurant"
// @Param tbl_num path string true "The id of the table"
// @Success 200 {object} ItemResponse "success"
// @Failure 400 {object} errors.ErrorResponse "Bad Request"
// @Failure 401 {object} errors.ErrorResponse "Unauthorized"
// @Failure 404 {object} errors.ErrorResponse "Not Found"
// @Failure 500 {object} errors.ErrorResponse "Internal Server Error"
func (r resource) get(c *gin.Context) {

	tableNum, _ := strconv.ParseInt(c.Param("tbl_num"), 10, 64)
	table, err := r.service.Get(c, c.Param("res_id"), tableNum)
	if err != nil {
		errors.Abort(c, err)
		return
	}

	c.JSON(200, ItemResponse{
		Item: table,
	})
}
