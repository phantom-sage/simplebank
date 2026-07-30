package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/phantom-sage/simplebank/util"
)

var testQueries *Queries
var testConn *pgxpool.Pool

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../../")
	if err != nil {
		log.Fatalln("can not load configuration file", err)
	}
	testConn, err = pgxpool.New(context.Background(), config.DBConnUrl)
	if err != nil {
		log.Fatalln("can not connect to db", err)
	}

	testQueries = New(testConn)
	os.Exit(m.Run())
}
