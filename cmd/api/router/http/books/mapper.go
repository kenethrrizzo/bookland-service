package books

import (
	"net/http"
	"strconv"

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
	// TODO: Validar request para que lleguen parámetros requeridos (Name, Author, Price)

	return &books.Book{
		Name:      bookRequest.Name,
		Author:    bookRequest.Author,
		CoverPage: bookRequest.Coverpage,
		Synopsis:  bookRequest.Synopsis,
		Price:     bookRequest.Price,
	}
}

// TODO: Mejorar mapper
func bookFormToBookDomain(w http.ResponseWriter, r *http.Request) (*books.Book, error) {
	author, err := strconv.Atoi(r.FormValue("author"))
	if err != nil {
		return nil, err
	}

	price, err := strconv.ParseFloat(r.FormValue("price"), 64)
	if err != nil {
		return nil, err
	}

	return &books.Book{
		Name:      r.FormValue("name"),
		Author:    author,
		CoverPage: r.FormValue("coverpage"),
		Synopsis:  r.FormValue("synopsis"),
		Price:     price,
	}, nil
}
