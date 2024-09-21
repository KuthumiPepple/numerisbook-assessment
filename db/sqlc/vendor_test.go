package db

import (
	"context"
	"testing"

	"github.com/kuthumipepple/numerisbook-assessment/util"
	"github.com/stretchr/testify/require"
)

func addRandomVendor(t *testing.T) Vendor {
	arg := AddVendorParams{
		Name:            util.RandomName(),
		Phone:           util.RandomPhone(),
		Address:         util.RandomAddress(),
		Email:           util.RandomEmail(),
		BankAccountName: util.RandomName(),
		BankAccountNo:   util.RandomInt(1000, 9999),
		BankName:        util.RandomName(),
	}

	vendor, err := testDb.AddVendor(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, vendor)
	require.Equal(t, arg.Name, vendor.Name)
	require.Equal(t, arg.Phone, vendor.Phone)
	require.Equal(t, arg.Address, vendor.Address)
	require.Equal(t, arg.Email, vendor.Email)
	require.Equal(t, arg.BankAccountName, vendor.BankAccountName)
	require.Equal(t, arg.BankAccountNo, vendor.BankAccountNo)
	require.Equal(t, arg.BankName, vendor.BankName)
	require.NotZero(t, vendor.ID)

	return vendor
}
func TestAddVendor(t *testing.T) {
	addRandomVendor(t)
}
