package main

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/ubo-dev/turkey-address-api/internal/repository"

	"github.com/ubo-dev/turkey-address-api/internal/api"
)

func main() {

	store, err := repository.NewMysqlStorage()
	if err != nil {
		log.Fatal(err)
	}

	if err := store.Init(); err != nil {
		log.Fatal(err)
	}

	server := api.NewAPIServer(":8080", store)
	server.Run()
}
