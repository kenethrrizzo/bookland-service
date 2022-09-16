package books

type BookRepository interface {
	GetAllBooks() ([]Book, error)
	GetBooksByGenre(string) ([]Book, error)
	GetBookByID(int) (*Book, error)
	CreateBook(*Book) (*Book, error)
	UpdateBook(*Book) (*Book, error)
	DeleteBook(int) error
}
