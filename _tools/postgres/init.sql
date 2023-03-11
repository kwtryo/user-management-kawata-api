CREATE TABLE accounts (
    id serial PRIMARY KEY,
    fullname VARCHAR(100),
    email text UNIQUE,
    phone VARCHAR(100),
    date TIMESTAMP NOT NULL
);
