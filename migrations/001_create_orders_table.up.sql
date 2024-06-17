CREATE TABLE orders (
    id SERIAL PRIMARY KEY,
    description TEXT,
    price NUMERIC(10, 2),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
