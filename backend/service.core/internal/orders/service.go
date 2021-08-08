package orders

import (
	"fmt"
	"regexp"
	"time"

	"github.com/COMP3900-9900-Capstone-Project/capstoneproject-comp3900-w16a-jamar/pkg/auth"
	"github.com/COMP3900-9900-Capstone-Project/capstoneproject-comp3900-w16a-jamar/pkg/entity"
	"github.com/COMP3900-9900-Capstone-Project/capstoneproject-comp3900-w16a-jamar/pkg/errors"
	"github.com/COMP3900-9900-Capstone-Project/capstoneproject-comp3900-w16a-jamar/service.core/internal/menu_items"
	"github.com/COMP3900-9900-Capstone-Project/capstoneproject-comp3900-w16a-jamar/service.core/internal/tables"
	"github.com/COMP3900-9900-Capstone-Project/capstoneproject-comp3900-w16a-jamar/service.core/internal/users"
	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type Service interface {
	// Returns details of an order tested
	Get(c *gin.Context, orderID string) (*entity.Order, error)
	// Returns details of an order based on res_id and table_num
	GetTableOrder(c *gin.Context, res_id string, table_num int64) (*entity.OrderResult, error)
	// Creates an order // tested
	Create(c *gin.Context, req CreateOrderRequest) (*entity.Order, error)
	// // Removes an order
	// Delete(c *gin.Context, orderID string) error
	// Updates an order tested
	Update(c *gin.Context, orderID string, req UpdateOrderRequest) (*entity.Order, error)

	// Move the table for an order
	MoveOrder(c *gin.Context, orderID string, req MoveOrderRequest) (*entity.Order, error)

	// Set order status tested
	SetOrderStatus(c *gin.Context, orderID string, status entity.OrderStatus) (*entity.Order, error)

	// Returns a list of orders a restaurant has tested
	GetRestaurantOrders(c *gin.Context, restaurantID string, status entity.OrderStatus, active bool) ([]*entity.OrderResult, error)
	// Returns a list of orders a customer has tested
	GetCustomerOrders(c *gin.Context, userID string) ([]*entity.OrderResult, error)

	// Creates a new order item : tested
	CreateOrderItem(c *gin.Context, req CreateOrderItemRequest) (*entity.OrderItem, error)
	// Updates an order item : tested
	UpdateOrderItem(c *gin.Context, req UpdateOrderItemRequest) (*entity.OrderItem, error)
	// Removes an order item : tested
	DeleteOrderItem(c *gin.Context, orderID string, menuItemID string) error
}

type CreateOrderRequest struct {
	RestaurantID string         `json:"restaurant_id"`
	Items        map[string]int `json:"items"`
	TableNumber  int64          `json:"table_num"`
}

type UpdateOrderRequest struct {
	RestaurantID string             `json:"restaurant_id"`
	Status       entity.OrderStatus `json:"status"`
	Items        map[string]int     `json:"items"`
	TableNumber  int64              `json:"table_num"`
}

type MoveOrderRequest struct {
	RestaurantID   string `json:"restaurant_id"`
	NewTableNumber int64  `json:"new_table_number"`
}

func (m CreateOrderRequest) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.RestaurantID, validation.Required, validation.Match(regexp.MustCompile("^res_[a-zA-Z0-9]{5,}$"))),
		validation.Field(&m.TableNumber, validation.Required),
		validation.Field(&m.Items, validation.Required),
	)
}

func (m UpdateOrderRequest) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.RestaurantID, validation.Required, validation.Match(regexp.MustCompile("^res_[a-zA-Z0-9]{5,}$"))),
		validation.Field(&m.TableNumber, validation.Required),
		validation.Field(&m.Status, validation.Required, validation.By(validateOrderStatus)),
		validation.Field(&m.Items, validation.Required),
	)
}

func (m MoveOrderRequest) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.RestaurantID, validation.Required, validation.Match(regexp.MustCompile("^res_[a-zA-Z0-9]{5,}$"))),
		validation.Field(&m.NewTableNumber, validation.Required),
	)
}

// Validate order status
func validateOrderStatus(value interface{}) error {
	s, _ := value.(entity.OrderStatus)
	status_list := entity.GetOrderStatus()
	var flag = false

	for i := range status_list {
		if status_list[i] == s {
			flag = true
			break
		}
	}

	if flag == false {
		return errors.BadRequest("Order status invalid")
	}
	return nil
}

type CreateOrderItemRequest struct {
	RestaurantID string `json:"restaurant_id"`
	OrderID      string `json:"order_id"`
	MenuItemID   string `json:"item_id"`
	Quantity     int    `json:"quantity"`
}

type UpdateOrderItemRequest = CreateOrderItemRequest

func (m CreateOrderItemRequest) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.OrderID, validation.Required, validation.Match(regexp.MustCompile("^ord_[a-zA-Z0-9]{5,}$"))),
		validation.Field(&m.MenuItemID, validation.Required, validation.Match(regexp.MustCompile("^itm_[a-zA-Z0-9]{5,}$"))),
		validation.Field(&m.Quantity, validation.Required, validation.Min(1)),
	)
}

// TODO validate Items in CreateOrderItemRequest

type service struct {
	repo         Repository
	userRepo     users.Repository
	menuItemRepo menu_items.Repository
	tablesRepo   tables.Repository
}

// NewService creates a order service
func NewService(repo Repository, userRepo users.Repository, menuItemRepo menu_items.Repository, tablesRepo tables.Repository) service {
	return service{repo, userRepo, menuItemRepo, tablesRepo}
}

// Gets an order including items
func (s service) Get(c *gin.Context, orderID string) (*entity.Order, error) {

	order, err := s.repo.Get(c, orderID)

	if err != nil {
		return nil, err
	}

	order_items, err := s.repo.GetOrderItems(c, order.ID)
	if err != nil {
		return nil, err
	}
	order.OrderItems = order_items

	return order, nil
}

func (s service) GetTableOrder(c *gin.Context, res_id string, table_num int64) (*entity.OrderResult, error) {

	order, err := s.repo.GetTableOrder(c, res_id, table_num)
	if order == nil {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	order_items, err := s.repo.GetOrderItemsResult(c, order.ID)
	if err != nil {
		return nil, err
	}
	if len(order_items) > 0 {
		order.OrderItems = order_items
	} else {
		order.OrderItems = []*entity.OrderItemResult{}
	}
	return order, nil
}

func (s service) Create(c *gin.Context, req CreateOrderRequest) (*entity.Order, error) {

	if err := req.Validate(); err != nil {
		return nil, errors.BadRequest(err.Error())
	}

	now := time.Now()
	order_id := entity.GenerateOrderID()
	total_cost := 0.0
	var order_items []*entity.OrderItem

	order := &entity.Order{
		ID:           order_id,
		RestaurantID: req.RestaurantID,
		UserID:       c.GetString("userID"),
		TotalCost:    0.0,
		Status:       entity.OrderStatus("ordered"),
		CreatedAt:    now,
		UpdatedAt:    now,
		OrderItems:   order_items,
		TableNumber:  req.TableNumber,
	}

	if err := s.repo.Create(c, order); err != nil {
		return nil, err
	}

	// create order_item entities for each item

	for item_id, quantity := range req.Items {
		order_item := &entity.OrderItem{
			OrderID:    order_id,
			MenuItemID: item_id,
			Quantity:   quantity,
		}
		if err := s.repo.CreateOrderItem(c, order_item); err != nil {
			return nil, err
		}
		// add to total cost
		menu_item, err := s.menuItemRepo.Get(c, req.RestaurantID, item_id)
		if err != nil {
			return nil, err
		}

		// Update total cost
		total_cost += float64(quantity) * menu_item.Price
		// add to item list
		order_items = append(order_items, order_item)
	}

	// Update order item list
	order, err := s.repo.Get(c, order_id)

	if err != nil {
		return nil, err
	}

	//order.OrderItems = order_items
	order.TotalCost = total_cost
	order.OrderItems = order_items
	err = s.repo.Update(c, order)
	if err != nil {
		return nil, err
	}

	// Set the table to taken when an order is created
	table, err := s.tablesRepo.GetTableByNumResID(c, order.TableNumber, order.RestaurantID)
	if table == nil {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	table.Status = entity.StatusType("Taken")
	if err := s.tablesRepo.Update(c, table); err != nil {
		return nil, err
	}

	fmt.Println(order)

	return order, nil
}

func (s service) Update(c *gin.Context, orderID string, req UpdateOrderRequest) (*entity.Order, error) {

	order, err := s.repo.Get(c, orderID)

	if err != nil {
		return nil, err
	}

	err = auth.IsManagerOf(c, order.RestaurantID)

	if err != nil {
		return nil, err
	}

	if err := req.Validate(); err != nil {
		return nil, errors.BadRequest(err.Error())
	}

	order.UpdatedAt = time.Now()
	//order.TableID = req.TableID

	if err := s.repo.Update(c, order); err != nil {
		return nil, err
	}

	return order, nil
}

func (s service) MoveOrder(c *gin.Context, orderID string, req MoveOrderRequest) (*entity.Order, error) {
	order, err := s.repo.Get(c, orderID)

	if err != nil {
		return nil, err
	}

	err = auth.IsManagerOf(c, order.RestaurantID)

	if err != nil {
		return nil, err
	}

	if err := req.Validate(); err != nil {
		return nil, errors.BadRequest(err.Error())
	}

	oldTable, err := s.tablesRepo.GetTableByNumResID(c, order.TableNumber, order.RestaurantID)
	if oldTable == nil {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	oldTable.Status = entity.StatusType("Free")
	if err := s.tablesRepo.Update(c, oldTable); err != nil {
		return nil, err
	}

	order.UpdatedAt = time.Now()
	order.TableNumber = req.NewTableNumber

	if err := s.repo.Update(c, order); err != nil {
		return nil, err
	}

	newTable, err := s.tablesRepo.GetTableByNumResID(c, req.NewTableNumber, order.RestaurantID)
	if newTable == nil {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	newTable.Status = entity.StatusType("Taken")
	if err := s.tablesRepo.Update(c, newTable); err != nil {
		return nil, err
	}

	return order, nil
}

func (s service) SetOrderStatus(c *gin.Context, orderID string, status entity.OrderStatus) (*entity.Order, error) {

	// Get order to find restaurant id
	order, err := s.repo.Get(c, orderID)

	if err != nil {
		return nil, err
	}

	err = auth.IsManagerOf(c, order.RestaurantID)

	if err != nil {
		return nil, err
	}

	var flag bool

	switch status {
	case "served":

		if order.Status == "served" {
			return nil, errors.BadRequest("Order is already served")
		}

		if order.Status == "completed" {
			return nil, errors.BadRequest("Order is already completed.")
		}

		if order.Status == "canceled" {
			return nil, errors.BadRequest("Order is canceled, cannot modify status.")
		}

		if order.Status == "ordered" {
			flag = true
		}

	case "completed":

		if order.Status == "ordered" {
			return nil, errors.BadRequest("Order is not yet served, please serve first")
		}

		if order.Status == "completed" {
			return nil, errors.BadRequest("Order is already completed.")
		}

		if order.Status == "canceled" {
			return nil, errors.BadRequest("Order is canceled, cannot modify status.")
		}

		if order.Status == "served" {
			flag = true
		}

	case "canceled":

		if order.Status == "completed" {
			return nil, errors.BadRequest("Cannot cancel a completed order")
		}

		if order.Status == "canceled" {
			return nil, errors.BadRequest("Order is canceled, cannot modify status.")
		}

		flag = true
	default:
		flag = false
	}

	if flag == true {
		// update order status
		order.UpdatedAt = time.Now()
		order.Status = status
		if err := s.repo.Update(c, order); err != nil {
			return nil, err
		}
		table, err := s.tablesRepo.GetTableByNumResID(c, order.TableNumber, order.RestaurantID)
		if table == nil {
			return nil, nil
		}
		if err != nil {
			return nil, err
		}
		if order.Status == "canceled" {
			table.Status = entity.StatusType("Free")
		}
		if order.Status == "completed" {
			table.Status = entity.StatusType("Free")
		}
		if err := s.tablesRepo.Update(c, table); err != nil {
			return nil, err
		}
	}

	return order, nil
}

func (s service) GetRestaurantOrders(c *gin.Context, restaurantID string, status entity.OrderStatus, active bool) ([]*entity.OrderResult, error) {

	err := auth.IsManagerOf(c, restaurantID)

	if err != nil {
		return nil, err
	}

	orders, err := s.repo.GetRestaurantOrders(c, restaurantID, status, active)

	if err != nil {
		return nil, err
	}

	for _, order := range orders {
		order_items, err := s.repo.GetOrderItemsResult(c, order.ID)
		if err != nil {
			return nil, err
		}
		if len(order_items) > 0 {
			order.OrderItems = order_items
		} else {
			order.OrderItems = []*entity.OrderItemResult{}
		}
	}
	return orders, nil
}

func (s service) GetCustomerOrders(c *gin.Context, customerID string) ([]*entity.OrderResult, error) {

	orders, err := s.repo.GetCustomerOrders(c, customerID)

	if err != nil {
		return nil, err
	}

	for _, order := range orders {
		order_items, err := s.repo.GetOrderItemsResult(c, order.ID)
		if err != nil {
			return nil, err
		}
		order.OrderItems = order_items
	}

	return orders, nil
}

func (s service) CreateOrderItem(c *gin.Context, req CreateOrderItemRequest) (*entity.OrderItem, error) {

	order, err := s.repo.Get(c, req.OrderID)

	if err != nil {
		return nil, err
	}

	err = auth.IsManagerOf(c, order.RestaurantID)

	if err != nil {
		return nil, err
	}

	if err := req.Validate(); err != nil {
		return nil, errors.BadRequest(err.Error())
	}

	// check if item exists in order already
	order_items, err := s.repo.GetOrderItems(c, req.OrderID)
	if err != nil {
		return nil, err
	}

	for _, item := range order_items {
		if item.MenuItemID == req.MenuItemID {
			// Increase quantity
			req.Quantity += item.Quantity
			return s.UpdateOrderItem(c, req)
		}
	}

	menu_item, err := s.menuItemRepo.Get(c, order.RestaurantID, req.MenuItemID)
	if err != nil {
		return nil, err
	}

	// Create order item
	order_item := &entity.OrderItem{
		OrderID:    req.OrderID,
		MenuItemID: req.MenuItemID,
		Quantity:   req.Quantity,
	}

	err = s.repo.CreateOrderItem(c, order_item)
	if err != nil {
		return nil, err
	}

	// Update order
	order.UpdatedAt = time.Now()
	order.TotalCost += menu_item.Price * float64(req.Quantity)
	err = s.repo.Update(c, order)
	if err != nil {
		return nil, err
	}

	table, err := s.tablesRepo.GetTableByNumResID(c, order.TableNumber, order.RestaurantID)
	if table == nil {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	if table.Status == entity.StatusType("Free") {
		table.Status = entity.StatusType("Taken")
		if err := s.tablesRepo.Update(c, table); err != nil {
			return nil, err
		}
	}

	return order_item, nil
}

func (s service) UpdateOrderItem(c *gin.Context, req UpdateOrderItemRequest) (*entity.OrderItem, error) {

	order, err := s.repo.Get(c, req.OrderID)

	if err != nil {
		return nil, err
	}

	err = auth.IsManagerOf(c, order.RestaurantID)

	if err != nil {
		return nil, err
	}

	err = req.Validate()
	if err != nil {
		return nil, errors.BadRequest(err.Error())
	}

	order_item, err := s.repo.GetOrderItem(c, req.OrderID, req.MenuItemID)
	if err != nil {
		return nil, err
	}

	if order_item.Quantity != req.Quantity {

		old_quantity := order_item.Quantity

		// Get menu item to update total cost laster
		menu_item, err := s.menuItemRepo.Get(c, order.RestaurantID, req.MenuItemID)
		if err != nil {
			return nil, err
		}

		// Update quantity
		order_item.Quantity = req.Quantity
		err = s.repo.UpdateOrderItem(c, order_item)
		if err != nil {
			return nil, err
		}

		// update order updated time and total price
		order, err := s.repo.Get(c, order_item.OrderID)
		if err != nil {
			return nil, err
		}

		// Update order detais
		order.TotalCost = order.TotalCost + menu_item.Price*(float64(req.Quantity)-float64(old_quantity))
		order.UpdatedAt = time.Now()
		err = s.repo.Update(c, order)
		if err != nil {
			return nil, err
		}
	}

	return order_item, nil
}

func (s service) DeleteOrderItem(c *gin.Context, orderID string, menuItemID string) error {

	order, err := s.repo.Get(c, orderID)

	if err != nil {
		return err
	}

	err = auth.IsManagerOf(c, order.RestaurantID)

	if err != nil {
		return err
	}

	order_item, err := s.repo.GetOrderItem(c, orderID, menuItemID)
	if err != nil {
		return err
	}

	// Get menu item for current quantity
	// Get menu item for unit price
	menu_item, err := s.menuItemRepo.Get(c, order.RestaurantID, menuItemID)
	if err != nil {
		return err
	}

	// delete item
	err = s.repo.DeleteOrderItem(c, orderID, menuItemID)
	if err != nil {
		return err
	}

	// update order
	quantity := order_item.Quantity
	order.UpdatedAt = time.Now()
	order.TotalCost -= float64(quantity) * menu_item.Price

	err = s.repo.Update(c, order)
	if err != nil {
		return err
	}

	return nil
}
