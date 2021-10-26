package repo

import "bookusecase/entity"

var bookList = []*entity.Book{
	{
		Name: "Brida",
		Author:   "Paulo Coelho",
	},
	{
		Name: "2 States",
		Author:   "Chetan Bhagat",
	},
	{
		Name: "The Alchemist",
		Author:   "Paulo Coelho",
	},
}

type repo struct{}

func NewInMemBookRepo() BookRepository {
	return &repo{}
}

func (*repo) Get() ([]*entity.Book, error) {
	return bookList, nil
}

func (*repo) Create(book *entity.Book) error {
	bookList = append(bookList, book)
	return nil
}