-- name: CreateProduct :one
INSERT INTO products (supplier_id, name, description, category, price, stock, photos)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING id, created_at;

-- name: GetProducts :many
SELECT id, supplier_id, name, description, category, price, stock, photos, created_at
FROM products;

-- name: SearchProductsByName :many
SELECT id, supplier_id, name, description, category, stock, photos, price, created_at
FROM products
WHERE name ILIKE $1;

-- name: GetProductsWithFilters :many
-- name: SearchProducts :many
SELECT *
FROM products
WHERE 
    (name ILIKE sqlc.arg('search_query') OR description ILIKE sqlc.arg('search_query') OR sqlc.arg('search_query') IS NULL) AND
    (price >= sqlc.arg('price_from') OR sqlc.arg('price_from') IS NULL) AND
    (price <= sqlc.arg('price_to') OR sqlc.arg('price_to') IS NULL) AND
    (category = sqlc.arg('category') OR sqlc.arg('category') IS NULL) AND
    (stock > 0 OR sqlc.arg('in_stock') IS NULL OR sqlc.arg('in_stock') = FALSE);

-- name: GetProductByID :one
SELECT id, supplier_id, name, description, category, price, stock, photos, created_at
FROM products
WHERE id = $1;

-- name: AddToArchive :exec
INSERT INTO archives (user_id, product_id)
VALUES ($1, $2);

-- name: GetArchivesByUserID :many
SELECT id, user_id, product_id, created_at
FROM archives
WHERE user_id = $1;

-- name: CreateOrder :exec
INSERT INTO orders (buyer_id, product_id, quantity) VALUES ($1, $2, $3);

-- name: GetOrders :many
SELECT id, buyer_id, product_id, quantity, status, created_at FROM orders;

-- name: CreateUser :one
INSERT INTO users (email, password, company_name, inn, role) VALUES ($1, $2, $3, $4, $5) RETURNING id, created_at;

-- name: GetUserByEmail :one
SELECT id, email, password, company_name, inn, role, created_at FROM users WHERE email = $1;