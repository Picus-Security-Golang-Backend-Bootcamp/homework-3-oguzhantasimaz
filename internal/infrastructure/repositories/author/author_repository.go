package author

import (
	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-3-oguzhantasimaz/internal/models/authors"
	"gorm.io/gorm"
)

type authorRepository struct {
	db *gorm.DB
}

func NewAuthorRepository(db *gorm.DB) authors.AuthorRepository {
	return &authorRepository{
		db: db,
	}
}

func (a *authorRepository) Migration() {
	a.db.Migrator().DropTable(&authors.Author{})
	a.db.AutoMigrate(&authors.Author{})
}

func (a *authorRepository) GetAllAuthors() ([]*authors.Author, error) {
	authors := make([]*authors.Author, 0)
	err := a.db.Find(&authors).Error
	return authors, err
}

func (a *authorRepository) GetAuthorByID(id int) (*authors.Author, error) {
	author := new(authors.Author)
	err := a.db.Where("id = ?", id).First(author).Error
	return author, err
}

func (a *authorRepository) CreateAuthor(author *authors.Author) (*authors.Author, error) {
	err := a.db.Create(author).Error
	return author, err
}

func (a *authorRepository) UpdateAuthor(author *authors.Author) (*authors.Author, error) {
	err := a.db.Save(author).Error
	return author, err
}

func (a *authorRepository) DeleteAuthor(id int) (err error) {
	err = a.db.Where("id = ?", id).Delete(&authors.Author{}).Error
	return
}

func (a *authorRepository) GetAuthorByName(name string) (*authors.Author, error) {
	author := new(authors.Author)
	err := a.db.Where("name = ?", name).First(author).Error
	return author, err
}

func (a *authorRepository) InsertSampleData() {
	authorList := []*authors.Author{
		{Name: "Erkut", Surname: "Tackin"},
		{Name: "Cem", Surname: "Karaca"},
		{Name: "Erkin", Surname: "Koray"},
		{Name: "Baris", Surname: "Manco"},
		{Name: "Baris", Surname: "Akarsu"},
	}
	for _, auth := range authorList {
		a.db.Create(&authors.Author{Name: auth.Name, Surname: auth.Surname})
	}
}
