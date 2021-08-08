package restaurants

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/COMP3900-9900-Capstone-Project/capstoneproject-comp3900-w16a-jamar/pkg/entity"
	"github.com/COMP3900-9900-Capstone-Project/capstoneproject-comp3900-w16a-jamar/pkg/errors"
)

// RegisterHandlers sets up the api routes for the core restaurants endpoints
func RegisterHandlers(r *gin.RouterGroup, service Service, userHandler, managerHandler gin.HandlerFunc) {
	res := resource{service}

	userGroup := r.Group("")
	userGroup.Use(userHandler)

	userGroup.GET("/restaurants", res.list)
	userGroup.GET("/restaurants/:res_id", res.get)

	userGroup.GET("/recommended-restaurants", res.recommendedList)

	managerGroup := r.Group("")
	managerGroup.Use(userHandler)
	managerGroup.Use(managerHandler)

	managerGroup.POST("/restaurants", res.create)
	managerGroup.PUT("/restaurants/:res_id", res.update)
}

type resource struct {
	service Service
}

type ItemResponse struct {
	Item *entity.Restaurant `json:"item"`
}

type ListResponse struct {
	Data []*entity.Restaurant `json:"data"`
}

// @Router /restaurants [post]
// @Tags Restaurants
// @Summary Creates a new restaurant attached to the user's account
// @Param request body restaurants.CreateRestaurantRequest true "The new restaurant's details"
// @Success 200 {object} ItemResponse "success"
// @Failure 400 {object} errors.ErrorResponse "Bad Request"
// @Failure 401 {object} errors.ErrorResponse "Unauthorized"
// @Failure 404 {object} errors.ErrorResponse "Not Found"
// @Failure 500 {object} errors.ErrorResponse "Internal Server Error"
func (r resource) create(c *gin.Context) {
	var req CreateRestaurantRequest

	if err := c.BindJSON(&req); err != nil {
		errors.Abort(c, errors.BadRequest(err.Error()))
		return
	}

	restaurant, err := r.service.Create(c, req)
	if err != nil {
		errors.Abort(c, err)
		return
	}
	c.JSON(200, ItemResponse{
		Item: restaurant,
	})
}

// @Router /restaurants [get]
// @Tags Restaurants
// @Summary Returns a list of restaurants
// @Description By default returns all restaurants in the system but managers can use the `mine=true` query param to only return restaurants they manage
// @Param mine query bool false "If it's true, only return restaurants managed by the user"
// @Success 200 {object} ListResponse "success"
// @Failure 400 {object} errors.ErrorResponse "Bad Request"
// @Failure 401 {object} errors.ErrorResponse "Unauthorized"
// @Failure 404 {object} errors.ErrorResponse "Not Found"
// @Failure 500 {object} errors.ErrorResponse "Internal Server Error"
func (r resource) list(c *gin.Context) {
	mine := false
	param := c.Query("mine")
	if param != "" {
		var err error
		mine, err = strconv.ParseBool(param)
		if err != nil {
			errors.Abort(c, errors.BadRequest("Couldn't parse query param 'mine': "+err.Error()))
			return
		}
	}

	restaurants, err := r.service.List(c, mine)
	if err != nil {
		errors.Abort(c, err)
		return
	}

	c.JSON(200, ListResponse{
		Data: restaurants,
	})
}

// @Router /restaurants/{res_id} [get]
// @Tags Restaurants
// @Summary Returns the details for the specified restaurant
// @Param res_id path string true "The id of the restaurant"
// @Success 200 {object} entity.Restaurant "success"
// @Failure 400 {object} errors.ErrorResponse "Bad Request"
// @Failure 401 {object} errors.ErrorResponse "Unauthorized"
// @Failure 404 {object} errors.ErrorResponse "Not Found"
// @Failure 500 {object} errors.ErrorResponse "Internal Server Error"
func (r resource) get(c *gin.Context) {
	restaurant, err := r.service.Get(c, c.Param("res_id"))
	if err != nil {
		errors.Abort(c, err)
		return
	}

	c.JSON(200, ItemResponse{
		Item: restaurant,
	})
}

// @Router /restaurants/{res_id} [put]
// @Tags Restaurants
// @Summary Updates a restaurant's details
// @Param res_id path string true "The id of the restaurant"
// @Param request body restaurants.UpdateRestaurantRequest true "The restaurant's updated details"
// @Success 200 {object} entity.Restaurant "success"
// @Failure 400 {object} errors.ErrorResponse "Bad Request"
// @Failure 401 {object} errors.ErrorResponse "Unauthorized"
// @Failure 404 {object} errors.ErrorResponse "Not Found"
// @Failure 500 {object} errors.ErrorResponse "Internal Server Error"
func (r resource) update(c *gin.Context) {
	var req UpdateRestaurantRequest
	if err := c.BindJSON(&req); err != nil {
		errors.Abort(c, errors.BadRequest(err.Error()))
		return
	}

	restaurant, err := r.service.Update(c, c.Param("res_id"), req)
	if err != nil {
		errors.Abort(c, err)
		return
	}

	c.JSON(200, ItemResponse{
		Item: restaurant,
	})
}

// @Router /recommended-restaurants [get]
// @Tags Restaurants
// @Summary Returns a list of restaurants, recommended for this specific user
// @Description By default returns all restaurants recommended based on this user's previous order history. For detailed explaination, please read our software report
// @Success 200 {object} ListResponse "success"
// @Failure 400 {object} errors.ErrorResponse "Bad Request"
// @Failure 401 {object} errors.ErrorResponse "Unauthorized"
// @Failure 404 {object} errors.ErrorResponse "Not Found"
// @Failure 500 {object} errors.ErrorResponse "Internal Server Error"
func (r resource) recommendedList(c *gin.Context) {

	restaurants, err := r.service.RecommendedList(c)
	if err != nil {
		errors.Abort(c, err)
		return
	}

	c.JSON(200, ListResponse{
		Data: restaurants,
	})
}
