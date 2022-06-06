package db;
import (
	"database/sql"
	"testing"
	"log"
	"os"
	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:baoduong1011@localhost:5432/simple_bank?sslmode=disable"
)

var testQueries *Queries
// because NewStore need Sql.DB 
var testDB *sql.DB

func TestMain(m *testing.M) {
	var err error
	testDB,err = sql.Open(dbDriver,dbSource)
	if(err!=nil) {
		log.Fatal("Can not connect to database!",err);
	}

	testQueries = New(testDB)
	os.Exit(m.Run())
	
}