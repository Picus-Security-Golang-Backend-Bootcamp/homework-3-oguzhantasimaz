package main

import (
	"os"

	log "github.com/sirupsen/logrus"

	infrastructure "github.com/Picus-Security-Golang-Backend-Bootcamp/homework-3-oguzhantasimaz/internal/infrastructure/repositories"
	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-3-oguzhantasimaz/internal/infrastructure/repositories/author"
	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-3-oguzhantasimaz/internal/infrastructure/repositories/book"
	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-3-oguzhantasimaz/internal/models/authors"
	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-3-oguzhantasimaz/internal/models/books"
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.WarnLevel)
}

func start() (books.BookRepository, authors.AuthorRepository) {
	db, err := infrastructure.NewMySQLDB("root:Ot123456@tcp(127.0.0.1:3306)/homework3?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Fatalln(err)
	}
	bookRepo := book.NewBookRepository(db)
	authorRepo := author.NewAuthorRepository(db)
	bookRepo.Migration()
	authorRepo.Migration()
	bookRepo.InsertSampleData()
	authorRepo.InsertSampleData()

	return bookRepo, authorRepo
}

func main() {
	bookRepo, authorRepo := start()

	authorList, err := authors.GetAllAuthors(authorRepo)
	if err != nil {
		log.Errorln(err)
	}

	for _, author := range authorList {
		log.Info(author)
	}

	bookList, err := books.GetAllBooks(bookRepo)
	if err != nil {
		log.Errorln(err)
	}

	for _, book := range bookList {
		log.Info(book)
	}

}
