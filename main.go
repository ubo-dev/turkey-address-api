package main

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	store, err := NewMysqlStorage()
	if err != nil {
		log.Fatal(err)
	}

	if err := store.Init(); err != nil {
		log.Fatal(err)
	}

	server := NewAPIServer(":8080", store)
	server.Run()
}
