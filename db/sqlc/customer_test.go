package db

import (
	"context"
	"testing"

	"github.com/kuthumipepple/numerisbook-assessment/util"
	"github.com/stretchr/testify/require"
)

func addRandomCustomer(t *testing.T) Customer {
	arg := AddCustomerParams{
		Name:    util.RandomName(),
		Phone:   util.RandomPhone(),
		Address: util.RandomAddress(),
		Email:   util.RandomEmail(),
	}

	customer, err := testStore.AddCustomer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, customer)
	require.Equal(t, arg.Name, customer.Name)
	require.Equal(t, arg.Phone, customer.Phone)
	require.Equal(t, arg.Address, customer.Address)
	require.Equal(t, arg.Email, customer.Email)
	require.NotZero(t, customer.ID)

	return customer
}

func TestAddCustomer(t *testing.T) {
	addRandomCustomer(t)
}
