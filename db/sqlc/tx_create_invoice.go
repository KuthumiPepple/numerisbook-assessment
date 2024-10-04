package db

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

type LineItemDetail struct {
	Description string `json:"description"`
	Quantity    int64  `json:"quantity"`
	UnitPrice   int64  `json:"unit_price"`
}

type CreateInvoiceTxParams struct {
	CustomerID       int64            `json:"customer_id"`
	VendorID         int64            `json:"vendor_id"`
	IssueDate        time.Time        `json:"issue_date"`
	DueDate          time.Time        `json:"due_date"`
	DiscountRate     int64            `json:"discount_rate"` // in basis points
	LineItemsDetails []LineItemDetail `json:"line_items"`
}

type InvoiceResult struct {
	Invoice   Invoice    `json:"invoice"`
	LineItems []LineItem `json:"line_items"`
}

func (store *SQLStore) CreateInvoiceTx(ctx context.Context, arg CreateInvoiceTxParams) (InvoiceResult, error) {
	var result InvoiceResult

	err := store.execTx(
		ctx,
		func(q *Queries) error {

			var err error
			invoice, err := q.AddNoItemsInvoice(
				ctx,
				AddNoItemsInvoiceParams{
					CustomerID:   arg.CustomerID,
					VendorID:     arg.VendorID,
					IssueDate:    arg.IssueDate,
					DueDate:      arg.DueDate,
					DiscountRate: arg.DiscountRate,
				},
			)
			if err != nil {
				return err
			}

			var sum int64
			for _, LineItemDetail := range arg.LineItemsDetails {
				itemTotalPrice := LineItemDetail.Quantity * LineItemDetail.UnitPrice
				lineItem, err := q.AddLineItem(
					ctx,
					AddLineItemParams{
						InvoiceNumber: invoice.InvoiceNumber,
						Description:   LineItemDetail.Description,
						Quantity:      LineItemDetail.Quantity,
						UnitPrice:     LineItemDetail.UnitPrice,
						TotalPrice:    itemTotalPrice,
					},
				)
				if err != nil {
					return err
				}
				sum += itemTotalPrice
				result.LineItems = append(result.LineItems, lineItem)
			}

			discount := (sum * arg.DiscountRate) / 10000

			updatedInvoice, err := q.UpdateInvoice(
				ctx,
				UpdateInvoiceParams{
					InvoiceNumber: invoice.InvoiceNumber,
					Subtotal: pgtype.Int8{
						Int64: sum,
						Valid: true,
					},
					Discount: pgtype.Int8{
						Int64: discount,
						Valid: true,
					},
					TotalAmount: pgtype.Int8{
						Int64: sum - discount,
						Valid: true,
					},
				},
			)
			if err != nil {
				return err
			}

			result.Invoice = updatedInvoice

			return err
		},
	)
	return result, err
}
