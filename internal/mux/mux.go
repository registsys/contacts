package mux

import (
	"net/http"

	"github.com/registsys/contacts/internal/handlers/contacts"
	"github.com/registsys/contacts/internal/services"
)

func New(s *services.Services) *http.ServeMux {

	mux := http.NewServeMux()

	contactsH := contacts.NewContactsHandler(s)

	mux.HandleFunc("/contacts", contactsH.ContactsHandler)

	return mux
}
