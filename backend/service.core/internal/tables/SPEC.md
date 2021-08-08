# Menu Items API spec

Given 
1) each restaurant has more than 0 tables
2) each table is attached to the restaurant


### Tables
** GET /restaurants/:res_id/tables
List all tables for restaurant :res_id
Used by:
    - Restaurant admin

** GET /restaurants/:res_id/tables/:tlb_id
One table's details
Used by:
    - Restaurant admin

** Post /restaurants/:res_id/tables
Create a new table for restaurant :res_id
Used by:
    - Restaurant admin

** POST  /restaurants/:res_id/tables/:tlb_id/occupy
Update a table status to occupied
Used by:
    - Restaurant admin

** POST  /restaurants/:res_id/tables/:tlb_id/free
Update a table status to free
Used by:
    - Restaurant admin

** Delete /restaurants/:res_id/tables/:tlb_id
Delete a table from the restaurant


Table entity
{
    id (int) PK
    restaurant_id PK
    status (status)
}
