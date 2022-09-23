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
	userRepository "github.com/kenethrrizzo/bookland-service/cmd/api/data/users"
	bookDomain "github.com/kenethrrizzo/bookland-service/cmd/api/domain/books"
	userDomain "github.com/kenethrrizzo/bookland-service/cmd/api/domain/users"
	router "github.com/kenethrrizzo/bookland-service/cmd/api/router/http"
	bookRouter "github.com/kenethrrizzo/bookland-service/cmd/api/router/http/books"
	userRouter "github.com/kenethrrizzo/bookland-service/cmd/api/router/http/users"
)

const (
	TEMP_DIRECTORY = "./tmp"
)

func main() {
	// * Crea un directorio temporal para guardar archivos
	err := os.MkdirAll(TEMP_DIRECTORY, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer os.RemoveAll(TEMP_DIRECTORY)

	config := config.LoadConfig()

	// Coneccion a la base de datos
	db, err := database.Connect(&config.Datasource)
	if err != nil {
		panic(err)
	}

	// Coneccion a la nube
	s3client, err := storage.Connect()
	if err != nil {
		panic(err)
	}

	/* files management */
	fileRepo := filebookRepository.NewStore(s3client)

	/* books */
	bookRepo := bookRepository.NewStore(db)
	bookService := bookDomain.NewService(bookRepo, fileRepo)
	bookHandler := bookRouter.NewHandler(bookService)

	/* users */
	userRepo := userRepository.NewStore(db)
	userService := userDomain.NewService(userRepo)
	userHandler := userRouter.NewHandler(userService)

	httpRouter := router.NewHTTPHandler(bookHandler, userHandler)

	err = http.ListenAndServe(fmt.Sprintf(":%s", config.Server.Port), httpRouter)
	if err != nil {
		panic(err)
	}
}
