package books

import (
	"github.com/kenethrrizzo/bookland-service/cmd/api/domain/books"
)

func bookDomaintoBookResponse(book *books.Book) *BookResponse {
	return &BookResponse{
		Id:        book.Id,
		Name:      book.Name,
		Author:    book.Author,
		CoverPage: book.CoverPage,
		Synopsis:  book.Synopsis,
		Price:     book.Price,
	}
}

func bookRequestToBookDomain(bookRequest *BookRequest) *books.Book {
	return &books.Book{
		Name:     bookRequest.Name,
		Author:   bookRequest.Author,
		Synopsis: bookRequest.Synopsis,
		Price:    bookRequest.Price,
	}
}
