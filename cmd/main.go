package main

import (
	"log"
	"net/http"

	"github.com/registsys/contacts/internal/mux"
	"github.com/registsys/contacts/internal/services"
)

func main() {

	s := services.NewServices()

	err := http.ListenAndServe(`:8080`, mux.New(s))
	if err != nil {
		log.Fatal(err)
	}
}
