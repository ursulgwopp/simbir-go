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

CREATE TABLE rents (
    id SERIAL PRIMARY KEY,
    transport_id INTEGER,
    user_id INTEGER,
    time_start TIMESTAMPTZ,
    time_end TIMESTAMPTZ DEFAULT '1970-01-01 00:00:00',
    price_of_unit DECIMAL,
    price_type TEXT,
    final_price DECIMAL DEFAULT 0,
    is_active BOOLEAN DEFAULT TRUE
);

CREATE TABLE blacklist (
    token TEXT
);

INSERT INTO accounts (username, hash_password, is_admin) VALUES ('admin', '313431736169693939787833692f3132313540d033e22ae348aeb5660fc2140aec35850c4da997', TRUE);

INSERT INTO transports (owner_id, can_be_rented, transport_type, model, color, identifier, description, latitude, longitude, minute_price, day_price) VALUES (1, TRUE, 'Car', 'Mercedes-Benz GLE350 Coupe AMG', 'black', 'в001ор178', '', 59.84, 30.25, 25, 25000);
INSERT INTO transports (owner_id, can_be_rented, transport_type, model, color, identifier, description, latitude, longitude, minute_price, day_price) VALUES (1, TRUE, 'Car', 'Mercedes-Benz E200', 'black', 'в002ор178', '', 59.84, 30.25, 18, 15000);
INSERT INTO transports (owner_id, can_be_rented, transport_type, model, color, identifier, description, latitude, longitude, minute_price, day_price) VALUES (1, TRUE, 'Car', 'Mercedes-Benz CLA180', 'white', 'в003ор178', '', 59.84, 30.25, 18, 15000);
-- INSERT INTO rents (transport_id, user_id, price_of_unit, price_type) VALUES (1, 1, 500, 'Days');
-- INSERT INTO rents (transport_id, user_id, price_of_unit, price_type) VALUES (1, 1, 500, 'Days');
-- INSERT INTO rents (transport_id, user_id, price_of_unit, price_type) VALUES (2, 1, 500, 'Days');