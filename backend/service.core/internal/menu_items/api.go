package menu_items

import (
	"github.com/COMP3900-9900-Capstone-Project/capstoneproject-comp3900-w16a-jamar/pkg/entity"
	"github.com/COMP3900-9900-Capstone-Project/capstoneproject-comp3900-w16a-jamar/pkg/errors"
	"github.com/gin-gonic/gin"
)

func RegisterHandlers(r *gin.RouterGroup, service Service, userHandler, managerHandler gin.HandlerFunc) {
	res := resource{service}

	userGroup := r.Group("")
	userGroup.Use(userHandler)
	// Gets a list of a visiable menu items
	userGroup.GET("/restaurants/:res_id/menu", res.listVisible)

	managerGroup := r.Group("")
	managerGroup.Use(userHandler)
	managerGroup.Use(managerHandler)

	// Gets a list of all menu items in a restaurant
	managerGroup.GET("/restaurants/:res_id/menu_items", res.list)
	// Creates a menu item
	managerGroup.POST("/restaurants/:res_id/menu_items", res.create)

	// Updates a menu item
	managerGroup.PUT("/restaurants/:res_id/menu_items/:itm_id", res.update)
	// Gets a menu item
	managerGroup.GET("/restaurants/:res_id/menu_items/:itm_id", res.get)
	// Deletes a menu item
	managerGroup.DELETE(("/restaurants/:res_id/menu_items/:itm_id"), res.delete)
}

type resource struct {
	service Service
}

type ItemResponse struct {
	Item *entity.MenuItem `json:"item"`
}

type ListResponse struct {
	Data []*entity.MenuItem `json:"data"`
}

type GetItemResponse struct {
	Item *entity.MenuItemResult `json:"item"`
}

type GetListResponse struct {
	Data []*entity.MenuItemResult `json:"data"`
}

// @Router /restaurants/{res_id}/menu [delete]
// @Tags MenuItems
// @Summary Deletes a menu item
// @Param res_id path string true "The id of the restaurant"
// @Success 200 {object} ""
// @Failure 400 {object} errors.ErrorResponse "Bad Request"
// @Failure 401 {object} errors.ErrorResponse "Unauthorized"
// @Failure 404 {object} errors.ErrorResponse "Not Found"
// @Failure 500 {object} errors.ErrorResponse "Internal Server Error"

func (r resource) delete(c *gin.Context) {

	err := r.service.Delete(c, c.Param("res_id"), c.Param("itm_id"))

	if err != nil {
		errors.Abort(c, err)
		return
	}

	c.JSON(200, "")
}

// @Router /restaurants/{res_id}/menu [get]
// @Tags MenuItems
// @Summary Gets a restaurant's current visiable menu items
// @Description Returns this restaurant's current displayed menu. It is used by customers only
// @Param res_id path string true "The id of the restaurant"
// @Success 200 {object} GetListResponse "success"
// @Failure 400 {object} errors.ErrorResponse "Bad Request"
// @Failure 401 {object} errors.ErrorResponse "Unauthorized"
// @Failure 404 {object} errors.ErrorResponse "Not Found"
// @Failure 500 {object} errors.ErrorResponse "Internal Server Error"
func (r resource) listVisible(c *gin.Context) {

	menu_items, err := r.service.ListVisible(c, c.Param("res_id"))

	if err != nil {
		errors.Abort(c, err)
		return
	}

	c.JSON(200,
		GetListResponse{
			Data: menu_items,
		},
	)
}

// @Router /restaurants/{res_id}/menu_items [get]
// @Tags MenuItems
// @Summary Gets all menu items associated to a restaurant
// @Param res_id path string true "The id of the restaurant"
// @Success 200 {object} GetListResponse "success"
// @Failure 400 {object} errors.ErrorResponse "Bad Request"
// @Failure 401 {object} errors.ErrorResponse "Unauthorized"
// @Failure 404 {object} errors.ErrorResponse "Not Found"
// @Failure 500 {object} errors.ErrorResponse "Internal Server Error"
func (r resource) list(c *gin.Context) {
	menu_items, err := r.service.List(c, c.Param("res_id"))

	if err != nil {
		errors.Abort(c, err)
		return
	}

	c.JSON(200,
		GetListResponse{
			Data: menu_items,
		},
	)
}

// @Router /restaurants/{res_id}/menu_items [post]
// @Tags MenuItems
// @Summary Creates a new menu item for a restaurant
// @Param res_id path string true "The id of the restaurant"
// @Success 200 {object} ItemResponse "success"
// @Failure 400 {object} errors.ErrorResponse "Bad Request"
// @Failure 401 {object} errors.ErrorResponse "Unauthorized"
// @Failure 404 {object} errors.ErrorResponse "Not Found"
// @Failure 500 {object} errors.ErrorResponse "Internal Server Error"
func (r resource) create(c *gin.Context) {
	var req CreateMenuItemRequest

	err := c.BindJSON(&req)

	if err != nil {
		errors.Abort(c, errors.BadRequest(err.Error()))
		return
	}

	menu_item, err := r.service.Create(c, c.Param("res_id"), req)

	if err != nil {
		errors.Abort(c, err)
		return
	}

	c.JSON(200, ItemResponse{
		Item: menu_item,
	})
}

// @Router /restaurants/{res_id}/menu_items/{itm_id} [put]
// @Tags MenuItems
// @Summary Updates a menu item for a restaurant
// @Param res_id path string true "The id of the restaurant"
// @Param itm_id path string true "The id of the item"
// @Param request body menu_items.UpdateMenuItemRequest true "The menu_item's updated details"
// @Success 200 {object} ItemResponse "success"
// @Failure 400 {object} errors.ErrorResponse "Bad Request"
// @Failure 401 {object} errors.ErrorResponse "Unauthorized"
// @Failure 404 {object} errors.ErrorResponse "Not Found"
// @Failure 500 {object} errors.ErrorResponse "Internal Server Error"
func (r resource) update(c *gin.Context) {
	var req UpdateMenuItemRequest

	err := c.BindJSON(&req)

	if err != nil {
		errors.Abort(c, errors.BadRequest(err.Error()))
		return
	}

	menu_item, err := r.service.Update(c, c.Param("res_id"), c.Param("itm_id"), req)

	if err != nil {
		errors.Abort(c, err)
		return
	}

	c.JSON(200, ItemResponse{
		Item: menu_item,
	})

}

// @Router /restaurants/{res_id}/menu_items/{itm_id} [get]
// @Tags MenuItems
// @Summary Get a menu item's details
// @Param res_id path string true "The id of the restaurant"
// @Param itm_id path string true "The id of the item"
// @Success 200 {object} GetItemResponse "success"
// @Failure 400 {object} errors.ErrorResponse "Bad Request"
// @Failure 401 {object} errors.ErrorResponse "Unauthorized"
// @Failure 404 {object} errors.ErrorResponse "Not Found"
// @Failure 500 {object} errors.ErrorResponse "Internal Server Error"
func (r resource) get(c *gin.Context) {

	menu_item, err := r.service.Get(c, c.Param("res_id"), c.Param("itm_id"))
	if err != nil {
		errors.Abort(c, err)
		return
	}

	c.JSON(200, GetItemResponse{
		Item: menu_item,
	})
}
