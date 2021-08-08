package categories

import (
	"github.com/COMP3900-9900-Capstone-Project/capstoneproject-comp3900-w16a-jamar/pkg/entity"
	"github.com/COMP3900-9900-Capstone-Project/capstoneproject-comp3900-w16a-jamar/pkg/errors"
	"github.com/gin-gonic/gin"
)

func RegisterHandlers(r *gin.RouterGroup, service Service, userHandler, managerHandler gin.HandlerFunc) {
	res := resource{service}

	userGroup := r.Group("")
	userGroup.Use(userHandler)
	// Gets a list of a visible categories
	userGroup.GET("/restaurants/:res_id/categories", res.list)
	// Get a specific category
	userGroup.GET("/restaurants/:res_id/categories/:cat_id", res.get)

	managerGroup := r.Group("")
	managerGroup.Use(userHandler)
	managerGroup.Use(managerHandler)

	// Creates a category for a restaurant
	managerGroup.POST("/restaurants/:res_id/categories", res.create)
}

type resource struct {
	service Service
}

type ItemResponse struct {
	Item *entity.MenuItemCategory `json:"item"`
}

type ListResponse struct {
	Data []*entity.MenuItemCategory `json:"data"`
}

// @Router /restaurants/{res_id}/categories [post]
// @Tags MenuItems
// @Summary Creates a new menu item category for a restaurant
// @Param res_id path string true "The id of the restaurant"
// @Success 200 {object} ItemResponse "success"
// @Failure 400 {object} errors.ErrorResponse "Bad Request"
// @Failure 401 {object} errors.ErrorResponse "Unauthorized"
// @Failure 404 {object} errors.ErrorResponse "Not Found"
// @Failure 500 {object} errors.ErrorResponse "Internal Server Error"
func (r resource) create(c *gin.Context) {
	var req CreateCategoryRequest

	err := c.BindJSON(&req)

	if err != nil {
		errors.Abort(c, errors.BadRequest(err.Error()))
		return
	}

	menu_item_category, err := r.service.Create(c, c.Param("res_id"), req)

	if err != nil {
		errors.Abort(c, err)
		return
	}

	c.JSON(200, ItemResponse{
		Item: menu_item_category,
	})
}

// @Router /restaurants/{res_id}/categories/{cat_id} [get]
// @Tags MenuItems
// @Summary Returns a category from an id
// @Param res_id path string true "The id of the restaurant"
// @Param cat_id path string true "The id of the category"
// @Success 200 {object} ItemResponse "success"
// @Failure 400 {object} errors.ErrorResponse "Bad Request"
// @Failure 401 {object} errors.ErrorResponse "Unauthorized"
// @Failure 404 {object} errors.ErrorResponse "Not Found"
// @Failure 500 {object} errors.ErrorResponse "Internal Server Error"
func (r resource) get(c *gin.Context) {
	menu_item_category, err := r.service.Get(c, c.Param("res_id"), c.Param("cat_id"))
	if err != nil {
		errors.Abort(c, err)
		return
	}

	c.JSON(200, ItemResponse{
		Item: menu_item_category,
	})
}

// @Router /restaurants/{res_id}/categories [get]
// @Tags MenuItems
// @Summary Gets all menu item categories associated to a restaurant
// @Param res_id path string true "The id of the restaurant"
// @Success 200 {object} ListResponse "success"
// @Failure 400 {object} errors.ErrorResponse "Bad Request"
// @Failure 401 {object} errors.ErrorResponse "Unauthorized"
// @Failure 404 {object} errors.ErrorResponse "Not Found"
// @Failure 500 {object} errors.ErrorResponse "Internal Server Error"
func (r resource) list(c *gin.Context) {
	menu_item_categories, err := r.service.List(c, c.Param("res_id"))

	if err != nil {
		errors.Abort(c, err)
		return
	}

	c.JSON(200,
		ListResponse{
			Data: menu_item_categories,
		},
	)
}
