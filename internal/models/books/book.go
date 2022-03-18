package books

import (
	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-3-oguzhantasimaz/internal/models/authors"
	"gorm.io/gorm"

	log "github.com/sirupsen/logrus"
)

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
	log.Infof("%s | %s | %s | %f", b.Title, b.Author.Name, b.StockCode, b.Price)
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

func UpdateBook(r BookRepository, book *Book) (*Book, error) {
	return r.UpdateBook(book)
}

func DeleteBook(r BookRepository, id int) error {
	return r.DeleteBook(id)
}
