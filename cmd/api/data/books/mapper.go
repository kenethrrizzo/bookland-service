package books

import (
	"database/sql"
	"strings"

	"github.com/kenethrrizzo/bookland-service/cmd/api/domain/books"
)

func bookDomainToBookSchema(book *books.Book) *Book {
	return &Book{
		ID:        book.ID,
		Name:      book.Name,
		Author:    book.Author,
		CoverPage: sql.NullString{String: book.CoverPage, Valid: true},
		Synopsis:  sql.NullString{String: book.Synopsis, Valid: true},
		Price:     book.Price,
		CreatedAt: book.CreatedAt,
		UpdatedAt: sql.NullTime{Time: book.UpdatedAt, Valid: true},
	}
}

func bookSchemaToBookDomain(bookSchema *Book) *books.Book {
	return &books.Book{
		ID:        bookSchema.ID,
		Name:      bookSchema.Name,
		Author:    bookSchema.Author,
		Genres:    strings.Split(bookSchema.Genres, ","),
		CoverPage: bookSchema.CoverPage.String,
		Synopsis:  bookSchema.Synopsis.String,
		Price:     bookSchema.Price,
		CreatedAt: bookSchema.CreatedAt,
		UpdatedAt: bookSchema.UpdatedAt.Time,
	}
}
