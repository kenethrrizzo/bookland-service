package http

import (
	"net/http"

	"github.com/gin-gonic/gin"

	bookRouter "github.com/kenethrrizzo/bookland-service/cmd/api/router/http/books"
	userRouter "github.com/kenethrrizzo/bookland-service/cmd/api/router/http/users"
)

const (
	BOOK_BASE_URL = "/books"
)

func NewHTTPHandler(bookHandler *bookRouter.BookHandler, userHandler *userRouter.UserHandler) http.Handler {
	router := gin.Default()

	usersGroup := router.Group("/users")
	{
		usersGroup.POST("/register", userHandler.Register)
		usersGroup.POST("/login", userHandler.Login)
	}

	booksGroup := router.Group("/books")
	{
		booksGroup.Use(userRouter.ValidateJWT)

		booksGroup.GET("/get", bookHandler.GetAllBooks)
		booksGroup.GET("/get/genre/:genre", bookHandler.GetBooksByGenre)
		booksGroup.GET("/get/:bookID", bookHandler.GetBookByID)
		booksGroup.POST("/register", bookHandler.RegisterNewBook)
		booksGroup.PUT("/update/:bookID", bookHandler.UpdateBook)
		booksGroup.DELETE("/delete/:bookID", bookHandler.DeleteBook)
	}

	return router
}
