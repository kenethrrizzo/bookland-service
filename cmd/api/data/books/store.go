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
	sqlSelect := "SELECT Id, Name, Author, CoverPage, Synopsis, Price, CreatedAt, UpdatedAt FROM Book"

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

	var booksDomain []books.Book

	for res.Next() {
		var bookRetrieved Book

		err := res.Scan(&bookRetrieved.Id, &bookRetrieved.Name, &bookRetrieved.Author, &bookRetrieved.CoverPage, &bookRetrieved.Synopsis, &bookRetrieved.Price, &bookRetrieved.CreatedAt, &bookRetrieved.UpdatedAt)
		if err != nil {
			appErr := domainErrors.NewAppErrorWithType(domainErrors.MapError)
			return nil, appErr
		}

		booksDomain = append(booksDomain, *toDomainModel(&bookRetrieved))
	}

	return booksDomain, nil
}

// TODO: Refactorizar código duplicado
func (s *Store) GetBooksByGenre(genre string) ([]books.Book, error) {
	sqlSelect := `SELECT b.Id, b.Name, b.Author, b.CoverPage, b.Synopsis, b.Price, b.CreatedAt, b.UpdatedAt FROM Book b
		INNER JOIN BookGenre bg ON bg.BookId = b.Id AND bg.GenreCode = ?`

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

	var booksDomain []books.Book

	for res.Next() {
		var bookRetrieved Book

		err := res.Scan(&bookRetrieved.Id, &bookRetrieved.Name, &bookRetrieved.Author, &bookRetrieved.CoverPage, &bookRetrieved.Synopsis, &bookRetrieved.Price, &bookRetrieved.CreatedAt, &bookRetrieved.UpdatedAt)
		if err != nil {
			appErr := domainErrors.NewAppErrorWithType(domainErrors.MapError)
			return nil, appErr
		}

		booksDomain = append(booksDomain, *toDomainModel(&bookRetrieved))
	}

	return booksDomain, nil
}

func (s *Store) GetBookByID(id int) (*books.Book, error) {
	sqlSelect := "SELECT Id, Name, Author, CoverPage, Synopsis, Price, CreatedAt, UpdatedAt FROM Book WHERE Id = ?"

	var bookRetrieved Book

	err := s.db.QueryRow(sqlSelect, id).Scan(&bookRetrieved.Id, &bookRetrieved.Name, &bookRetrieved.Author, &bookRetrieved.CoverPage, &bookRetrieved.Synopsis, &bookRetrieved.Price, &bookRetrieved.CreatedAt, &bookRetrieved.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			appErr := domainErrors.NewAppErrorWithType(domainErrors.NotFound)
			return nil, appErr
		}
		appErr := domainErrors.NewAppError(errors.Wrap(err, readError), domainErrors.RepositoryError)
		return nil, appErr
	}

	return toDomainModel(&bookRetrieved), nil
}

func (s *Store) CreateBook(book *books.Book) (*books.Book, error) {
	sqlInsert := "INSERT INTO Book (Name, Author, CoverPage, Synopsis, Price) values (?, ?, ?, ?, ?)"

	bookEntity := toDBModel(book)

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

	book.Id = int(lastId)

	return book, nil
}

func (s *Store) UpdateBook(book *books.Book) (*books.Book, error) {
	sqlUpdate := "UPDATE Book SET Name = ?, Author = ?, CoverPage = ?, Synopsis = ?, Price = ?, UpdatedAt = ? WHERE Id = ?"

	bookEntity := toDBModel(book)

	_, err := s.db.Exec(sqlUpdate, bookEntity.Name, bookEntity.Author, bookEntity.CoverPage, bookEntity.Synopsis, bookEntity.Price, time.Now().UTC(), bookEntity.Id)
	if err != nil {
		appErr := domainErrors.NewAppError(errors.Wrap(err, updateError), domainErrors.RepositoryError)
		return nil, appErr
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
