package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
)

var testDb *Queries

func TestMain(m *testing.M) {
	testDbSource := "postgresql://root:secret@localhost:5432/numerisbookdb?sslmode=disable"

	connPool, err := pgxpool.New(context.Background(), testDbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testDb = New(connPool)

	os.Exit(m.Run())
}
