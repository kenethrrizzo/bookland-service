package books

import (
	"log"

	"github.com/kenethrrizzo/bookland-service/cmd/api/domain/files"
	"github.com/sirupsen/logrus"
)

type BookService interface {
	GetBookByID(int) (*Book, error)
	GetAllBooks() ([]Book, error)
	RegisterNewBook(*Book, string) (*Book, error)
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

func (svc *Service) RegisterNewBook(book *Book, coverImg string) (*Book, error) {
	coverPageName, err := svc.fileRepo.UploadFile(coverImg)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	book.CoverPage = *coverPageName

	return svc.bookRepo.CreateBook(book)
}

func (svc *Service) UpdateBookCoverImage(id int, coverImg string) (*Book, error) {
	coverPageName, err := svc.fileRepo.UploadFile(coverImg)
	if err != nil {
		log.Fatalln(err)
	}

	return svc.bookRepo.UpdateBookCoverImage(id, *coverPageName)
}

func (svc *Service) DeleteBook(id int) error {
	return svc.bookRepo.DeleteBook(id)
}
