package main

import (
	"database/sql"
	"log"

	"github.com/baoduong1011/Project_Golang/api"
	db "github.com/baoduong1011/Project_Golang/db/sqlc"
	util "github.com/baoduong1011/Project_Golang/utils"
	_ "github.com/lib/pq"
	// "golang.org/x/tools/0.20220407163324-91bcfb1bdf9c/go/analysis/passes/nilness"
	// "golang.org/x/tools/go/analysis/passes/nilfunc"
)

// const (
// 	dbDriver = "postgres"
// 	dbSource = "postgresql://root:baoduong1011@localhost:5432/simple_bank?sslmode=disable"
// 	serverAddress = "0.0.0.0:8080"
// )

func main() {

	config , err := util.LoadiConfig(".")

	if(err != nil) {
		log.Fatal("Can not load config file to database!",err);
	}

	conn , err := sql.Open(config.DBDriver,config.DBSource)
	if(err!=nil) {
		log.Fatal("Can not connect to database!",err);
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)

	if err != nil {
		log.Fatal("Can not access to Server: ",err)
	}
}