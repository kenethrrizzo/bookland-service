package http

import (
	"net/http"

	"github.com/gin-gonic/gin"

	bookHandler "github.com/kenethrrizzo/bookland-service/cmd/api/router/http/books"
)

const (
	BOOK_BASE_URL = "/books"
)

func NewHTTPHandler(bookHandler *bookHandler.BookHandler) http.Handler {
	router := gin.Default()

	booksGroup := router.Group("/books")
	{
		booksGroup.GET("/get", bookHandler.GetAllBooks)
		booksGroup.GET("/get/:bookID", bookHandler.GetBookByID)
		booksGroup.POST("/register", bookHandler.RegisterNewBook)
		booksGroup.PUT("/update/:bookID", bookHandler.UpdateBook)
		booksGroup.DELETE("/delete/:bookID", bookHandler.DeleteBook)
	}

	return router
}
