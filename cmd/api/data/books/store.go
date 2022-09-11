package books

import (
	"database/sql"
	"log"
	"strconv"

	"github.com/kenethrrizzo/bookland-service/cmd/api/domain/books"
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
			log.Println("no se ha encontrado un registro en la base de datos", err)
			return nil, err
			// TODO: Implementar errores personalizados
		}
		log.Fatalln("ha ocurrido un error al obtener libros: ", err)
		return nil, err
		// TODO: Implementar errores personalizados
	}
	defer res.Close()

	var booksRes []books.Book

	for res.Next() {
		log.Println("obteniendo registro...")
		var br Book // br: books retrieved
		err := res.Scan(&br.Id, &br.Name, &br.Author, &br.CoverPage, &br.Synopsis, &br.Price, &br.CreatedAt, &br.UpdatedAt)
		if err != nil {
			log.Println("ha ocurrido un error al mapear el resultado a la entidad: ", err)
			return nil, err
			// TODO: Implementar errores personalizados
		}
		booksRes = append(booksRes, *toDomainModel(&br))
	}
	log.Println("cantidad de libros recuperados: " + strconv.Itoa(len(booksRes)))

	return booksRes, nil
}

func (s *Store) GetBookByID(id int) (*books.Book, error) {
	sqlSelect := "SELECT Id, Name, Author, CoverPage, Synopsis, Price, CreatedAt, UpdatedAt FROM Book WHERE Id = ?"

	var br Book // br: books retrieved
	err := s.db.QueryRow(sqlSelect, id).Scan(&br.Id, &br.Name, &br.Author, &br.CoverPage, &br.Synopsis, &br.Price, &br.CreatedAt, &br.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("no se ha encontrado un registro en la base de datos", err)
			return nil, err
			// TODO: Implementar errores personalizados
		}
		log.Fatalln("ha ocurrido un error al obtener un libro por id: ", err)
		return nil, err
		// TODO: Implementar errores personalizados
	}
	log.Println("libro recuperado: " + br.Name)

	return toDomainModel(&br), nil
}

func (s *Store) CreateBook(book *books.Book) (*books.Book, error) {
	return nil, nil
}
