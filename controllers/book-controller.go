package controllers

import (
	"bookusecase/entity"
	"bookusecase/service"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type BookController interface {
	GetProduct(http.ResponseWriter, *http.Request)
	CreateProduct(http.ResponseWriter, *http.Request)
}

type controller struct {
	svc service.BookService
}

func NewBookController(svc service.BookService) BookController {
	return &controller{
		svc: svc,
	}
}

func (c controller) GetProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	products, err := c.svc.Get()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	res, _ := json.Marshal(products)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func (c controller) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var newProduct *entity.Book
	resp, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(resp, &newProduct)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = c.svc.Create(newProduct)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message": "Created product successfully!!"}`))
}
