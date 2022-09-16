package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/kenethrrizzo/bookland-service/cmd/api/config"
	bookRepository "github.com/kenethrrizzo/bookland-service/cmd/api/data/books"
	"github.com/kenethrrizzo/bookland-service/cmd/api/data/connections/database"
	"github.com/kenethrrizzo/bookland-service/cmd/api/data/connections/storage"
	filebookRepository "github.com/kenethrrizzo/bookland-service/cmd/api/data/files"
	bookDomain "github.com/kenethrrizzo/bookland-service/cmd/api/domain/books"
	router "github.com/kenethrrizzo/bookland-service/cmd/api/router/http"
	bookHandler "github.com/kenethrrizzo/bookland-service/cmd/api/router/http/books"
)

const (
	TEMP_DIRECTORY = "./tmp"
)

func main() {
	// * Crea un directorio temporal para guardar archivos
	err := os.Mkdir(TEMP_DIRECTORY, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer os.RemoveAll(TEMP_DIRECTORY)

	config := config.LoadConfig()

	db, err := database.Connect(&config.Datasource)
	if err != nil {
		panic(err)
	}

	s3client, err := storage.Connect()
	if err != nil {
		panic(err)
	}

	/* files management */
	fileRepo := filebookRepository.NewStore(s3client)

	/* books */
	bookRepo := bookRepository.NewStore(db)
	bookService := bookDomain.NewService(bookRepo, fileRepo)
	bookHandler := bookHandler.NewHandler(bookService)

	httpRouter := router.NewHTTPHandler(bookHandler)

	err = http.ListenAndServe(fmt.Sprintf(":%s", config.Server.Port), httpRouter)
	if err != nil {
		panic(err)
	}
}
