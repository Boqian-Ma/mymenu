BEGIN;

-- * ===== CORE SERVICE ===== *

CREATE TABLE IF NOT EXISTS categories (
    id text check (id ~ '^cat_[a-zA-Z0-9]{5,}$'),
    name text not null,

    created_at timestamp not null,
    updated_at timestamp not null,

    restaurant_id text references restaurants (id),
    primary key(id)
);

CREATE TABLE IF NOT EXISTS menu_items (
    id text check (id ~ '^itm_[a-zA-Z0-9]{5,}$'),

    name text not null,
    description text not null,
    price decimal not null default 0.00,
    allergy allergy_type,

    created_at timestamp not null,
    updated_at timestamp not null,

    is_special boolean not null default FALSE,
    is_menu boolean not null default FALSE,

    file text,

    category_id text references categories (id),
    restaurant_id text references restaurants (id),
    primary key(id)
);

-- * ===== Restaurant 1 Quay ===== *
INSERT INTO CATEGORIES VALUES ('cat_Rtnp0C3aWs', 'Main', '2021-07-19 11:38:06.653134', '2021-07-19 11:38:06.653134', 'res_C6lNdP8uZu');
INSERT INTO CATEGORIES VALUES ('cat_LZOrhEhDFK', 'Desserts', '2021-07-19 11:38:06.653134', '2021-07-19 11:38:06.653134', 'res_C6lNdP8uZu');
INSERT INTO CATEGORIES VALUES ('cat_6kgVQOu3aX', 'Drinks', '2021-07-19 11:38:06.653134', '2021-07-19 11:38:06.653134', 'res_C6lNdP8uZu');

INSERT INTO MENU_ITEMS VALUES ('itm_YCwMQJwQf2','Smoked eel','green walnuts, Oscietra caviar sea cucumber crackling', 999, 'Nuts', '2021-07-19 11:45:10.645704','2021-07-19 11:45:10.645704', 't', 't', '48.jpg','cat_Rtnp0C3aWs','res_C6lNdP8uZu');
INSERT INTO MENU_ITEMS VALUES ('itm_YCwMQJwQf3','Tasmanian rock lobster','ginger scented milk curd super chicken broth, sudachi citrus', 3099 , 'Shellfish','2021-07-19 11:45:10.645704','2021-07-19 11:45:10.645704', 'f', 't','38.jpg', 'cat_Rtnp0C3aWs','res_C6lNdP8uZu');
INSERT INTO MENU_ITEMS VALUES ('itm_YCwMQJwQf4','Bone marrow noodles','koji butter, Southern squid', 1699 , 'Nuts','2021-07-19 11:45:10.645704','2021-07-19 11:45:10.645704', 't', 't', '39.jpg','cat_Rtnp0C3aWs','res_C6lNdP8uZu');

INSERT INTO MENU_ITEMS VALUES ('itm_YCwMQJwQf5','Nikka Caramel Macchiato','Coffee ice cream, cacao nibs, whisky foam', 1499 , 'Eggs','2021-07-19 11:45:10.645704','2021-07-19 11:45:10.645704', 't', 't', '40.jpg','cat_LZOrhEhDFK','res_C6lNdP8uZu');
INSERT INTO MENU_ITEMS VALUES ('itm_YCwMQJwQf6','Tofu Cheesecake','Toasted soybean cookie, yoghurt & raspberry sorbet, hibiscus flower', 1699 , 'Sea Food','2021-07-19 11:45:10.645704','2021-07-19 11:45:10.645704', 'f', 't','41.jpg', 'cat_LZOrhEhDFK','res_C6lNdP8uZu');

INSERT INTO MENU_ITEMS VALUES ('itm_YCwMQJwQf7','Bone marrow noodles','Yuzu sake, elderflower, grapefruit', 899 , 'Shellfish','2021-07-19 11:45:10.645704','2021-07-19 11:45:10.645704', 'f', 't','42.jpg', 'cat_6kgVQOu3aX','res_C6lNdP8uZu');
INSERT INTO MENU_ITEMS VALUES ('itm_YCwMQJwQf8','Sake Sbagliato','Sparkling wine, Lillet Blanc, Campari', 899, 'Dairy','2021-07-19 11:45:10.645704','2021-07-19 11:45:10.645704', 't', 't', '43.jpg','cat_6kgVQOu3aX','res_C6lNdP8uZu');


-- * ===== Restaurant 2 Sokyo ===== *
INSERT INTO CATEGORIES VALUES ('cat_Rtnp0C3aW1', 'Breakfast', '2021-07-19 11:38:06.653134', '2021-07-19 11:38:06.653134', 'res_C6lNdP8uZa');
INSERT INTO CATEGORIES VALUES ('cat_LZOrhEhDF2', 'Lunch', '2021-07-19 11:38:06.653134', '2021-07-19 11:38:06.653134', 'res_C6lNdP8uZa');

INSERT INTO MENU_ITEMS VALUES ('itm_YCwMQJwQa2','Chilli crab omelette','Spanner crab, seaweed rice, Sambal butter', 1999, 'Nuts', '2021-07-19 11:45:10.645704','2021-07-19 11:45:10.645704', 'f', 't', '44.jpg','cat_Rtnp0C3aW1','res_C6lNdP8uZa');
INSERT INTO MENU_ITEMS VALUES ('itm_YCwMQJwQa3','Yuzu Waffle','Yuzu marmalade, raspberries, whipped cream', 3099 , 'Nuts', '2021-07-19 11:45:10.645704','2021-07-19 11:45:10.645704', 'f', 't', '45.jpg','cat_Rtnp0C3aW1','res_C6lNdP8uZa');

INSERT INTO MENU_ITEMS VALUES ('itm_YCwMQJwQa4','Sashimi Platter','Chef’s selection 24-piece sashimi', 6969, 'Dairy','2021-07-19 11:45:10.645704','2021-07-19 11:45:10.645704', 'f', 't', '46.jpg','cat_LZOrhEhDF2','res_C6lNdP8uZa');
INSERT INTO MENU_ITEMS VALUES ('itm_YCwMQJwQa5','Dengakuman','Miso glazed toothfish, Japanese salsa, pickled cucumber', 6099 , 'Eggs','2021-07-19 11:45:10.645704','2021-07-19 11:45:10.645704', 'f', 't', '47.jpg','cat_LZOrhEhDF2','res_C6lNdP8uZa');



-- -- * ===== Restaurant 3 The Little Cup and Saucer ===== *

-- INSERT INTO CATEGORIES VALUES ('cat_Rtnp0C3aW3', 'Breakfast', '2021-07-19 11:38:06.653134', '2021-07-19 11:38:06.653134', 'res_C6lNdP8uZ1');
-- INSERT INTO CATEGORIES VALUES ('cat_LZOrhEhDF4', 'Desserts', '2021-07-19 11:38:06.653134', '2021-07-19 11:38:06.653134', 'res_C6lNdP8uZ1');

-- INSERT INTO MENU_ITEMS VALUES ('itm_YCwMQJwQ12','Eggs','green walnuts, Oscietra caviar sea cucumber crackling', 9,'2021-07-19 11:45:10.645704','2021-07-19 11:45:10.645704', 'f', 't', '','cat_Rtnp0C3aW3','res_C6lNdP8uZ1');
-- INSERT INTO MENU_ITEMS VALUES ('itm_YCwMQJwQ23','French toast','ginger scented milk curd super chicken broth, sudachi citrus', 30 ,'2021-07-19 11:45:10.645704','2021-07-19 11:45:10.645704', 'f', 't','', 'cat_Rtnp0C3aW3','res_C6lNdP8uZ1');
-- INSERT INTO MENU_ITEMS VALUES ('itm_YCwMQJwQ34','Hot Cereal','koji butter, Southern squid', 19 ,'2021-07-19 11:45:10.645704','2021-07-19 11:45:10.645704', 't', 't', '','cat_Rtnp0C3aW3','res_C6lNdP8uZ1');

-- INSERT INTO MENU_ITEMS VALUES ('itm_YCwMQJwQa8','Coffee','Black coffee', 8 ,'2021-07-19 11:45:10.645704','2021-07-19 11:45:10.645704', 't', 't', '','cat_6kgVQOu3F4','res_C6lNdP8uZ1');
-- INSERT INTO MENU_ITEMS VALUES ('itm_YCwMQJwQa8','Latte','Columbia', 8 ,'2021-07-19 11:45:10.645704','2021-07-19 11:45:10.645704', 'f', 't', '','cat_6kgVQOu3F4','res_C6lNdP8uZ1');


-- -- * ===== Restaurant 4 Pancakes On The Rocks ===== *

-- INSERT INTO CATEGORIES VALUES ('cat_Rtnp0C3aa4', 'Main', '2021-07-19 11:38:06.653134', '2021-07-19 11:38:06.653134', 'res_C6lNdP8uZ2');
-- INSERT INTO CATEGORIES VALUES ('cat_LZOrhEhDb5', 'Desserts', '2021-07-19 11:38:06.653134', '2021-07-19 11:38:06.653134', 'res_C6lNdP8uZ2');

-- INSERT INTO MENU_ITEMS VALUES ('itm_YCwMQJwQ12','Eggs','green walnuts, Oscietra caviar sea cucumber crackling', 9,'2021-07-19 11:45:10.645704','2021-07-19 11:45:10.645704', 'f', 't', '','cat_Rtnp0C3aa4','res_C6lNdP8uZ2');
-- INSERT INTO MENU_ITEMS VALUES ('itm_YCwMQJwQ23','French toast','ginger scented milk curd super chicken broth, sudachi citrus', 30 ,'2021-07-19 11:45:10.645704','2021-07-19 11:45:10.645704', 'f', 't','', 'cat_Rtnp0C3aa4','res_C6lNdP8uZ2');
-- INSERT INTO MENU_ITEMS VALUES ('itm_YCwMQJwQ34','Hot Cereal','koji butter, Southern squid', 19 ,'2021-07-19 11:45:10.645704','2021-07-19 11:45:10.645704', 't', 't', '','cat_Rtnp0C3aa4','res_C6lNdP8uZ2');

-- INSERT INTO MENU_ITEMS VALUES ('itm_YCwMQJwQa8','Coffee','Black coffee', 8 ,'2021-07-19 11:45:10.645704','2021-07-19 11:45:10.645704', 't', 't', '','cat_LZOrhEhDb5','res_C6lNdP8uZ2');
-- INSERT INTO MENU_ITEMS VALUES ('itm_YCwMQJwQa8','Latte','Columbia', 8 ,'2021-07-19 11:45:10.645704','2021-07-19 11:45:10.645704', 'f', 't', '','cat_LZOrhEhDb5','res_C6lNdP8uZ2');




COMMIT;
