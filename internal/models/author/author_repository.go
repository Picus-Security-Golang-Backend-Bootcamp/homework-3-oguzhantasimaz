package author

type AuthorRepository interface {
	GetAllAuthors() ([]*Author, error)
	GetAuthorByID(id int) (*Author, error)
	CreateAuthor(author *Author) (*Author, error)
	UpdateAuthor(author *Author) (*Author, error)
	DeleteAuthor(id int) error
}
