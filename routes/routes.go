package routes

import (
	"net/http"
	"log"
	"go/go-lang-api/handlers"
	"go/go-lang-api/middleware"
	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.Use(middleware.ContentTypeMiddleware)
	r.HandleFunc("/", handlers.Home)
	r.HandleFunc("/api/personalidades", handlers.Personalidades).Methods("GET")
	r.HandleFunc("/api/personalidades/{id}", handlers.Personalidade).Methods("GET")
	r.HandleFunc("/api/personalidades", handlers.CreatePersonalidade).Methods("POST")
	r.HandleFunc("/api/personalidades/{id}", handlers.DeletePersonalidade).Methods("DELETE")
	r.HandleFunc("/api/personalidades/{id}", handlers.EditPersonalidade).Methods("PUT")
	log.Panic(http.ListenAndServe(":8080", r))
}