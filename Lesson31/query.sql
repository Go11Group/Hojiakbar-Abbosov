CREATE TABLE IF NOT EXISTS product (
    id UUID PRIMARY KEY,
    name VARCHAR,
    category VARCHAR,
    cost INT
);

CREATE INDEX idx_product_name ON product (name);

CREATE INDEX idx_product_category ON product (category);
