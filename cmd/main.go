package main

import (
	"log"
	"net/http"

	"github.com/registsys/contacts/internal/config"
	"github.com/registsys/contacts/internal/mux"
	"github.com/registsys/contacts/internal/services"
	"github.com/registsys/contacts/internal/storage"
)

func main() {

	cfg := config.New("./config.yaml")

	store := storage.New(cfg.PostgresDSN)

	serv := services.NewServices(store)

	err := http.ListenAndServe(`:8080`, mux.New(serv))
	if err != nil {
		log.Fatal(err)
	}
}
