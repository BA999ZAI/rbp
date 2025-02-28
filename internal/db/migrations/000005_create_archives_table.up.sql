CREATE TABLE archives (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id) NOT NULL,
    product_id INT REFERENCES products(id) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);