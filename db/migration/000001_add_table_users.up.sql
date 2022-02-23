CREATE TABLE users
(
    "id"         serial PRIMARY KEY NOT NULL,
    "first_name" text               NOT NULL,
    "last_name"  text               NOT NULL,
    "age"        int                NOT NULL
);