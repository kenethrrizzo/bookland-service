package books

import (
	"github.com/kenethrrizzo/bookland-service/cmd/api/domain/files"
)

type BookService interface {
	GetBookByID(int) (*Book, error)
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

func (svc *Service) GetBookByID(id int) (*Book, error) {
	return svc.bookRepo.GetBookByID(id)
}

func (svc *Service) GetAllBooks() ([]Book, error) {
	return svc.bookRepo.GetAllBooks()
}

func (svc *Service) RegisterNewBook(book *Book) (*Book, error) {
	if book.CoverPage != "" {
		coverPageURL, err := svc.fileRepo.UploadFile(book.CoverPage)
		if err != nil {
			return nil, err
		}

		book.CoverPage = *coverPageURL
	}

	return svc.bookRepo.CreateBook(book)
}

func (svc *Service) UpdateBook(book *Book, bookID int) (*Book, error) {
	book.Id = bookID

	if book.CoverPage != "" {
		coverPageURL, err := svc.fileRepo.UploadFile(book.CoverPage)
		if err != nil {
			return nil, err
		}

		book.CoverPage = *coverPageURL
	}

	return svc.bookRepo.UpdateBook(book)
}

func (svc *Service) DeleteBook(id int) error {
	return svc.bookRepo.DeleteBook(id)
}
