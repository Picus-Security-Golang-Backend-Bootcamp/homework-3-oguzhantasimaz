package book

import (
	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-3-oguzhantasimaz/internal/models/authors"
	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-3-oguzhantasimaz/internal/models/books"
	"gorm.io/gorm"

	log "github.com/sirupsen/logrus"
)

type bookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) books.BookRepository {
	return &bookRepository{
		db: db,
	}
}

func (b *bookRepository) Migration() {
	b.db.Migrator().DropTable(&books.Book{})
	b.db.AutoMigrate(&books.Book{})
}

func (b *bookRepository) GetAllBooks() ([]*books.Book, error) {
	books := make([]*books.Book, 0)
	err := b.db.Find(&books).Error
	return books, err
}
func (b *bookRepository) GetBookByID(id int) (*books.Book, error) {
	book := new(books.Book)
	err := b.db.Where("id = ?", id).First(book).Error
	return book, err
}
func (b *bookRepository) GetBookByTitle(title string) (*books.Book, error) {
	book := new(books.Book)
	//parameter title includes wildcard character '%'
	err := b.db.Where("title LIKE ?", "%"+title+"%").First(book).Error
	return book, err
}
func (b *bookRepository) CreateBook(book *books.Book) (*books.Book, error) {
	err := b.db.Create(book).Error
	return book, err
}
func (b *bookRepository) UpdateBook(book *books.Book) (*books.Book, error) {
	err := b.db.Save(book).Error
	return book, err
}
func (b *bookRepository) DeleteBook(id int) error {
	//update book isDeleted field to true
	book, err := b.GetBookByID(id)
	if err != nil {
		return err
	}
	book.IsDeleted = true
	book, err = b.UpdateBook(book)
	if err != nil {
		return err
	}
	return err
}
func (b *bookRepository) BuyBook(id int, count int) (*books.Book, error) {
	book, err := b.GetBookByID(id)
	if err != nil {
		return nil, err
	}
	if book.StockCount < count {
		log.Fatal("Not enough stock")
	}
	book.StockCount -= count
	book, err = b.UpdateBook(book)
	if err != nil {
		return nil, err
	}
	return book, err
}

func (r *bookRepository) InsertSampleData() {
	log.Info("Inserting sample data")
	bookList := []books.Book{
		{Title: "Beyaz Ev", StockCode: "GPRG", StockCount: 10, Isbn: 123456789, PageCount: 200, Price: 10.0, IsDeleted: false, Author: authors.Author{Name: "Erkut", Surname: "Tackin"}},
		{Title: "Resimdeki Gozyasi", StockCode: "CPRG", StockCount: 10, Isbn: 123456788, PageCount: 200, Price: 10.0, IsDeleted: false, Author: authors.Author{Name: "Cem", Surname: "Karaca"}},
		{Title: "Hayat Bir Teselli", StockCode: "JPRG", StockCount: 10, Isbn: 123456787, PageCount: 200, Price: 10.0, IsDeleted: false, Author: authors.Author{Name: "Erkin", Surname: "Koray"}},
		{Title: "Unutamadim", StockCode: "PPRG", StockCount: 10, Isbn: 123456786, PageCount: 200, Price: 10.0, IsDeleted: false, Author: authors.Author{Name: "Baris", Surname: "Manco"}},
		{Title: "Amasra", StockCode: "RPRG", StockCount: 10, Isbn: 123456785, PageCount: 200, Price: 10.0, IsDeleted: false, Author: authors.Author{Name: "Baris", Surname: "Akarsu"}},
	}
	r.db.Create(&bookList)
}
