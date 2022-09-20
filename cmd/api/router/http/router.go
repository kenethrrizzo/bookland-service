package http

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	bookRouter "github.com/kenethrrizzo/bookland-service/cmd/api/router/http/books"
	userRouter "github.com/kenethrrizzo/bookland-service/cmd/api/router/http/users"

	"github.com/kenethrrizzo/bookland-service/cmd/api/router/http/middlewares/auth"
)

const (
	BOOK_BASE_URL = "/books"
)

func NewHTTPHandler(bookHandler *bookRouter.BookHandler, userHandler *userRouter.UserHandler) http.Handler {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
		ExposeHeaders:    []string{"Content-Length"},
	}))

	usersGroup := router.Group("/users")
	{
		usersGroup.POST("/register", userHandler.Register)
		usersGroup.POST("/login", userHandler.Login)
	}

	booksGroup := router.Group("/books")
	{
		booksGroup.Use(auth.ValidateJWT)

		booksGroup.GET("/get", bookHandler.GetAllBooks)
		booksGroup.GET("/get/genre/:genre", bookHandler.GetBooksByGenre)
		booksGroup.GET("/get/:bookID", bookHandler.GetBookByID)
		booksGroup.POST("/register", bookHandler.RegisterNewBook)
		booksGroup.PUT("/update/:bookID", bookHandler.UpdateBook)
		booksGroup.DELETE("/delete/:bookID", bookHandler.DeleteBook)
	}

	return router
}
