package handlers

import (
	"net/http"

	"github.com/gorilla/pat"
)

func GetRouter() *pat.Router {
	r := pat.New()

	r.Add("OPTIONS", "/v1/items", http.HandlerFunc(HomeHandler))
	r.Add("OPTIONS", "/v1/item/{upc}", http.HandlerFunc(HomeHandler))

	r.Get("/v1/items", GetItemsHandler)

	r.Get("/v1/item/{upc}", GetItemInfoHandler)

	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./web")))

	r.Get("/", HomeHandler)

	return r
}
