package db

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestGetInvoice(t *testing.T) {
	arg := createRandomInvoice(t)
	res, err := testStore.GetInvoice(context.Background(), arg.Invoice.InvoiceNumber)
	require.NoError(t, err)
	require.NotEmpty(t, res)

	// check invoice
	gotInvoice := res.Invoice
	require.NotEmpty(t, gotInvoice)
	require.Equal(t, arg.Invoice.InvoiceNumber, gotInvoice.InvoiceNumber)
	require.Equal(t, arg.Invoice.CustomerID, gotInvoice.CustomerID)
	require.Equal(t, arg.Invoice.VendorID, gotInvoice.VendorID)
	require.WithinDuration(t, arg.Invoice.IssueDate, gotInvoice.IssueDate, time.Second)
	require.WithinDuration(t, arg.Invoice.DueDate, gotInvoice.DueDate, time.Second)
	require.Equal(t, arg.Invoice.Status, gotInvoice.Status)
	require.Equal(t, arg.Invoice.Subtotal, gotInvoice.Subtotal)
	require.Equal(t, arg.Invoice.DiscountRate, gotInvoice.DiscountRate)
	require.Equal(t, arg.Invoice.Discount, gotInvoice.Discount)
	require.Equal(t, arg.Invoice.TotalAmount, gotInvoice.TotalAmount)

	// check line items
	gotLineItems := res.LineItems
	require.Len(t, gotLineItems, len(arg.LineItems))
	for _, lineItem := range gotLineItems {
		require.NotZero(t, lineItem.ID)
		require.Equal(t, arg.Invoice.InvoiceNumber, lineItem.InvoiceNumber)
	}
}
