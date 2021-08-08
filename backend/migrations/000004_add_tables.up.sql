BEGIN;

-- * ===== CORE SERVICE ===== *

DROP TYPE IF EXISTS statustype;
CREATE TYPE statustype AS ENUM ('Taken', 'Free');

CREATE TABLE IF NOT EXISTS tables (
    table_num integer check (table_num > 0),
    created_at timestamp not null,
    updated_at timestamp not null,

    restaurant_id text references restaurants (id),
    status statustype not null,
    num_seats integer check (num_seats > 0),

    primary key(table_num, restaurant_id )
);


INSERT INTO TABLES VALUES (1, '2021-07-20 14:34:13.205848', '2021-07-20 14:34:13.205848', 'res_C6lNdP8uZa', 'Free', 2);
INSERT INTO TABLES VALUES (2, '2021-07-20 14:34:13.205848', '2021-07-20 14:34:13.205848', 'res_C6lNdP8uZa', 'Free', 2);
INSERT INTO TABLES VALUES (3, '2021-07-20 14:34:13.205848', '2021-07-20 14:34:13.205848', 'res_C6lNdP8uZa', 'Free', 2);

INSERT INTO TABLES VALUES (1, '2021-07-20 14:34:13.205848', '2021-07-20 14:34:13.205848', 'res_C6lNdP8uZu', 'Free', 2);
INSERT INTO TABLES VALUES (2, '2021-07-20 14:34:13.205848', '2021-07-20 14:34:13.205848', 'res_C6lNdP8uZu', 'Free', 2);
INSERT INTO TABLES VALUES (3, '2021-07-20 14:34:13.205848', '2021-07-20 14:34:13.205848', 'res_C6lNdP8uZu', 'Free', 2);

COMMIT;