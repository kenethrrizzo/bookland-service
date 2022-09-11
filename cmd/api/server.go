package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/kenethrrizzo/bookland-service/cmd/api/config"
	"github.com/kenethrrizzo/bookland-service/cmd/api/data/database"
	router "github.com/kenethrrizzo/bookland-service/cmd/api/router/http"
)

func main() {
	config := config.LoadConfig()

	// Connecting to database
	db, err := database.Connect(&config.Datasource)
	if err != nil {
		log.Fatalln("ha ocurrido un error al conectarse a la base de datos: ", err)
	}

	httpRouter := router.NewHTTPHandler(db)
	log.Println("Listening server in port: " + config.Server.Port)
	err = http.ListenAndServe(fmt.Sprintf(":%s", config.Server.Port), httpRouter)
	if err != nil {
		panic(err)
	}
}
