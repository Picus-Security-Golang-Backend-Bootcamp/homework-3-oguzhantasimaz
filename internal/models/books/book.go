package books

import (
	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-3-oguzhantasimaz/internal/models/authors"
	"gorm.io/gorm"

	log "github.com/sirupsen/logrus"
)

// Book represents a book
type Book struct {
	gorm.Model
	ID         int
	Title      string
	StockCode  string
	StockCount int
	Isbn       int
	PageCount  int
	Price      float64
	IsDeleted  bool
	AuthorID   int
	Author     authors.Author `gorm:"references:ID"`
}

func (b *Book) Print() {
	log.Infof("\nBook:\n %d | %s | %s | %d | %d | %d | %f | %v ", b.ID, b.Title, b.StockCode, b.StockCount, b.Isbn, b.PageCount, b.Price, b.IsDeleted)
}

func CreateBook(r BookRepository, book *Book) (*Book, error) {
	return r.CreateBook(book)
}

func GetAllBooks(r BookRepository) ([]*Book, error) {
	return r.GetAllBooks()
}

func GetBookByID(r BookRepository, id int) (*Book, error) {
	return r.GetBookByID(id)
}

func GetBookByTitle(r BookRepository, title string) (*Book, error) {
	return r.GetBookByTitle(title)
}

func BuyBook(r BookRepository, id int, count int) (*Book, error) {
	return r.BuyBook(id, count)
}

func UpdateBook(r BookRepository, book *Book) (*Book, error) {
	return r.UpdateBook(book)
}

func DeleteBook(r BookRepository, id int) error {
	return r.DeleteBook(id)
}
