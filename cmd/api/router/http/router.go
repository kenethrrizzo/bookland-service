package http

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"

	bookHandler "github.com/kenethrrizzo/bookland-service/cmd/api/router/http/books"
)

const (
	BOOK_BASE_URL = "/books"
)

func NewHTTPHandler(bookHandler *bookHandler.BookHandler) http.Handler {
	router := chi.NewRouter()

	/* Book routes */
	router.Get(fmt.Sprintf("%s/get-all", BOOK_BASE_URL), bookHandler.GetAllBooks)
	router.Get(fmt.Sprintf("%s/get-by-id/{bookID}", BOOK_BASE_URL), bookHandler.GetBookByID)
	router.Post(fmt.Sprintf("%s/register", BOOK_BASE_URL), bookHandler.RegisterNewBook)
	router.Put(fmt.Sprintf("%s/update-cover-page", BOOK_BASE_URL), bookHandler.UpdateBookCoverImage)
	router.Delete(fmt.Sprintf("%s/delete/{bookID}", BOOK_BASE_URL), bookHandler.DeleteBook)

	return router
}
