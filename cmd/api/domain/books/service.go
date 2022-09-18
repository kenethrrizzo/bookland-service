package books

import (
	"github.com/kenethrrizzo/bookland-service/cmd/api/domain/files"
)

type BookService interface {
	GetBookByID(int) (*Book, error)
	GetBooksByGenre(string) ([]Book, error)
	GetAllBooks() ([]Book, error)
	RegisterNewBook(*Book) (*Book, error)
	UpdateBook(*Book, int) (*Book, error)
	DeleteBook(int) error
}

type Service struct {
	bookRepo BookRepository
	fileRepo files.FileRepository
}

func NewService(bookRepo BookRepository, fileRepo files.FileRepository) *Service {
	return &Service{bookRepo, fileRepo}
}

func (s *Service) GetAllBooks() ([]Book, error) {
	return s.bookRepo.GetAllBooks()
}

func (s *Service) GetBooksByGenre(genre string) ([]Book, error) {
	return s.bookRepo.GetBooksByGenre(genre)
}

func (s *Service) GetBookByID(id int) (*Book, error) {
	return s.bookRepo.GetBookByID(id)
}

func (s *Service) RegisterNewBook(book *Book) (*Book, error) {
	validateGenre(book.Genres)

	if book.CoverPage != "" {
		coverPageURL, err := s.fileRepo.UploadFile(book.CoverPage)
		if err != nil {
			return nil, err
		}

		book.CoverPage = *coverPageURL
	}

	return s.bookRepo.CreateBook(book)
}

func (s *Service) UpdateBook(book *Book, bookID int) (*Book, error) {
	validateGenre(book.Genres)

	book.ID = bookID

	if book.CoverPage != "" {
		coverPageURL, err := s.fileRepo.UploadFile(book.CoverPage)
		if err != nil {
			return nil, err
		}

		book.CoverPage = *coverPageURL
	}

	return s.bookRepo.UpdateBook(book)
}

func (s *Service) DeleteBook(id int) error {
	return s.bookRepo.DeleteBook(id)
}
