package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
)

var testStore Store

func TestMain(m *testing.M) {
	testStoreSource := "postgresql://root:secret@localhost:5432/numerisbookdb?sslmode=disable"

	connPool, err := pgxpool.New(context.Background(), testStoreSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testStore = NewStore(connPool)

	os.Exit(m.Run())
}
