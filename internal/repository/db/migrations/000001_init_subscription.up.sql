CREATE TABLE IF NOT EXISTS subscription (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    price INTEGER  NOT NULL,
    user_uuid UUID NOT NULL,
    date_start DATE NOT NULL,
    date_end DATE
);