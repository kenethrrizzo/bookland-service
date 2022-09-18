package books

import (
	"database/sql"
	"time"

	"github.com/pkg/errors"

	"github.com/kenethrrizzo/bookland-service/cmd/api/domain/books"
	domainErrors "github.com/kenethrrizzo/bookland-service/cmd/api/domain/errors"
)

const (
	createError       = "error creando un nuevo libro"
	updateError       = "error actualizando el libro"
	deleteError       = "error eliminando libro"
	readError         = "error buscando un libro en la base de datos"
	listError         = "error obteniendo libros de la base de datos"
	lastInsertIDError = "error obteniendo último ID insertado"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db}
}

func (s *Store) GetAllBooks() ([]books.Book, error) {
	var booksDomain []books.Book

	sqlSelect := `SELECT b.Id, b.Name, b.Author, b.CoverPage, b.Synopsis, b.Price, b.CreatedAt, b.UpdatedAt, group_concat(bg.GenreCode) 
		FROM Book b INNER JOIN BookGenre bg ON bg.BookId = b.Id GROUP BY b.Id`
	res, err := s.db.Query(sqlSelect)
	if err != nil {
		if err == sql.ErrNoRows {
			appErr := domainErrors.NewAppErrorWithType(domainErrors.NotFound)
			return nil, appErr
		}
		appErr := domainErrors.NewAppError(errors.Wrap(err, listError), domainErrors.RepositoryError)
		return nil, appErr
	}
	defer res.Close()

	for res.Next() {
		var book Book

		err := res.Scan(&book.ID, &book.Name, &book.Author, &book.CoverPage, &book.Synopsis, &book.Price, &book.CreatedAt, &book.UpdatedAt, &book.Genres)
		if err != nil {
			appErr := domainErrors.NewAppErrorWithType(domainErrors.MapError)
			return nil, appErr
		}

		booksDomain = append(booksDomain, *bookSchemaToBookDomain(&book))
	}

	return booksDomain, nil
}

// TODO: CORREGIR PARA QUE APAREZCAN TODOS LOS CODES
func (s *Store) GetBooksByGenre(genre string) ([]books.Book, error) {
	var booksDomain []books.Book

	sqlSelect := `SELECT b.Id, b.Name, b.Author, b.CoverPage, b.Synopsis, b.Price, b.CreatedAt, b.UpdatedAt, group_concat(bg.GenreCode) as 'Genres'
		FROM Book b INNER JOIN BookGenre bg ON bg.BookId = b.Id WHERE bg.GenreCode = ? GROUP BY b.Id`
	res, err := s.db.Query(sqlSelect, genre)
	if err != nil {
		if err == sql.ErrNoRows {
			appErr := domainErrors.NewAppErrorWithType(domainErrors.NotFound)
			return nil, appErr
		}
		appErr := domainErrors.NewAppError(errors.Wrap(err, listError), domainErrors.RepositoryError)
		return nil, appErr
	}
	defer res.Close()

	for res.Next() {
		var book Book
		err := res.Scan(&book.ID, &book.Name, &book.Author, &book.CoverPage, &book.Synopsis, &book.Price, &book.CreatedAt, &book.UpdatedAt, &book.Genres)
		if err != nil {
			appErr := domainErrors.NewAppErrorWithType(domainErrors.MapError)
			return nil, appErr
		}

		booksDomain = append(booksDomain, *bookSchemaToBookDomain(&book))
	}

	return booksDomain, nil
}

func (s *Store) GetBookByID(id int) (*books.Book, error) {
	var book Book

	sqlSelect := "SELECT Id, Name, Author, CoverPage, Synopsis, Price, CreatedAt, UpdatedAt FROM Book WHERE Id = ?"
	err := s.db.QueryRow(sqlSelect, id).Scan(&book.ID, &book.Name, &book.Author, &book.CoverPage, &book.Synopsis, &book.Price, &book.CreatedAt, &book.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			appErr := domainErrors.NewAppErrorWithType(domainErrors.NotFound)
			return nil, appErr
		}
		appErr := domainErrors.NewAppError(errors.Wrap(err, readError), domainErrors.RepositoryError)
		return nil, appErr
	}

	return bookSchemaToBookDomain(&book), nil
}

func (s *Store) CreateBook(book *books.Book) (*books.Book, error) {
	bookEntity := bookDomainToBookSchema(book)

	sqlInsert := "INSERT INTO Book (Name, Author, CoverPage, Synopsis, Price) values (?, ?, ?, ?, ?)"
	res, err := s.db.Exec(sqlInsert, bookEntity.Name, bookEntity.Author, bookEntity.CoverPage, bookEntity.Synopsis, bookEntity.Price)
	if err != nil {
		appErr := domainErrors.NewAppError(errors.Wrap(err, createError), domainErrors.RepositoryError)
		return nil, appErr
	}

	lastId, err := res.LastInsertId()
	if err != nil {
		appErr := domainErrors.NewAppError(errors.Wrap(err, lastInsertIDError), domainErrors.RepositoryError)
		return nil, appErr
	}

	book.ID = int(lastId)

	if len(book.Genres) == 1 && book.Genres[0] == "" {
		return book, nil
	}

	for _, genre := range book.Genres {
		bookGenreEntity := &BookGenre{
			BookID:    book.ID,
			GenreCode: genre,
		}

		sqlInsert = "INSERT INTO BookGenre (BookId, GenreCode) values (?, ?)"
		_, err = s.db.Exec(sqlInsert, bookGenreEntity.BookID, bookGenreEntity.GenreCode)
		if err != nil {
			appErr := domainErrors.NewAppError(errors.Wrap(err, createError), domainErrors.RepositoryError)
			return nil, appErr
		}
	}

	return book, nil
}

func (s *Store) UpdateBook(book *books.Book) (*books.Book, error) {
	bookEntity := bookDomainToBookSchema(book)

	sqlUpdate := "UPDATE Book SET Name = ?, Author = ?, CoverPage = ?, Synopsis = ?, Price = ?, UpdatedAt = ? WHERE Id = ?"
	_, err := s.db.Exec(sqlUpdate, bookEntity.Name, bookEntity.Author, bookEntity.CoverPage, bookEntity.Synopsis, bookEntity.Price, time.Now().UTC(), bookEntity.ID)
	if err != nil {
		appErr := domainErrors.NewAppError(errors.Wrap(err, updateError), domainErrors.RepositoryError)
		return nil, appErr
	}

	if len(book.Genres) == 1 && book.Genres[0] == "" {
		return book, nil
	}

	for _, genre := range book.Genres {
		bookGenreEntity := &BookGenre{
			BookID:    book.ID,
			GenreCode: genre,
		}

		sqlUpdate = "UPDATE BookGenre SET GenreCode = ? WHERE BookId = ?"
		_, err = s.db.Exec(sqlUpdate, bookGenreEntity.GenreCode, bookGenreEntity.BookID)
		if err != nil {
			appErr := domainErrors.NewAppError(errors.Wrap(err, updateError), domainErrors.RepositoryError)
			return nil, appErr
		}
	}

	return book, nil
}

func (s *Store) DeleteBook(id int) error {
	sqlDelete := "UPDATE Book SET Status = ? WHERE Id = ?"
	_, err := s.db.Exec(sqlDelete, "I", id)
	if err != nil {
		appErr := domainErrors.NewAppError(errors.Wrap(err, deleteError), domainErrors.RepositoryError)
		return appErr
	}

	return nil
}
