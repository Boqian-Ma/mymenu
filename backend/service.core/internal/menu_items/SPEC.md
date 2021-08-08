
# Menu Items API spce

Given
1) each restaurant has exactly 1 menu
2) menu items have to be attached to the restaurant

We can remove the explicit concept of a menu entirely from the backend
- Every restaurant will have a list of menu items
- Each item will contain an 'available' field (true or false)
- If a manager calls get /restaurants/:id/menu_items
    - It returns all items
- If a customer callls get /restaurants/:id/menu
    - It returns only 'available' items

- Item 'category' will just be a string field with the category name in lowercase
    - Need to validate this on item creation & deletion to ensure we don't have data integrity issues

## PHASE 1
### MENU ITEMS
**GET /restaurants/:id/menu_items**
List all menu items available for restaurant :id
- Can only be used by managers of the restaurant (customers get a 403)

**POST /restaurants/:id/menu_items**
Create a new menu item for the restaurant :id
- Can only be used by managers of the restaurant (customers get a 403)

**PUT /restaurants/:id/menu_items/:id**
Update a menu item, this api will also be used to mark items as 'available'
- Can only be used by managers of the restaurant (customers get a 403)

**DELETE /restaurants/:id/menu_items/:id**
Removes a menu item from a restaurant
- Can only be used by managers of the restaurant (customers get a 403)

@Adamn changed my mind ur right I reckon we should build this - good practice!
**GET /restaurants/:id/menu_items/:id**
returns a specific menu item
- Can only be used by managers of the restaurant (customers get a 403)

## PHASE 2
### MENU
**GET /restaurants/:id/menu**
Returns a list of all the 'available' menu items for a specific restaurant
Intended to be used by customers 
