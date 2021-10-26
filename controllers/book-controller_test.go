package controllers

import (
	"bookusecase/controllers/mocks"
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetProductWhenEmpty(t *testing.T) {
	// Test case 1 : When there are no products initially
	req, err := http.NewRequest(http.MethodGet, "/book", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	serviceMock := &mocks.ServiceMock{}
	c := NewBookController(serviceMock)
	handler := http.HandlerFunc(c.GetProduct)
	handler.ServeHTTP(rr, req)
	status := rr.Code
	if status == http.StatusNoContent {
		resp := "{\"message\": \"There are no products to display!!\"}"
		if rr.Body.String() !=  resp {
			t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), resp)
		}
	}
}

func TestGetProduct(t *testing.T) {
	// Test case 2 : When there are valid products
	req2, err := http.NewRequest(http.MethodGet, "/book/", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr2 := httptest.NewRecorder()
	serviceMock := &mocks.ServiceMock{}
	c := NewBookController(serviceMock)
	
	handler2 := http.HandlerFunc(c.GetProduct)
	handler2.ServeHTTP(rr2, req2)
	status2 := rr2.Code
	if status2 == http.StatusOK {
		expected := `[{"name":"Brida","author":"Paulo Coelho"}]`
		if rr2.Body.String() != expected {
			t.Errorf("handler returned unexpected body: got %v want %v",
				rr2.Body.String(), expected)
		}
	}
}

func TestCreateProduct(t *testing.T) {
	var jsonStr = []byte(`{"BookName":"Brida", "Author": "Paulo Coelho"}`)
	req2, err := http.NewRequest(http.MethodPost, "/book/", bytes.NewBuffer([]byte(jsonStr)))
	if err != nil {
		t.Fatal(err)
	}
	rr2 := httptest.NewRecorder()
	serviceMock := &mocks.ServiceMock{}
	c := NewBookController(serviceMock)

	handler2 := http.HandlerFunc(c.CreateProduct)
	handler2.ServeHTTP(rr2, req2)
	status2 := rr2.Code
	if status2 == http.StatusOK {
		expected := `[{"BookName": "Brida", "Author": "Paulo Coelho"}]`
		if rr2.Body.String() != expected {
			t.Errorf("handler returned unexpected body: got %v want %v",
				rr2.Body.String(), expected)
		}
	}
}

func TestCreateProductFailure(t *testing.T){
	// Test ioutil.ReadAll error scenario 
	testRequest := httptest.NewRequest(http.MethodPost, "/something", errReader(0))
	rr := httptest.NewRecorder()
	serviceMock := &mocks.ServiceMock{}
	c := NewBookController(serviceMock)
	handler := http.HandlerFunc(c.CreateProduct)
	handler.ServeHTTP(rr, testRequest)
	status := rr.Code
	if status == http.StatusBadRequest {
		fmt.Println("error case is ok")
	}
}

type errReader int

func (errReader) Read(p []byte) (n int, err error) {
	return 0, errors.New("test error")
}