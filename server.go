package main

import (
	"github.com/hiennguyen1610/api-yelp/handler"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func main() {
	mux := httprouter.New()
	mux.GET("/search", handler.Log(handler.Search))

	server := http.Server{
		Addr: "127.0.0.1:8080",
		Handler: mux,
	}
	server.ListenAndServe()
}