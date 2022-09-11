package main

import (
	"fmt"
	"net/http"

	"github.com/kenethrrizzo/bookland-service/cmd/api/config"
	"github.com/kenethrrizzo/bookland-service/cmd/api/data/database"
	router "github.com/kenethrrizzo/bookland-service/cmd/api/router/http"
)

func main() {
	config := config.LoadConfig()

	db, err := database.Connect(&config.Datasource)
	if err != nil {
		panic(err)
	}

	httpRouter := router.NewHTTPHandler(db)
	err = http.ListenAndServe(fmt.Sprintf(":%s", config.Server.Port), httpRouter)
	if err != nil {
		panic(err)
	}
}
