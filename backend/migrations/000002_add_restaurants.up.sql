
BEGIN;

-- * ===== CORE SERVICE ===== *
DROP TYPE IF EXISTS cuisine_type;
CREATE TYPE cuisine_type AS ENUM ('Asian', 'Indian', 'European', 'Mediterranean', 'North American', 'South American');


CREATE TABLE IF NOT EXISTS restaurants (
    id text check (id ~ '^res_[a-zA-Z0-9]{5,}$'),
    created_at timestamp not null,
    updated_at timestamp not null,

    name text not null,
    type text not null,

    cuisine cuisine_type not null,

    location text not null,
    email text,
    phone text,
    website text,
    business_hours text,
    file text,

    primary key(id)
);

CREATE INDEX idx_restaurants_type ON restaurants (type);
CREATE INDEX idx_restaurants_location ON restaurants (location);

CREATE TABLE IF NOT EXISTS restaurant_members (
    user_id text references users (id),
    restaurant_id text references restaurants (id),

    primary key (user_id, restaurant_id)
);

-- * ===== Restaurant Asian ===== *

INSERT INTO RESTAURANTS VALUES ('res_C6lNdP8uZu', '2021-07-19 11:21:08.516376', '2021-07-19 11:21:08.516376', 'Quay', 'bar', 'Asian', 'Sydney', 'booking@quay.com.au', '02 9251 5600', 'www.quay.com.au', '', 'quay.jpg');
INSERT INTO restaurant_members values ('usr_GWc7JuwWCd', 'res_C6lNdP8uZu');

INSERT INTO RESTAURANTS VALUES ('res_C6lNdP8uZa', '2021-07-19 11:21:08.516376', '2021-07-19 11:21:08.516376', 'Sokyo', 'diner','Asian', 'Sydney',  'STARRESERVATIONS@STAR.COM.AU', '02 9777 9000', 'https://www.star.com.au/sydney/eat-and-drink/signature-dining/sokyo', '', 'sokyo.jpg');
INSERT INTO restaurant_members values ('usr_GWc7JuwWCd', 'res_C6lNdP8uZa');

-- * ===== Restaurant European ===== *

-- INSERT INTO RESTAURANTS VALUES ('res_C6lNdP8uZ1', '2021-07-19 11:21:08.516376', '2021-07-19 11:21:08.516376', 'The Little Cup and Saucer', 'Cafe', 'European', 'Sydney', 'info@thelittlecupandsaucer.com.au', '02 9591 8886', 'https://www.facebook.com/thelittlecupandsaucer/', '', 'yeet');
-- INSERT INTO restaurant_members values ('usr_GWc7JuwWCd', 'res_C6lNdP8uZ1');

-- INSERT INTO RESTAURANTS VALUES ('res_C6lNdP8uZ2', '2021-07-19 11:21:08.516376', '2021-07-19 11:21:08.516376', 'Pancakes On The Rocks', 'Diner','European', '22 Playfair St, The Rocks NSW 2000',  'STARRESERVATIONS@STAR.COM.AU', '02 9247 6371', 'https://pancakesontherocks.com.au/', '', 'yeet');
-- INSERT INTO restaurant_members values ('usr_GWc7JuwWCd', 'res_C6lNdP8uZ2');


-- INSERT INTO RESTAURANTS VALUES ('res_C6lNdP8uZ3', '2021-07-19 11:21:08.516376', '2021-07-19 11:21:08.516376', 'Curry King Express', 'Diner', 'Indian', 'Sydney', 'info@thelittlecupandsaucer.com.au', '02 9591 8886', 'https://www.facebook.com/thelittlecupandsaucer/', '', 'yeet');
-- INSERT INTO restaurant_members values ('usr_GWc7JuwWCd', 'res_C6lNdP8uZ3');

-- INSERT INTO RESTAURANTS VALUES ('res_C6lNdP8uZ4', '2021-07-19 11:21:08.516376', '2021-07-19 11:21:08.516376', 'Pancakes On The Rocks', 'Diner','Indian', '22 Playfair St, The Rocks NSW 2000',  'STARRESERVATIONS@STAR.COM.AU', '02 9247 6371', 'https://pancakesontherocks.com.au/', '', 'yeet');
-- INSERT INTO restaurant_members values ('usr_GWc7JuwWCd', 'res_C6lNdP8uZ4');

-- COPY RESTAURANTS FROM '/home/lubuntu/capstoneproject-comp3900-w16a-jamar/backend/migrations/restaurant_members.csv' DELIMITER ',' CSV HEADER;

-- COPY restaurant_members FROM 'restaurants.csv' DELIMITER ',' CSV HEADER;



COMMIT;
