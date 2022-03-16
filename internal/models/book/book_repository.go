package book

import "github.com/Picus-Security-Golang-Backend-Bootcamp/homework-3-oguzhantasimaz/internal/models/author"

type BookRepository interface {
	GetAllBooks() ([]*Book, error)
	GetBookByID(id int) (*Book, error)
	CreateBook(title string, stockCode string, stockCount int, isbn int, pageCount int, price float64, isDeleted bool, author *author.Author) (*Book, error)
	UpdateBook(id int, title string, stockCode string, stockCount int, isbn int, pageCount int, price float64, isDeleted bool, author *author.Author) (*Book, error)
	DeleteBook(id int) error
}
