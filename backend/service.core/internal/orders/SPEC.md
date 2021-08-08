
# Orders API

** GET /orders?res_id?u_id?status
Gets a list of orders depend on query parameters

** POST /orders
Creates an order
- Used by:
    - Restaurant manager
    - Customer

** PUT /orders/:order_id
Update details of an order
    - Modify an order
        - Qantity 
- Used by:
    - Restaurant manager

** POST /orders/:order_id/complete
Updates status to completed
- Used by:
    - Restaurant manager

** POST /orders/:order_id/cook
Updates status to cooked
- Used by:
    - Restaurant manager


body:
    {
        res_id
        ....
    }

## Restaurant order
### Orders
** GET /restaurants/:res_id/orders?status
List all orders this restaurant

- status: a list of statuses
    returns a list of matching orders
- Used by:
    - Restaurant manager

** GET /restaurants/:res_id/orders/:order_id
Get details of one order
- Used by:
    - Restaurant manager
    - customer (only accessing their orders)

** PUT /restaurants/:res_id/orders/:order_id
Update details of an order
    - Modify an order
        - Qantity 
- Used by:
    - Restaurant manager

** POST /restaurants/:res_id/orders
Create an order
- Used by:
    - Restaurant manager
    - Customer

** POST /restaurants/:res_id/orders/:order_id/cook
Set the order status to "cooked"
- Used by 
    - Restaurant manager

** POST /restaurants/:res_id/orders/:order_id/complete
Set the order status to "completed"
- Used by 
    - Restaurant manager

## User order history
** GET /users/:u_id/orders
List all order history for user :u_id
- Used by:
    - Customer



Order entity:
{
    id (PK)
    user_id (FK)
    restaurant_id (FK)
    table_id (FK)
    CreatedAt
    ModifiedAt
    status (string)
}

order_item:
{
    order_id
    menu_item_id
    quantity (int, > 0)
}

Status flow
1. Ordered
2. Cooked
3. Complete