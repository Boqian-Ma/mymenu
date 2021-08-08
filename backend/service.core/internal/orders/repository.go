package orders

import (
	"github.com/COMP3900-9900-Capstone-Project/capstoneproject-comp3900-w16a-jamar/pkg/db"
	"github.com/COMP3900-9900-Capstone-Project/capstoneproject-comp3900-w16a-jamar/pkg/entity"
	"github.com/gin-gonic/gin"
)

type Repository interface {
	// Returns details of an order
	Get(c *gin.Context, orderID string) (*entity.Order, error)
	// Return details of order based on res_id and table_num
	GetTableOrder(c *gin.Context, res_id string, table_num int64) (*entity.OrderResult, error)
	// Creates an order
	Create(c *gin.Context, order *entity.Order) error
	// Removes an order
	Delete(c *gin.Context, orderID string) error
	// Updates an order
	Update(c *gin.Context, order *entity.Order) error
	// Returns a list of orders a restaurant has
	GetRestaurantOrders(c *gin.Context, restaurantID string, status entity.OrderStatus, active bool) ([]*entity.OrderResult, error)
	// Returns a list of orders a customer has
	GetCustomerOrders(c *gin.Context, userID string) ([]*entity.OrderResult, error)

	// Gets an order item
	GetOrderItem(c *gin.Context, orderID string, menuItemID string) (*entity.OrderItem, error)
	// Creates a order item relationship
	CreateOrderItem(c *gin.Context, order_item *entity.OrderItem) error
	// Deletes an order item
	DeleteOrderItem(c *gin.Context, orderID string, menuItemID string) error
	// Updates an order item
	UpdateOrderItem(c *gin.Context, order_item *entity.OrderItem) error
	// Gets all items associated to an order
	GetOrderItems(c *gin.Context, orderID string) ([]*entity.OrderItem, error)
	// Gets an order item
	GetOrderItemResult(c *gin.Context, orderID string, menuItemID string) (*entity.OrderItemResult, error)
	// Gets all items associated to an order
	GetOrderItemsResult(c *gin.Context, orderID string) ([]*entity.OrderItemResult, error)
}

type repository struct {
	db db.DB
}

// NewRepository returns a new repostory that can be used to access menu item data
func NewRepository(db db.DB) Repository {
	return repository{db}
}

func (r repository) Get(c *gin.Context, orderID string) (*entity.Order, error) {
	var order entity.Order

	if err := r.db.Select(c, &order, "select * from orders where id=$1", orderID); err != nil {
		return nil, err
	}
	return &order, nil
}

func (r repository) GetTableOrder(c *gin.Context, res_id string, table_num int64) (*entity.OrderResult, error) {
	var order entity.OrderResult

	if err := r.db.Select(c, &order, "select * from orders where restaurant_id=$1 and table_num=$2 and (status='ordered' or status='served')", res_id, table_num); err != nil {
		return nil, err
	}
	return &order, nil
}

func (r repository) Create(c *gin.Context, order *entity.Order) error {
	return r.db.Insert(c, "orders", order)
}

func (r repository) Delete(c *gin.Context, orderID string) error {
	return r.db.Delete(c, "delete from orders where id=$1", orderID)
}

func (r repository) Update(c *gin.Context, order *entity.Order) error {
	return r.db.Update(c, order, "update orders set ... where id=$1", order.ID)
}

func (r repository) GetRestaurantOrders(c *gin.Context, restaurantID string, status entity.OrderStatus, active bool) ([]*entity.OrderResult, error) {

	var orders []*entity.OrderResult

	if status == "" {
		if active {
			if err := r.db.Select(c, &orders, "select * from orders where restaurant_id=$1 and (status='ordered' or status='served') order by table_num", restaurantID); err != nil {
				return nil, err
			}
		} else {
			if err := r.db.Select(c, &orders, "select * from orders where restaurant_id=$1 order by table_num", restaurantID); err != nil {
				return nil, err
			}
		}

	} else {
		if err := r.db.Select(c, &orders, "select * from orders where restaurant_id=$1 and status=$2 order by table_num", restaurantID, status); err != nil {
			return nil, err
		}
	}

	return orders, nil
}

func (r repository) GetCustomerOrders(c *gin.Context, userID string) ([]*entity.OrderResult, error) {
	var orders []*entity.OrderResult

	if err := r.db.Select(c, &orders, "select * from orders where user_id=$1", userID); err != nil {
		return nil, err
	}

	return orders, nil
}

func (r repository) CreateOrderItem(c *gin.Context, order_item *entity.OrderItem) error {
	return r.db.Insert(c, "orders_items", order_item)
}

func (r repository) DeleteOrderItem(c *gin.Context, orderID string, menuItemID string) error {
	return r.db.Delete(c, "delete from orders_items where order_id=$1 and item_id=$2", orderID, menuItemID)
}

func (r repository) UpdateOrderItem(c *gin.Context, order_item *entity.OrderItem) error {
	return r.db.Update(c, order_item, "update orders_items set ... where order_id=$1 and item_id=$2", order_item.OrderID, order_item.MenuItemID)
}

func (r repository) GetOrderItems(c *gin.Context, orderID string) ([]*entity.OrderItem, error) {
	var order_items []*entity.OrderItem

	if err := r.db.Select(c, &order_items, "select * from orders_items where order_id=$1", orderID); err != nil {
		return nil, err
	}
	return order_items, nil
}

func (r repository) GetOrderItem(c *gin.Context, orderID string, menuItemID string) (*entity.OrderItem, error) {
	var order_item entity.OrderItem

	if err := r.db.Select(c, &order_item, "select * from orders_items where order_id=$1 and item_id=$2", orderID, menuItemID); err != nil {
		return nil, err
	}
	return &order_item, nil
}

func (r repository) GetOrderItemsResult(c *gin.Context, orderID string) ([]*entity.OrderItemResult, error) {
	var order_items []*entity.OrderItemResult

	if err := r.db.Select(c, &order_items, "select o.*, i.name as item_name, i.price as item_price from orders_items o join menu_items i on o.item_id=i.id where o.order_id=$1", orderID); err != nil {
		return nil, err
	}
	return order_items, nil
}

func (r repository) GetOrderItemResult(c *gin.Context, orderID string, menuItemID string) (*entity.OrderItemResult, error) {
	var order_item entity.OrderItemResult

	if err := r.db.Select(c, &order_item, "select o.*, i.name as item_name, i.price as item_price from orders_items o join menu_items i on o.item_id=i.id where o.order_id=$1 and o.item_id=$2", orderID, menuItemID); err != nil {
		return nil, err
	}
	return &order_item, nil
}
