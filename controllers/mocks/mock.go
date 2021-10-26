package mocks

import (
	"bookusecase/entity"
)

type ServiceMock struct {}

func (s *ServiceMock) Create(product *entity.Book) error {
	return nil
}

func (s *ServiceMock) Get() ([]*entity.Book, error) {
	return []*entity.Book{{
		Name: "Brida",
		Author: "Paulo Coelho",
	}}, nil
}
