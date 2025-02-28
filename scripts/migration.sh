#!/bin/bash

DB_HOST="localhost"
DB_PORT="5432"
DB_USER="user"
DB_PASSWORD="password"
DB_NAME="dbname"

psql "host=$DB_HOST port=$DB_PORT user=$DB_USER password=$DB_PASSWORD dbname=$DB_NAME sslmode=disable" -f ./internal/db/migrations/000001_create_users_table.up.sql
psql "host=$DB_HOST port=$DB_PORT user=$DB_USER password=$DB_PASSWORD dbname=$DB_NAME sslmode=disable" -f ./internal/db/migrations/000002_create_products_table.up.sql
psql "host=$DB_HOST port=$DB_PORT user=$DB_USER password=$DB_PASSWORD dbname=$DB_NAME sslmode=disable" -f ./internal/db/migrations/000003_create_orders_table.up.sql