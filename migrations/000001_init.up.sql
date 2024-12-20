CREATE TABLE accounts (
    id SERIAL PRIMARY KEY,
    username TEXT UNIQUE,
    hash_password TEXT,
    balance DECIMAL DEFAULT 0,
    is_admin BOOLEAN DEFAULT FALSE
);

CREATE TABLE transports (
    id SERIAL PRIMARY KEY,
    owner_id INTEGER,
    can_be_rented BOOLEAN,
    transport_type TEXT,
    model TEXT,
    color TEXT,
    identifier TEXT,
    description TEXT DEFAULT '',
    latitude DECIMAL,
    longitude DECIMAL,
    minute_price DECIMAL DEFAULT 0,
    day_price DECIMAL DEFAULT 0
);

CREATE TABLE blacklist (
    token TEXT
);

INSERT INTO accounts (username, hash_password, is_admin) VALUES ('admin', '313431736169693939787833692f3132313540d033e22ae348aeb5660fc2140aec35850c4da997', TRUE);