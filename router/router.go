package router

import "net/http"

type Router interface {
	Get(uri string, f func(rw http.ResponseWriter, r *http.Request))
	Post(uri string, f func(rw http.ResponseWriter, r *http.Request))
	Serve(string)
}
