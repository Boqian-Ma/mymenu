
BEGIN;

-- * ===== CORE SERVICE ===== *
DROP TYPE IF EXISTS allergy_type;
CREATE TYPE allergy_type AS ENUM ('', 'Nuts', 'Sea Food', 'Shellfish', 'Dairy', 'Eggs');


CREATE TABLE IF NOT EXISTS users (
    id text check (id ~ '^usr_[a-zA-Z0-9]{5,}$'),
    created_at timestamp not null,
    updated_at timestamp not null,

    allergy allergy_type,

    account_type text not null,
    email text unique,
    name text not null,

    hashed_password text not null,
    salt text not null,

    primary key(id)
);

CREATE UNIQUE INDEX idx_users_email ON users (email);

CREATE TABLE IF NOT EXISTS sessions (
    token text not null,
    user_id text references users (id),

    primary key (token, user_id)
);

CREATE UNIQUE INDEX idx_sessions_token ON sessions (token);


INSERT INTO USERS VALUES ('usr_GWc7JuwWCd', '2021-07-19 11:00:38.635435', '2021-07-19 11:00:38.635435', '', 'manager', 'manager@example.com', 'manager1', 'cf54e9d9da8e2a41d2ab21b7364045ce27b4e5ceaf29bc61f4fb96d5198ca7f7', '4496aeb3-e635-4c21-92ea-ac79e520a407');
INSERT INTO USERS VALUES ('usr_zyutoK4dUd', '2021-07-19 11:00:38.635435', '2021-07-19 11:00:38.635435', 'Nuts','customer', 'customer@example.com', 'customer1', '59333debb7db19aaf95e87d63f6a0f12721b8393275123b03cc3a535e39623f9', 'e2657876-652b-406b-ab67-d5bf3bd8fea6');


COMMIT;
