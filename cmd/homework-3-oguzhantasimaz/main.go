package main

import (
	"github.com/onsi/ginkgo/reporters/stenographer/support/go-colorable"
	log "github.com/sirupsen/logrus"

	infrastructure "github.com/Picus-Security-Golang-Backend-Bootcamp/homework-3-oguzhantasimaz/internal/infrastructure/repositories"
	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-3-oguzhantasimaz/internal/infrastructure/repositories/author"
	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-3-oguzhantasimaz/internal/infrastructure/repositories/book"
	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-3-oguzhantasimaz/internal/models/authors"
	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-3-oguzhantasimaz/internal/models/books"
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.TextFormatter{ForceColors: true})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(colorable.NewColorableStdout())

	// Only log the warning severity or above.
	log.SetLevel(log.TraceLevel)

}

func start() (books.BookRepository, authors.AuthorRepository) {
	db, err := infrastructure.NewMySQLDB("root:Ot123456@tcp(127.0.0.1:3306)/homework3?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Fatalln(err)
	}
	bookRepo := book.NewBookRepository(db)
	authorRepo := author.NewAuthorRepository(db)
	authorRepo.Migration()
	bookRepo.Migration()
	bookRepo.InsertSampleData()

	return bookRepo, authorRepo
}

func main() {
	bookRepo, authorRepo := start()

	authorList, err := authors.GetAllAuthors(authorRepo)
	if err != nil {
		log.Errorln(err)
	}

	for _, author := range authorList {
		author.Print()
	}

	bookList, err := books.GetAllBooks(bookRepo)
	if err != nil {
		log.Errorln(err)
	}

	for _, book := range bookList {
		book.Print()
	}

}
