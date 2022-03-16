package book

import (
	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-3-oguzhantasimaz/internal/models/author"
	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-3-oguzhantasimaz/internal/models/book"
	"gorm.io/gorm"
)

type bookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) book.BookRepository {
	return &bookRepository{
		db: db,
	}
}

func (b *bookRepository) GetAllBooks() ([]*book.Book, error) {

}
func (b *bookRepository) GetBookByID(id int) (*book.Book, error) {

}
func (b *bookRepository) CreateBook(title string, stockCode string, stockCount int, isbn int, pageCount int, price float64, isDeleted bool, author *author.Author) (*book.Book, error) {

}
func (b *bookRepository) UpdateBook(id int, title string, stockCode string, stockCount int, isbn int, pageCount int, price float64, isDeleted bool, author *author.Author) (*book.Book, error) {

}
func (b *bookRepository) DeleteBook(id int) error {

}
