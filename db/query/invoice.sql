-- name: AddNoItemsInvoice :one
INSERT INTO invoices (
    customer_id, 
    vendor_id, 
    issue_date, 
    due_date, 
    discount_rate
) VALUES (
    $1, $2, $3, $4, $5
) RETURNING *;

-- name: AddLineItem :one
INSERT INTO line_items (
    invoice_number,
    description,
    quantity,
    unit_price,
    total_price
) VALUES ($1, $2, $3, $4, $5) RETURNING *;

-- name: GetInvoiceLineItems :many
SELECT * FROM line_items WHERE invoice_number = $1;

-- name: UpdateInvoice :one
UPDATE invoices
SET
    status = COALESCE(sqlc.narg(status), status),
    subtotal = COALESCE(sqlc.narg(subtotal), subtotal),
    discount_rate = COALESCE(sqlc.narg(discount_rate), discount_rate),
    discount = COALESCE(sqlc.narg(discount), discount),
    total_amount = COALESCE(sqlc.narg(total_amount), total_amount)
WHERE
    invoice_number = sqlc.arg(invoice_number)
RETURNING *;