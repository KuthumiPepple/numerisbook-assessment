-- name: AddCustomer :one
INSERT INTO customers (name, phone, address, email)
VALUES ($1, $2, $3, $4)
RETURNING *;