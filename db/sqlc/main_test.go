package db

import (
	"database/sql"
	"log"
	"os"
	"testing"
	_ "github.com/lib/pq"
)


const (
dbSource = "postgresql://root:mysecretpassword@localhost:5433/simple_bank?sslmode=disable"
dbDriver = "postgres"
)
var testQueries *Queries

func  TestMain(m *testing.M) {
	conn, err := sql.Open(dbDriver, dbSource)

	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}
	testQueries = New(conn)
	os.Exit(m.Run())
}



