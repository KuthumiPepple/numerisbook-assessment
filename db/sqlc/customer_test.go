package db

import (
	"context"
	"testing"

	"github.com/kuthumipepple/numerisbook-assessment/util"
	"github.com/stretchr/testify/require"
)

func TestAddCustomer(t *testing.T) {
	arg := AddCustomerParams{
		Name:    util.RandomName(),
		Phone:   util.RandomPhone(),
		Address: util.RandomAddress(),
		Email:   util.RandomEmail(),
	}

	customer, err := testDb.AddCustomer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, customer)
	require.Equal(t, arg.Name, customer.Name)
	require.Equal(t, arg.Phone, customer.Phone)
	require.Equal(t, arg.Address, customer.Address)
	require.Equal(t, arg.Email, customer.Email)
	require.NotZero(t, customer.ID)
}
