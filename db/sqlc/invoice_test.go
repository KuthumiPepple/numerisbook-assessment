package db

import (
	"context"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/kuthumipepple/numerisbook-assessment/util"
	"github.com/stretchr/testify/require"
)

func addRandomNoItemsInvoice(t *testing.T) Invoice {
	customer := addRandomCustomer(t)
	vendor := addRandomVendor(t)

	arg := AddNoItemsInvoiceParams{
		CustomerID:   customer.ID,
		VendorID:     vendor.ID,
		IssueDate:    time.Now(),
		DueDate:      time.Now().AddDate(0, 0, 30),
		DiscountRate: util.RandomInt(1, 10),
	}

	invoice, err := testStore.AddNoItemsInvoice(context.Background(), arg)
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

	lineItem, err := testStore.AddLineItem(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, lineItem)
	require.NotZero(t, lineItem.ID)
	require.Equal(t, arg.InvoiceNumber, lineItem.InvoiceNumber)
	require.Equal(t, arg.Description, lineItem.Description)
	require.Equal(t, arg.Quantity, lineItem.Quantity)
	require.Equal(t, arg.UnitPrice, lineItem.UnitPrice)
	require.Equal(t, arg.Quantity*arg.UnitPrice, lineItem.TotalPrice)
}

func TestGetInvoiceLineItems(t *testing.T) {
	invoice := addRandomNoItemsInvoice(t)

	n := 5

	existedLineItems := make(map[int64]bool)

	for i := 0; i < n; i++ {
		arg := AddLineItemParams{
			InvoiceNumber: invoice.InvoiceNumber,
			Description:   util.RandomString(10),
			Quantity:      util.RandomInt(1, 10),
			UnitPrice:     util.RandomInt(100, 1000),
		}
		arg.TotalPrice = arg.Quantity * arg.UnitPrice

		lineItem, err := testStore.AddLineItem(context.Background(), arg)
		require.NoError(t, err)
		require.NotEmpty(t, lineItem)
		existedLineItems[lineItem.ID] = true
	}

	lineItems, err := testStore.GetInvoiceLineItems(context.Background(), invoice.InvoiceNumber)
	require.NoError(t, err)
	require.Len(t, lineItems, n)

	for _, lineItem := range lineItems {
		require.NotZero(t, lineItem)
		require.Equal(t, invoice.InvoiceNumber, lineItem.InvoiceNumber)

		_, ok := existedLineItems[lineItem.ID]
		require.True(t, ok)
	}
}

func TestUpdateInvoiceOnlyStatus(t *testing.T) {
	oldInvoice := addRandomNoItemsInvoice(t)

	newStatus := util.RandomStatusExcludingDraft()

	updatedInvoice, err := testStore.UpdateInvoice(
		context.Background(),
		UpdateInvoiceParams{
			InvoiceNumber: oldInvoice.InvoiceNumber,
			Status: pgtype.Text{
				String: newStatus,
				Valid:  true,
			},
		},
	)

	require.NoError(t, err)
	require.NotEqual(t, oldInvoice.Status, updatedInvoice.Status)
	require.Equal(t, newStatus, updatedInvoice.Status)
	require.Equal(t, oldInvoice.Subtotal, updatedInvoice.Subtotal)
	require.Equal(t, oldInvoice.DiscountRate, updatedInvoice.DiscountRate)
	require.Equal(t, oldInvoice.Discount, updatedInvoice.Discount)
	require.Equal(t, oldInvoice.TotalAmount, updatedInvoice.TotalAmount)
}

func TestUpdateInvoiceOnlyDiscountRate(t *testing.T) {
	oldInvoice := addRandomNoItemsInvoice(t)

	newDiscountRate := util.RandomInt(1, 10)

	updatedInvoice, err := testStore.UpdateInvoice(
		context.Background(),
		UpdateInvoiceParams{
			InvoiceNumber: oldInvoice.InvoiceNumber,
			DiscountRate: pgtype.Int8{
				Int64: newDiscountRate,
				Valid: true,
			},
		},
	)

	require.NoError(t, err)
	require.NotEqual(t, oldInvoice.DiscountRate, updatedInvoice.DiscountRate)
	require.Equal(t, newDiscountRate, updatedInvoice.DiscountRate)
	require.Equal(t, oldInvoice.Status, updatedInvoice.Status)
	require.Equal(t, oldInvoice.Subtotal, updatedInvoice.Subtotal)
	require.Equal(t, oldInvoice.Discount, updatedInvoice.Discount)
	require.Equal(t, oldInvoice.TotalAmount, updatedInvoice.TotalAmount)
}

func TestUpdateInvoiceOnlyDiscountrateDiscountSubtotalTotalamount(t *testing.T) {
	oldInvoice := addRandomNoItemsInvoice(t)

	newSubtotal := util.RandomInt(1000, 10000) * 100
	newDiscountRate := util.RandomInt(1, 10)
	newDiscount := newSubtotal * newDiscountRate / 100
	newTotalAmount := newSubtotal - newDiscount

	updatedInvoice, err := testStore.UpdateInvoice(
		context.Background(),
		UpdateInvoiceParams{
			InvoiceNumber: oldInvoice.InvoiceNumber,
			Subtotal: pgtype.Int8{
				Int64: newSubtotal,
				Valid: true,
			},
			DiscountRate: pgtype.Int8{
				Int64: newDiscountRate,
				Valid: true,
			},
			Discount: pgtype.Int8{
				Int64: newDiscount,
				Valid: true,
			},
			TotalAmount: pgtype.Int8{
				Int64: newTotalAmount,
				Valid: true,
			},
		},
	)

	require.NoError(t, err)
	require.NotEqual(t, oldInvoice.Subtotal, updatedInvoice.Subtotal)
	require.NotEqual(t, oldInvoice.DiscountRate, updatedInvoice.DiscountRate)
	require.NotEqual(t, oldInvoice.Discount, updatedInvoice.Discount)
	require.NotEqual(t, oldInvoice.TotalAmount, updatedInvoice.TotalAmount)
	require.Equal(t, newSubtotal, updatedInvoice.Subtotal)
	require.Equal(t, newDiscountRate, updatedInvoice.DiscountRate)
	require.Equal(t, newDiscount, updatedInvoice.Discount)
	require.Equal(t, newTotalAmount, updatedInvoice.TotalAmount)
	require.Equal(t, oldInvoice.Status, updatedInvoice.Status)
}

func TestUpdateInvoiceAllUpdatableFields(t *testing.T) {
	oldInvoice := addRandomNoItemsInvoice(t)

	newStatus := util.RandomStatusExcludingDraft()
	newSubtotal := util.RandomInt(1000, 10000) * 100
	newDiscountRate := util.RandomInt(1, 10)
	newDiscount := newSubtotal * newDiscountRate / 100
	newTotalAmount := newSubtotal - newDiscount

	updatedInvoice, err := testStore.UpdateInvoice(
		context.Background(),
		UpdateInvoiceParams{
			InvoiceNumber: oldInvoice.InvoiceNumber,
			Status: pgtype.Text{
				String: newStatus,
				Valid:  true,
			},
			Subtotal: pgtype.Int8{
				Int64: newSubtotal,
				Valid: true,
			},
			DiscountRate: pgtype.Int8{
				Int64: newDiscountRate,
				Valid: true,
			},
			Discount: pgtype.Int8{
				Int64: newDiscount,
				Valid: true,
			},
			TotalAmount: pgtype.Int8{
				Int64: newTotalAmount,
				Valid: true,
			},
		},
	)

	require.NoError(t, err)
	require.NotEqual(t, oldInvoice.Status, updatedInvoice.Status)
	require.NotEqual(t, oldInvoice.Subtotal, updatedInvoice.Subtotal)
	require.NotEqual(t, oldInvoice.DiscountRate, updatedInvoice.DiscountRate)
	require.NotEqual(t, oldInvoice.Discount, updatedInvoice.Discount)
	require.NotEqual(t, oldInvoice.TotalAmount, updatedInvoice.TotalAmount)
	require.Equal(t, newStatus, updatedInvoice.Status)
	require.Equal(t, newSubtotal, updatedInvoice.Subtotal)
	require.Equal(t, newDiscountRate, updatedInvoice.DiscountRate)
	require.Equal(t, newDiscount, updatedInvoice.Discount)
	require.Equal(t, newTotalAmount, updatedInvoice.TotalAmount)
}
