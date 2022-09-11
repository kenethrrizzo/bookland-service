package books

import (
	"database/sql"
	"log"

	"github.com/kenethrrizzo/bookland-service/cmd/api/domain/books"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db}
}

func (s *Store) GetBookByID(id int) (*books.Book, error) {
	sqlSelect := "SELECT Id, Name, Author, CoverPage, Synopsis, Price, CreatedAt, UpdatedAt FROM Book WHERE Id = ?"

	var br Book
	err := s.db.QueryRow(sqlSelect, id).Scan(&br.Id, &br.Name, &br.Author, &br.CoverPage, &br.Synopsis, &br.Price, &br.CreatedAt, &br.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("no se ha encontrado un registro en la base de datos", err)
			return nil, err
			// TODO: Implementar errores personalizados
		}
		log.Fatalln("ha ocurrido un error al obtener un libro por id: ", err)
		return nil, err
	}
	log.Println("Libro recuperado: " + br.Name)

	return toDomainModel(&br), nil
}

func (s *Store) GetAllBooks() ([]books.Book, error) {
	return nil, nil
}

func (s *Store) CreateBook(book *books.Book) (*books.Book, error) {
	return nil, nil
}
