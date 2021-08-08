BEGIN;

-- * ===== CORE SERVICE ===== *

CREATE TABLE IF NOT EXISTS orders (
    id text check (id ~ '^ord_[a-zA-Z0-9]{5,}$'),

    total_cost decimal not null default 0.00 check (total_cost >= 0.0),

    created_at timestamp not null,
    updated_at timestamp not null,

    status text not null,

    table_num integer check (table_num > 0),

    restaurant_id text references restaurants (id),
    user_id text references users (id),
    foreign key(table_num, restaurant_id) references tables(table_num, restaurant_id),

    primary key(id)

);

CREATE TABLE IF NOT EXISTS orders_items (
    order_id text references orders (id) not null,
    item_id text references menu_items (id) not null,

    quantity integer not null check (quantity > 0),

    primary key (order_id, item_id)
);

ALTER TABLE orders_items DROP CONSTRAINT "orders_items_item_id_fkey";

COMMIT;