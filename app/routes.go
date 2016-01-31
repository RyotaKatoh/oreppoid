package app

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func BuildRouter() http.Handler {
	publicRouter := mux.NewRouter().PathPrefix("api").Subrouter()
	registerPublicRoutes(publicRouter)

	return publicRouter

}

func registerPublicRoutes(r *mux.Router) {

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { fmt.Println("hogehoge") }).Methods("GET")

}
