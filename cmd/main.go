package main

import (
	"github.com/registsys/contacts/internal/handlers/contacts"
	"github.com/registsys/contacts/internal/store"
	"log"
	"net/http"

	"github.com/registsys/contacts/internal/mux"
)

func main() {

	s := store.NewStore()

	// где то здесь может быть объявлен консьюмер или какой нибудь воркер которые на вход принимают store
	// так же обычно есть еще слой бизнес логики, он тоже может быть одинаково исопльован консьюмером и HTTP/GRPC хендлерами

	d := mux.Deps{
		ContactCreateHandler: contacts.NewContactCreateHandler(s),
		ContactListHandler:   contacts.NewContactListHandler(s),
	}
	err := http.ListenAndServe(`:8080`, mux.New(&d)) // на самом деле так не делается
	if err != nil {
		log.Fatal(err) // никогда не паникуем
	}
}
