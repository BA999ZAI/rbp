CREATE TYPE order_status AS ENUM ('new', 'processing', 'completed', 'canceled');

CREATE TABLE orders (
    id SERIAL PRIMARY KEY,
    buyer_id INT REFERENCES users(id) NOT NULL,
    product_id INT REFERENCES products(id) NOT NULL,
    quantity INT NOT NULL,
    status order_status DEFAULT 'new' NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);