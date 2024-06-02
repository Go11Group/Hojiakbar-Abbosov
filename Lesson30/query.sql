CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) NOT NULL,
    email VARCHAR(100) NOT NULL UNIQUE,
    password VARCHAR(100) NOT NULL
);

CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    description TEXT,
    price NUMERIC(10, 2) NOT NULL,
    stock_quantity INT NOT NULL
);

INSERT INTO users (username, email, password) VALUES
('john_doe', 'john@example.com', 'securepassword1'),
('jane_smith', 'jane@example.com', 'securepassword2'),
('alice_jones', 'alice@example.com', 'securepassword3'),
('bob_brown', 'bob@example.com', 'securepassword4'),
('charlie_black', 'charlie@example.com', 'securepassword5');

INSERT INTO products (name, description, price, stock_quantity) VALUES
('Laptop', 'A powerful laptop with 16GB RAM and 512GB SSD', 1200.00, 10),
('Smartphone', 'A smartphone with a 6.5-inch display and 128GB storage', 699.99, 50),
('Headphones', 'Noise-cancelling over-ear headphones', 199.99, 30),
('Smartwatch', 'A smartwatch with fitness tracking features', 149.99, 20),
('Tablet', 'A tablet with a 10-inch display and 64GB storage', 399.99, 15);
