CREATE TABLE accounts (
    id SERIAL PRIMARY KEY,
    username TEXT UNIQUE,
    hash_password TEXT,
    balance DECIMAL DEFAULT 0,
    is_admin BOOLEAN DEFAULT FALSE
);

CREATE TABLE blacklist (
    token TEXT
);