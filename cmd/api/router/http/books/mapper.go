package books

import (
	"strings"

	"github.com/kenethrrizzo/bookland-service/cmd/api/domain/books"
)

func bookDomaintoBookResponse(book *books.Book) *BookResponse {
	return &BookResponse{
		ID:        book.ID,
		Name:      book.Name,
		Author:    book.Author,
		Genres:    book.Genres,
		CoverPage: book.CoverPage,
		Synopsis:  book.Synopsis,
		Price:     book.Price,
	}
}

func bookRequestToBookDomain(bookRequest *BookRequest) *books.Book {
	return &books.Book{
		Name:     bookRequest.Name,
		Author:   bookRequest.Author,
		Genres:   strings.Split(bookRequest.Genres, ","),
		Synopsis: bookRequest.Synopsis,
		Price:    bookRequest.Price,
	}
}
