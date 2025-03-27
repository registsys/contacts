package mux

import (
	"net/http"

	"github.com/registsys/contacts/internal/handlers/contacts"
	"github.com/registsys/contacts/internal/store"
)

func New() *http.ServeMux {
	mux := http.NewServeMux()
	store := store.NewStore()

	mux.HandleFunc("/contact", contacts.NewContactCreateHandler(store))
	mux.HandleFunc("/contact/list", contacts.NewContactListHandler(store))
	return mux
}
