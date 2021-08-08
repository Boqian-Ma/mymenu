package entity

// Represents a relationship between an order and a menu item
type OrderItem struct {
	OrderID    string `json:"-" db:"order_id"`
	MenuItemID string `json:"item_id" db:"item_id"`

	Quantity int `json:"quantity" db:"quantity"`
}

type OrderItemResult struct {
	OrderID    string `json:"-" db:"order_id"`
	MenuItemID string `json:"item_id" db:"item_id"`

	Quantity      int    `json:"quantity" db:"quantity"`
	MenuItemName  string `json:"item_name" db:"item_name"`
	MenuItemPrice int    `json:"item_price" db:"item_price"`
}
