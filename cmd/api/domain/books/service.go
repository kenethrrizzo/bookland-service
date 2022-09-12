package books

import "github.com/kenethrrizzo/bookland-service/cmd/api/domain/files"

type BookService interface {
	GetBookByID(int) (*Book, error)
	GetAllBooks() ([]Book, error)
	RegisterNewBook(*Book) (*Book, error)
	UpdateBookCoverImage(int, string) (*Book, error)
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
	return svc.bookRepo.CreateBook(book)
}

func (svc *Service) UpdateBookCoverImage(id int, newCoverPageURI string) (*Book, error) {
	return svc.bookRepo.UpdateBookCoverImage(id, newCoverPageURI)
}

func (svc *Service) DeleteBook(id int) error {
	return svc.bookRepo.DeleteBook(id)
}
