CREATE TABLE "users"
(
    "ID"        serial PRIMARY KEY NOT NULL,
    "FirstName" text               NOT NULL,
    "LastName"  text               NOT NULL,
    "Age"       int                NOT NULL
);