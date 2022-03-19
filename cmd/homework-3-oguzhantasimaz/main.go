package main

import (
	"fmt"
	"os"
	"path"
	"strconv"

	"github.com/onsi/ginkgo/reporters/stenographer/support/go-colorable"
	log "github.com/sirupsen/logrus"

	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-3-oguzhantasimaz/helper"
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
	args := os.Args

	if len(args) == 1 {
		projectName := path.Base(args[0])
		fmt.Printf("\n %s uygulamasında kullanabileceğiniz komutlar : \n search => arama işlemi için \n list => listeleme işlemi için\n buy => satin alma islemi icin\n delete => silme islemi icin", projectName)
		os.Exit(1)
	} else {
		var err error
		args = args[1:]
		if args[0] == "search" {
			var book *books.Book

			//check argument to see if it is an int or string
			if helper.IsInt(args[1]) {
				//declare error variable
				book, err = books.GetBookByID(bookRepo, helper.StringToInt(args[1]))
				if err != nil {
					log.Errorln(err)
				}
			} else {
				book, err = books.GetBookByTitle(bookRepo, args[1])
				if err != nil {
					log.Errorln(err)
				}
			}
			//if book is not found
			if book != nil {
				book.Print()
				author, err := authors.GetAuthorByID(authorRepo, book.AuthorID)
				if err != nil {
					log.Fatal(err)
				}
				author.Print()
			} else {
				fmt.Println("\nBook not found")
			}
		} else if args[0] == "list" {
			bookList, err := books.GetAllBooks(bookRepo)
			if err != nil {
				log.Errorln(err)
			}
			for _, book := range bookList {
				book.Print()
				author, err := authors.GetAuthorByID(authorRepo, book.AuthorID)
				if err != nil {
					log.Fatal(err)
				}
				author.Print()
			}
		} else if args[0] == "buy" {
			if len(args) >= 2 {
				//get id and count args as integer and buy book by id and count
				id, err := strconv.Atoi(args[1])
				if err != nil {
					//create error if id is not integer
					error := fmt.Errorf("%s is not integer", args[1])
					log.Fatal(error)
				}
				count, err := strconv.Atoi(args[2])
				if err != nil {
					//create error if count is not integer
					error := fmt.Errorf("%s is not integer", args[2])
					log.Fatal(error)
				}
				//buy book by id and count
				book, err := books.BuyBook(bookRepo, id, count)
				if err != nil {
					log.Fatal(err)
				}
				fmt.Printf("\n Bought book: \n")
				book.Print()
				author, err := authors.GetAuthorByID(authorRepo, book.AuthorID)
				if err != nil {
					log.Fatal(err)
				}
				author.Print()
			} else {
				log.Fatal("\n Please enter book id and count")
			}
		} else if args[0] == "delete" {
			//delete book by id
			if len(args) >= 2 {
				//convert id to integer
				id, err := strconv.Atoi(args[1])
				if err != nil {
					//create error if id is not integer
					error := fmt.Errorf("%s is not integer", args[1])
					log.Fatal(error)
				}
				err = books.DeleteBook(bookRepo, id)
				if err != nil {
					log.Errorln(err)
				}
				fmt.Println("\nDeleted book id: " + args[1])
			} else {
				log.Fatal("\n Please enter book id")
			}
		}
	}

}
