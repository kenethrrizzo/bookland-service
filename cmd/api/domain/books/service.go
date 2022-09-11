package books

type BookService interface {
	GetBookByID(int) (*Book, error)
	GetAllBooks() ([]Book, error)
	RegisterNewBook(*Book) (*Book, error)
}

type Service struct {
	repo BookRepository
}

func NewService(repo BookRepository) *Service {
	return &Service{repo}
}

func (svc *Service) GetBookByID(id int) (*Book, error) {
	return svc.repo.GetBookByID(id)
}

func (svc *Service) GetAllBooks() ([]Book, error) {
	return svc.repo.GetAllBooks()
}

func (svc *Service) RegisterNewBook(book *Book) (*Book, error) {
	return svc.repo.CreateBook(book)
}
