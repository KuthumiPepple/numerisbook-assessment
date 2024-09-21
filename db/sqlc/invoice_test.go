package db

import (
	"context"
	"testing"
	"time"

	"github.com/kuthumipepple/numerisbook-assessment/util"
	"github.com/stretchr/testify/require"
)

func addRandomNoItemsInvoice(t *testing.T) Invoice {
	customer := addRandomCustomer(t)
	vendor := addRandomVendor(t)

	arg := AddNoItemsInvoiceParams{
		CustomerID:   customer.ID,
		VendorID:     vendor.ID,
		IssueDate:    util.RandomPastDate(),
		DueDate:      util.RandomFutureDate(),
		DiscountRate: util.RandomInt(1, 10),
	}

	invoice, err := testDb.AddNoItemsInvoice(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, invoice)
	require.NotZero(t, invoice.InvoiceNumber)
	require.Equal(t, arg.CustomerID, invoice.CustomerID)
	require.Equal(t, arg.VendorID, invoice.VendorID)
	require.WithinDuration(t, arg.IssueDate, invoice.IssueDate, time.Second)
	require.WithinDuration(t, arg.DueDate, invoice.DueDate, time.Second)
	require.Equal(t, arg.DiscountRate, invoice.DiscountRate)
	require.Equal(t, "draft", invoice.Status)
	require.Zero(t, invoice.Subtotal)
	require.Zero(t, invoice.Discount)
	require.Zero(t, invoice.TotalAmount)
	require.NotEmpty(t, invoice.BillingCurrency)
	require.NotEmpty(t, invoice.Note)

	return invoice
}

func TestAddNoItemsInvoice(t *testing.T) {
	addRandomNoItemsInvoice(t)
}

func TestAddLineItem(t *testing.T) {
	invoice := addRandomNoItemsInvoice(t)

	arg := AddLineItemParams{
		InvoiceNumber: invoice.InvoiceNumber,
		Description:   util.RandomString(10),
		Quantity:      util.RandomInt(1, 10),
		UnitPrice:     util.RandomInt(100, 1000),
	}
	arg.TotalPrice = arg.Quantity * arg.UnitPrice

	lineItem, err := testDb.AddLineItem(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, lineItem)
	require.NotZero(t, lineItem.ID)
	require.Equal(t, arg.InvoiceNumber, lineItem.InvoiceNumber)
	require.Equal(t, arg.Description, lineItem.Description)
	require.Equal(t, arg.Quantity, lineItem.Quantity)
	require.Equal(t, arg.UnitPrice, lineItem.UnitPrice)
	require.Equal(t, arg.Quantity*arg.UnitPrice, lineItem.TotalPrice)
}
