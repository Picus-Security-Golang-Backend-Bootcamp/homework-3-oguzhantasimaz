package book

import "github.com/Picus-Security-Golang-Backend-Bootcamp/homework-3-oguzhantasimaz/internal/models/author"

type Book struct {
	id         int
	title      string
	stockCode  string
	stockCount int
	isbn       int
	pageCount  int
	price      float64
	isDeleted  bool
	author     *author.Author
}

func CreateBook(r *BookRepository, title string, stockCode string, stockCount int, isbn int, pageCount int, price float64, isDeleted bool, author *author.Author) *Book {
	return r.CreateBook(title, stockCode, stockCount, isbn, pageCount, price, isDeleted, author)
}
