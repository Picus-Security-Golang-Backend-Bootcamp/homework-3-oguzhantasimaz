package books

type BookRepository interface {
	Migration()
	InsertSampleData()
	GetAllBooks() ([]*Book, error)
	GetBookByID(id int) (*Book, error)
	CreateBook(book *Book) (*Book, error)
	UpdateBook(book *Book) (*Book, error)
	DeleteBook(id int) error
}
