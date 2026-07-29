package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
)

const connString = "postgres://root:secret@localhost:5432/simple_bank?sslmode=disable"

var testQueries *Queries
var testConn *pgxpool.Pool

func TestMain(m *testing.M) {
	var err error

	testConn, err = pgxpool.New(context.Background(), connString)
	if err != nil {
		log.Fatalln("can not connect to db", err)
	}

	testQueries = New(testConn)
	os.Exit(m.Run())
}
