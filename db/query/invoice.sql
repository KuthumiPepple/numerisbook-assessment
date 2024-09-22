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

-- name: GetInvoice :one
SELECT * FROM invoices WHERE invoice_number = $1;