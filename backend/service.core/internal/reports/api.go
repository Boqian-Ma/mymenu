package reports

import (
	"github.com/COMP3900-9900-Capstone-Project/capstoneproject-comp3900-w16a-jamar/pkg/entity"
	"github.com/COMP3900-9900-Capstone-Project/capstoneproject-comp3900-w16a-jamar/pkg/errors"
	"github.com/gin-gonic/gin"
)

func RegisterHandlers(r *gin.RouterGroup, service Service, userHandler gin.HandlerFunc, managerHandler gin.HandlerFunc) {

	res := resource{service}

	managerGroup := r.Group("")
	managerGroup.Use(userHandler)
	managerGroup.Use(managerHandler)

	managerGroup.GET("/restaurants/:res_id/report", res.get)
}

type resource struct {
	service Service
}

type ItemResponse struct {
	Item *entity.HomePageReport `json:"item"`
}

// @Router /restaurants/{res_id}/report [get]
// @Tags Restaurants
// @Summary Returns the details for the specified restaurant. Currently only support "home" as a type.
// @Param res_id path string true "The id of the restaurant"
// @Param type	query string true "type of report"
// @Success 200 {object} entity.HomePageReport "success"
// @Failure 400 {object} errors.ErrorResponse "Bad Request"
// @Failure 401 {object} errors.ErrorResponse "Unauthorized"
// @Failure 404 {object} errors.ErrorResponse "Not Found"
// @Failure 500 {object} errors.ErrorResponse "Internal Server Error"
func (r resource) get(c *gin.Context) {

	reportType := c.Query("type")

	if reportType == "" {
		errors.Abort(c, errors.BadRequest("Bad Request"))
	}

	report, err := r.service.Get(c, c.Param("res_id"), reportType)

	if err != nil {
		errors.Abort(c, err)
		return
	}

	c.JSON(200, ItemResponse{
		Item: report,
	})

}
