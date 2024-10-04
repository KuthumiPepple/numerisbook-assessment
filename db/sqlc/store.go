package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Store defines all functions to execute database queries and transactions
type Store interface {
	Querier
	CreateInvoiceTx(ctx context.Context, arg CreateInvoiceTxParams) (InvoiceResult, error)
	GetInvoice(ctx context.Context, id int64) (InvoiceResult, error)
}

// SQLStore provides all functions to execute SQL queries and transactions.
type SQLStore struct {
	connPool *pgxpool.Pool
	*Queries
}

func NewStore(connPool *pgxpool.Pool) Store {
	return &SQLStore{
		connPool: connPool,
		Queries:  New(connPool),
	}
}

const getInvoice = `
SELECT 
    i.invoice_number,
    i.customer_id,
    i.vendor_id,
    i.issue_date,
    i.due_date,
    i.status,
    i.subtotal,
    i.discount_rate,
    i.discount,
    i.total_amount,
    i.billing_currency,
    i.note,
    li.id,
	li.invoice_number,
    li.description,
    li.quantity,
    li.unit_price,
	li.total_price
FROM 
    invoices i
JOIN 
    line_items li ON i.invoice_number = li.invoice_number
WHERE 
    i.invoice_number = $1;
`

func (store *SQLStore) GetInvoice(ctx context.Context, id int64) (InvoiceResult, error) {
	var result InvoiceResult
	rows, err := store.connPool.Query(ctx, getInvoice, id)
	if err != nil {
		return result, err
	}
	defer rows.Close()

	var invoice Invoice
	var invoiceSet bool
	lineItems := []LineItem{}

	for rows.Next() {
		var li LineItem
		var tempInvoice Invoice
		err := rows.Scan(
			&tempInvoice.InvoiceNumber,
			&tempInvoice.CustomerID,
			&tempInvoice.VendorID,
			&tempInvoice.IssueDate,
			&tempInvoice.DueDate,
			&tempInvoice.Status,
			&tempInvoice.Subtotal,
			&tempInvoice.DiscountRate,
			&tempInvoice.Discount,
			&tempInvoice.TotalAmount,
			&tempInvoice.BillingCurrency,
			&tempInvoice.Note,
			&li.ID,
			&li.InvoiceNumber,
			&li.Description,
			&li.Quantity,
			&li.UnitPrice,
			&li.TotalPrice,
		)
		if err != nil {
			return result, err
		}

		if !invoiceSet {
			invoice = tempInvoice
			invoiceSet = true
		}
		lineItems = append(lineItems, li)
	}
	if err = rows.Err(); err != nil {
		return result, err
	}

	result.Invoice = invoice
	result.LineItems = lineItems

	return result, nil
}