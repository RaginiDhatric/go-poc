package service

import (
	"bookusecase/entity"
	"bookusecase/repo"
)

var(
	bookRepo repo.BookRepository = repo.NewMongoDBRepo()
)

type BookService interface {
	Get() ([]*entity.Book, error)
	Create(*entity.Book) error
}

type service struct{}

func NewBookService() BookService{
	return &service{}
}

func (*service) Get() ([]*entity.Book, error){
	return bookRepo.Get()
}

func (*service) Create(book *entity.Book) error{
	return bookRepo.Create(book)
}