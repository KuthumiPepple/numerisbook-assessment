-- name: AddVendor :one
INSERT INTO vendors (name, phone, address, email, bank_account_name, bank_account_no, bank_name)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING *;