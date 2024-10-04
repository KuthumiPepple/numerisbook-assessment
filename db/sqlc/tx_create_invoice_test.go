package db

import (
	"context"
	"testing"
	"time"

	"github.com/kuthumipepple/numerisbook-assessment/util"
	"github.com/stretchr/testify/require"
)
func createRandomInvoice(t *testing.T) InvoiceResult{
	customer := addRandomCustomer(t)
	vendor := addRandomVendor(t)

	n := 10

	randomLineItems := make([]LineItemDetail, 0, n)
	var expSubtotal int64

	for i := 0; i < n; i++ {
		lineItem := LineItemDetail{
			Description: util.RandomString(10),
			Quantity:    util.RandomInt(1, 20),
			UnitPrice:   util.RandomInt(100, 1000),
		}
		expSubtotal += lineItem.Quantity * lineItem.UnitPrice
		randomLineItems = append(randomLineItems, lineItem)
	}

	arg := CreateInvoiceTxParams{
		CustomerID:       customer.ID,
		VendorID:         vendor.ID,
		IssueDate:        time.Now(),
		DueDate:          time.Now().AddDate(0, 0, 30),
		DiscountRate:     util.RandomInt(1, 10000),
		LineItemsDetails: randomLineItems,
	}
	expDiscount := (expSubtotal * arg.DiscountRate) / 10000
	expTotal := expSubtotal - expDiscount

	res, err := testStore.CreateInvoiceTx(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, res)

	// check invoice
	invoice := res.Invoice
	require.NotEmpty(t, invoice)
	require.NotZero(t, invoice.InvoiceNumber)
	require.Equal(t, arg.CustomerID, invoice.CustomerID)
	require.Equal(t, arg.VendorID, invoice.VendorID)
	require.WithinDuration(t, arg.IssueDate, invoice.IssueDate, time.Second)
	require.WithinDuration(t, arg.DueDate, invoice.DueDate, time.Second)
	require.Equal(t, arg.DiscountRate, invoice.DiscountRate)
	require.Equal(t, "draft", invoice.Status)
	require.Equal(t, expSubtotal, invoice.Subtotal)
	require.Equal(t, expDiscount, invoice.Discount)
	require.Equal(t, expTotal, invoice.TotalAmount)

	// check line items
	lineItems := res.LineItems
	require.Len(t, lineItems, n)
	for _, lineItem := range lineItems {
		require.NotZero(t, lineItem.ID)
		require.Equal(t, invoice.InvoiceNumber, lineItem.InvoiceNumber)
	}
	return res
}
func TestCreateInvoiceTx(t *testing.T) {
	createRandomInvoice(t)
}
