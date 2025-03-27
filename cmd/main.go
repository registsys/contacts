package main

import (
	"net/http"

	"github.com/registsys/contacts/internal/mux"
)

func main() {

	err := http.ListenAndServe(`:8080`, mux.New())
	if err != nil {
		panic(err)
	}
}
