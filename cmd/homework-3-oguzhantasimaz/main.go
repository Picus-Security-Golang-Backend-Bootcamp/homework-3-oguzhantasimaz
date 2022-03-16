package main

import "github.com/Picus-Security-Golang-Backend-Bootcamp/homework-3-oguzhantasimaz/internal/infrastructure/repositories/book"

func main() {
	newBook := book.NewBookRepository(db)
}
