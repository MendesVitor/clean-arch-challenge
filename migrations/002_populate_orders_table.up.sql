-- migrations/002_populate_orders_table.up.sql

INSERT INTO orders (description, price, created_at) VALUES 
('Order 1', 19.99, NOW()),
('Order 2', 29.99, NOW()),
('Order 3', 39.99, NOW()),
('Order 4', 49.99, NOW()),
('Order 5', 59.99, NOW());
