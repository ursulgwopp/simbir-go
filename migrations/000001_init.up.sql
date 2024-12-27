CREATE TABLE accounts (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) UNIQUE NOT NULL,
    hash_password VARCHAR(255) NOT NULL,
    balance DECIMAL DEFAULT 0,
    is_admin BOOLEAN DEFAULT FALSE
);

CREATE TABLE transports (
    id SERIAL PRIMARY KEY,
    owner_id INTEGER REFERENCES accounts(id) ON DELETE CASCADE,
    can_be_rented BOOLEAN NOT NULL,
    transport_type VARCHAR(255) NOT NULL,
    model VARCHAR(255) NOT NULL,
    color VARCHAR(255) NOT NULL,
    identifier VARCHAR(255) UNIQUE NOT NULL,
    description VARCHAR(255) DEFAULT '',
    latitude DECIMAL NOT NULL,
    longitude DECIMAL NOT NULL,
    minute_price DECIMAL DEFAULT 0,
    day_price DECIMAL DEFAULT 0
);

CREATE TABLE rents (
    id SERIAL PRIMARY KEY,
    transport_id INTEGER REFERENCES transports(id),
    user_id INTEGER REFERENCES accounts(id),
    time_start TIMESTAMPTZ NOT NULL,
    time_end TIMESTAMPTZ DEFAULT '1970-01-01 00:00:00',
    price_of_unit DECIMAL NOT NULL,
    price_type VARCHAR(255) NOT NULL,
    final_price DECIMAL DEFAULT 0,
    is_active BOOLEAN DEFAULT TRUE
);

CREATE TABLE blacklist (
    token TEXT
);

INSERT INTO accounts (username, hash_password, is_admin) VALUES ('admin', '313431736169693939787833692f3132313540d033e22ae348aeb5660fc2140aec35850c4da997', TRUE);

INSERT INTO transports (owner_id, can_be_rented, transport_type, model, color, identifier, description, latitude, longitude, minute_price, day_price) VALUES (1, TRUE, 'Car', 'Mercedes-Benz GLE350 Coupe', 'Black', 'B001OP178', '', 59.84, 30.25, 25, 25000);
INSERT INTO transports (owner_id, can_be_rented, transport_type, model, color, identifier, description, latitude, longitude, minute_price, day_price) VALUES (1, TRUE, 'Car', 'Mercedes-Benz E200', 'Black', 'B002OP178', '', 59.84, 30.25, 18, 15000);
INSERT INTO transports (owner_id, can_be_rented, transport_type, model, color, identifier, description, latitude, longitude, minute_price, day_price) VALUES (1, TRUE, 'Car', 'Mercedes-Benz CLA180 AMG', 'White', 'B003OP178', '', 59.84, 30.25, 18, 15000);