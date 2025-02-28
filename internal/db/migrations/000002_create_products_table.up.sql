CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    supplier_id INT REFERENCES users(id) NOT NULL,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    category VARCHAR(255) NOT NULL,
    price NUMERIC(10, 2) NOT NULL,
    stock INT DEFAULT 0 NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE INDEX idx_products_name ON products(name);
CREATE INDEX idx_products_category_price ON products(category, price);