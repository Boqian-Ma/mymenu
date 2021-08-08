package orders

import (
	"fmt"
	"strconv"

	"github.com/COMP3900-9900-Capstone-Project/capstoneproject-comp3900-w16a-jamar/pkg/entity"
	"github.com/COMP3900-9900-Capstone-Project/capstoneproject-comp3900-w16a-jamar/pkg/errors"
	"github.com/gin-gonic/gin"
)

func RegisterHandlers(r *gin.RouterGroup, service Service, userHandler, managerHandler gin.HandlerFunc) {
	res := resource{service}

	userGroup := r.Group("")
	userGroup.Use(userHandler)

	// Gets a list of a visiable menu items
	userGroup.GET("/orders", res.getOrders)
	userGroup.POST("/orders", res.createOrder)
	userGroup.GET("/order", res.getTableOrder)

	managerGroup := r.Group("")
	managerGroup.Use(userHandler)
	managerGroup.Use(managerHandler)

	// Updates an order details
	managerGroup.PUT("/orders/:ord_id", res.updateOrder)
	// Move an order to a different table
	managerGroup.PUT("/orders/:ord_id/move", res.moveOrder)
	// Change an order's status to cooked
	managerGroup.POST("/orders/:ord_id/serve", res.orderServed)
	// Change an order's stauts to complete
	managerGroup.POST("/orders/:ord_id/complete", res.orderCompleted)
	// Change an order's stauts to cancel
	managerGroup.POST(("/orders/:ord_id/cancel"), res.orderCanceled)

	// Creates a new order Item an order details
	//managerGroup.GET("/orders/:ord_id/items/:itm_id", res.createOrderItem)
	// Creates a new order Item an order details
	managerGroup.POST("/orders/:ord_id", res.createOrderItem)
	// Updates an order item (quantity)
	managerGroup.PUT("/orders/:ord_id/items/:itm_id", res.updateOrderItem)
	// Deletes an order item from an order
	managerGroup.DELETE("/orders/:ord_id/items/:itm_id", res.deleteOrderItem)
}

type resource struct {
	service Service
}

type ItemResponse struct {
	Item *entity.Order `json:"item"`
}

type ListResponse struct {
	Data []*entity.Order `json:"data"`
}

type OrderItemResponse struct {
	Item *entity.OrderItem `json:"item"`
}

type GetItemResponse struct {
	Item *entity.OrderResult `json:"item"`
}

type GetListResponse struct {
	Data []*entity.OrderResult `json:"data"`
}

type GetOrderItemResponse struct {
	Item *entity.OrderItemResult `json:"item"`
}

// @Router /orders [get]
// @Tags Orders
// @Summary Gets a list of orders depending on the query parameters
// @Param res_id query string false "If supplied by user, only return orders made by this restaurant"
// @Param usr_id query string false "If supplied by user, only return orders made by this specific user"
// @Param status query string false "If supplied, return a list of order based on the status code. Should only be used by restaurant"
// @Success 200 {object} GetListResponse "success"
// @Failure 400 {object} errors.ErrorResponse "Bad Request"
// @Failure 401 {object} errors.ErrorResponse "Unauthorized"
// @Failure 404 {object} errors.ErrorResponse "Not Found"
// @Failure 500 {object} errors.ErrorResponse "Internal Server Error"
func (r resource) getOrders(c *gin.Context) {
	// Get res order + status passed

	res_id := string(c.Query("res_id"))
	usr_id := string(c.Query("usr_id"))
	status := entity.OrderStatus(c.Query("status"))
	activeParam := c.Query("active")
	active := false
	// Get restaurant orders
	if res_id != "" {
		if usr_id != "" {
			errors.Abort(c, errors.BadRequest("Incorrect query parameter usage"))
			return
		}
		if activeParam != "" {
			var err error
			active, err = strconv.ParseBool(activeParam)
			if err != nil {
				errors.Abort(c, errors.BadRequest("Couldn't parse query param 'active': "+err.Error()))
				return
			}
		}

		orders, err := r.service.GetRestaurantOrders(c, res_id, status, active)

		if err != nil {
			errors.Abort(c, err)
			return
		}

		c.JSON(200,
			GetListResponse{
				Data: orders,
			},
		)
	}

	// Get customer orders
	// Tested
	if usr_id != "" {
		if res_id != "" {
			errors.Abort(c, errors.BadRequest("Incorrect query parameter usage"))
			return
		}

		if usr_id != c.GetString("userID") {
			errors.Abort(c, errors.BadRequest("Permission Error:Please only look for your own order"))
			return
		}

		orders, err := r.service.GetCustomerOrders(c, usr_id)

		if err != nil {
			errors.Abort(c, err)
			return
		}

		c.JSON(200,
			GetListResponse{
				Data: orders,
			},
		)
	}

}

// @Router /order [get]
// @Tags Orders
// @Summary Gets a list of orders depending on the query parameters
// @Param res_id query string true "If supplied by user, only return orders made by this table"
// @Param table_num query string true "If supplied by user, only return orders made by this specific user"
// @Success 200 {object} GetItemResponse "success"
// @Failure 400 {object} errors.ErrorResponse "Bad Request"
// @Failure 401 {object} errors.ErrorResponse "Unauthorized"
// @Failure 404 {object} errors.ErrorResponse "Not Found"
// @Failure 500 {object} errors.ErrorResponse "Internal Server Error"
func (r resource) getTableOrder(c *gin.Context) {
	res_id := string(c.Query("res_id"))
	table_num_str := string(c.Query("table_num"))

	if res_id == "" || table_num_str == "" {
		errors.Abort(c, errors.BadRequest("Must provide a res_id and table_num"))
		return
	}
	table_num, _ := strconv.ParseInt(table_num_str, 10, 64)
	order, err := r.service.GetTableOrder(c, res_id, table_num)
	if err != nil {
		errors.Abort(c, err)
		return
	}
	if order == nil {
		order = &entity.OrderResult{}
	}
	fmt.Println(order)
	c.JSON(200, GetItemResponse{
		Item: order,
	})
}

// @Router /orders [post]
// @Tags Orders
// @Summary Creates a order. Used by both customer and users
// @Param request body orders.CreateOrderRequest true "The new order's details"
// @Success 200 {object} ItemResponse "success"
// @Failure 400 {object} errors.ErrorResponse "Bad Request"
// @Failure 401 {object} errors.ErrorResponse "Unauthorized"
// @Failure 404 {object} errors.ErrorResponse "Not Found"
// @Failure 500 {object} errors.ErrorResponse "Internal Server Error"
func (r resource) createOrder(c *gin.Context) {
	// Create order: restaurant passed

	var req CreateOrderRequest

	if err := c.BindJSON(&req); err != nil {
		errors.Abort(c, errors.BadRequest(err.Error()))
		return
	}

	order, err := r.service.Create(c, req)
	if err != nil {
		errors.Abort(c, err)
		return
	}

	c.JSON(200, ItemResponse{
		Item: order,
	})
}

// @Router /orders/{ord_id} [put]
// @Tags Orders
// @Summary Updates an order
// @Param request body orders.UpdateOrderRequest true "The order's updated details"
// @Success 200 {object} ItemResponse "success"
// @Failure 400 {object} errors.ErrorResponse "Bad Request"
// @Failure 401 {object} errors.ErrorResponse "Unauthorized"
// @Failure 404 {object} errors.ErrorResponse "Not Found"
// @Failure 500 {object} errors.ErrorResponse "Internal Server Error"
func (r resource) updateOrder(c *gin.Context) {
	var req UpdateOrderRequest

	err := c.BindJSON(&req)

	if err != nil {
		errors.Abort(c, errors.BadRequest(err.Error()))
		return
	}

	order, err := r.service.Update(c, c.Param("ord_id"), req)

	if err != nil {
		errors.Abort(c, err)
		return
	}

	c.JSON(200, ItemResponse{
		Item: order,
	})
}

// @Router /orders/{ord_id}/move [put]
// @Tags Orders
// @Summary Updates an order
// @Param request body orders.MoveOrderRequest true "Move an order to different table"
// @Success 200 {object} ItemResponse "success"
// @Failure 400 {object} errors.ErrorResponse "Bad Request"
// @Failure 401 {object} errors.ErrorResponse "Unauthorized"
// @Failure 404 {object} errors.ErrorResponse "Not Found"
// @Failure 500 {object} errors.ErrorResponse "Internal Server Error"
func (r resource) moveOrder(c *gin.Context) {
	var req MoveOrderRequest

	err := c.BindJSON(&req)

	if err != nil {
		errors.Abort(c, errors.BadRequest(err.Error()))
		return
	}

	order, err := r.service.MoveOrder(c, c.Param("ord_id"), req)

	if err != nil {
		errors.Abort(c, err)
		return
	}

	c.JSON(200, ItemResponse{
		Item: order,
	})
}

// @Router /orders/{ord_id}/cook [post]
// @Tags Orders
// @Summary Sets an order status to "served"
// @Param ord_id path string true "The id of the order"
// @Param request body orders.UpdateOrderRequest true "The order's updated details"
// @Success 200 {object} ItemResponse "success"
// @Failure 400 {object} errors.ErrorResponse "Bad Request"
// @Failure 401 {object} errors.ErrorResponse "Unauthorized"
// @Failure 404 {object} errors.ErrorResponse "Not Found"
// @Failure 500 {object} errors.ErrorResponse "Internal Server Error"
func (r resource) orderServed(c *gin.Context) {

	order, err := r.service.SetOrderStatus(c, c.Param("ord_id"), entity.OrderStatus("served"))

	if err != nil {
		errors.Abort(c, err)
		return
	}

	c.JSON(200, ItemResponse{
		Item: order,
	})

}

// @Router /orders/{ord_id}/complete [post]
// @Tags Orders
// @Summary Sets an order status to "completed"
// @Param ord_id path string true "The id of the order"
// @Param request body orders.UpdateOrderRequest true "The order's updated details"
// @Success 200 {object} ItemResponse "success"
// @Failure 400 {object} errors.ErrorResponse "Bad Request"
// @Failure 401 {object} errors.ErrorResponse "Unauthorized"
// @Failure 404 {object} errors.ErrorResponse "Not Found"
// @Failure 500 {object} errors.ErrorResponse "Internal Server Error"
func (r resource) orderCompleted(c *gin.Context) {

	order, err := r.service.SetOrderStatus(c, c.Param("ord_id"), entity.OrderStatus("completed"))

	if err != nil {
		errors.Abort(c, err)
		return
	}

	c.JSON(200, ItemResponse{
		Item: order,
	})
}

// @Router /orders/{ord_id}/cancel [post]
// @Tags Orders
// @Summary Sets an order status to "canceled"
// @Param ord_id path string true "The id of the order"
// @Param request body orders.UpdateOrderRequest true "The order's updated details"
// @Success 200 {object} ItemResponse "success"
// @Failure 400 {object} errors.ErrorResponse "Bad Request"
// @Failure 401 {object} errors.ErrorResponse "Unauthorized"
// @Failure 404 {object} errors.ErrorResponse "Not Found"
// @Failure 500 {object} errors.ErrorResponse "Internal Server Error"
func (r resource) orderCanceled(c *gin.Context) {

	order, err := r.service.SetOrderStatus(c, c.Param("ord_id"), entity.OrderStatus("canceled"))

	if err != nil {
		errors.Abort(c, err)
		return
	}

	c.JSON(200, ItemResponse{
		Item: order,
	})
}

// @Router /orders/{ord_id}/items/{itm_id} [get]
// @Tags Orders
// @Summary Gets the details of an order item
// @Param ord_id path string true "The id of the order"
// @Param itm_id path string true "The id of the item in the order"
// @Success 200 {object} ItemResponse "success"
// @Failure 400 {object} errors.ErrorResponse "Bad Request"
// @Failure 401 {object} errors.ErrorResponse "Unauthorized"
// @Failure 404 {object} errors.ErrorResponse "Not Found"
// @Failure 500 {object} errors.ErrorResponse "Internal Server Error"
// func (r resource) getOrderItem(c *gin.Context) {
// 	order_item, err := r.service.GetOrd
// }

// @Router /orders/{ord_id} [post]
// @Tags Orders
// @Summary Creates a new order item in an order
// @Param ord_id path string true "The id of the order"
// @Success 200 {object} OrderItemResponse "success"
// @Failure 400 {object} errors.ErrorResponse "Bad Request"
// @Failure 401 {object} errors.ErrorResponse "Unauthorized"
// @Failure 404 {object} errors.ErrorResponse "Not Found"
// @Failure 500 {object} errors.ErrorResponse "Internal Server Error"
func (r resource) createOrderItem(c *gin.Context) {
	var req CreateOrderItemRequest

	err := c.BindJSON(&req)

	if err != nil {
		errors.Abort(c, errors.BadRequest(err.Error()))
		return
	}

	order_item, err := r.service.CreateOrderItem(c, req)

	if err != nil {
		errors.Abort(c, err)
		return
	}

	c.JSON(200, OrderItemResponse{
		Item: order_item,
	})
}

// @Router /orders/{ord_id} [put]
// @Tags Orders
// @Summary Updates an order item in an order.
// @Param ord_id path string true "The id of the order"
// @Param request body orders.UpdateOrderItemRequest true "The order item's details"
// @Success 200 {object} OrderItemResponse "success"
// @Failure 400 {object} errors.ErrorResponse "Bad Request"
// @Failure 401 {object} errors.ErrorResponse "Unauthorized"
// @Failure 404 {object} errors.ErrorResponse "Not Found"
// @Failure 500 {object} errors.ErrorResponse "Internal Server Error"
func (r resource) updateOrderItem(c *gin.Context) {

	/*
		itm_id := string(c.Query("itm_id"))

		if itm_id == "" {
			errors.Abort(c, errors.BadRequest("Bad Request"))
			return
		}
	*/

	var req UpdateOrderItemRequest

	err := c.BindJSON(&req)

	if err != nil {
		errors.Abort(c, errors.BadRequest(err.Error()))
		return
	}

	order_item, err := r.service.UpdateOrderItem(c, req)

	if err != nil {
		errors.Abort(c, err)
		return
	}

	c.JSON(200, OrderItemResponse{
		Item: order_item,
	})

}

// @Router /orders/{ord_id}/items/{itm_id} [delete]
// @Tags Orders
// @Summary Deletes an order item from an order
// @Param ord_id path string true "The id of the order"
// @Param itm_id path string true "The id of the item in the order"
// @Param request body orders.UpdateOrderItemRequest true "The order item's details"
// @Success 200 {object} ItemResponse "success"
// @Failure 400 {object} errors.ErrorResponse "Bad Request"
// @Failure 401 {object} errors.ErrorResponse "Unauthorized"
// @Failure 404 {object} errors.ErrorResponse "Not Found"
// @Failure 500 {object} errors.ErrorResponse "Internal Server Error"
func (r resource) deleteOrderItem(c *gin.Context) {

	err := r.service.DeleteOrderItem(c, c.Param("ord_id"), c.Param("itm_id"))

	if err != nil {
		errors.Abort(c, err)
		return
	}

	c.JSON(200, "")
}
