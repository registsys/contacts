package mux

import (
	"net/http"

	"github.com/registsys/contacts/internal/handlers/contacts"
	"github.com/registsys/contacts/internal/services"
)

func New(s *services.Services) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/contact", contacts.NewContactCreateHandler(&s.Contacts))
	mux.HandleFunc("/contact/list", contacts.NewContactListHandler(&s.Contacts))
	return mux
}
