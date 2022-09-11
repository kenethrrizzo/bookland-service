package books

type BookRepository interface {
	GetBookByID(int) (*Book, error)
	GetAllBooks() ([]Book, error)
	CreateBook(*Book) (*Book, error)
}