package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Store defines all functions to execute database queries and transactions
type Store interface {
	Querier
	CreateInvoiceTx(ctx context.Context, arg CreateInvoiceTxParams) (CreateInvoiceTxResult, error)
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
