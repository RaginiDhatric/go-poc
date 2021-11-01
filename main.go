package main

import (
	"bookusecase/controllers"
	"bookusecase/router"
	"bookusecase/service"
)

var (
	svc               service.BookService        = service.NewBookService()
	productController controllers.BookController = controllers.NewBookController(svc)
	muxRouter         router.Router              = router.NewMuxRouter()
)

func main() {
	muxRouter.Get("/book", productController.GetProduct)
	muxRouter.Post("/book", productController.CreateProduct)
	muxRouter.Serve(":9091")
}
