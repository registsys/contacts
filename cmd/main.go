package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/registsys/contacts/internal/config"
	"github.com/registsys/contacts/internal/mux"
	"github.com/registsys/contacts/internal/services"
	"github.com/registsys/contacts/internal/storage"
	"github.com/registsys/contacts/internal/storage/inmemory"
	"github.com/registsys/contacts/internal/storage/pg"
)

func main() {

	cfg := config.New("./config.yaml")

	var store storage.StorageI

	store, err := pg.New(cfg.PostgresDSN)
	if err != nil {
		if err == pg.ErrNotConfigured {
			fmt.Printf("failed to create pg store: %v\nusing inmemory store\n", err)
			store = inmemory.New()
		} else {
			log.Fatal(err)
		}
	}

	serv := services.NewServices(store)

	err = http.ListenAndServe(`:8080`, mux.New(serv))
	if err != nil {
		log.Fatal(err)
	}
}
