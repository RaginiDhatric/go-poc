package repo

import (
	"bookusecase/entity"
)

type BookRepository interface{
	Get() ([]*entity.Book, error)
	Create(*entity.Book) error
}