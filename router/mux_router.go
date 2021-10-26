package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

type muxRouter struct{}

var (
	muxDispatcher = mux.NewRouter()
)

func NewMuxRouter() Router {
	return &muxRouter{}
}

func (m *muxRouter) Get(url string, f func(rw http.ResponseWriter, r *http.Request)) {
	muxDispatcher.HandleFunc(url, f).Methods("GET")
}

func (m *muxRouter) Post(url string, f func(rw http.ResponseWriter, r *http.Request)) {
	muxDispatcher.HandleFunc(url, f).Methods("POST")
}

func (m *muxRouter) Serve(port string) {
	http.ListenAndServe(port, muxDispatcher)
}



// var RegisterRoutes = func(router *mux.Router) {
// 	b := &Book{}
// 	router.HandleFunc("/book/", b.CreateProduct).Methods("POST")
// 	router.HandleFunc("/book/", b.GetProduct).Methods("GET")
// }
