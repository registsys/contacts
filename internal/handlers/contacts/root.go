package contacts

import (
	"net/http"

	"github.com/registsys/contacts/internal/handlers"
	"github.com/registsys/contacts/internal/services"
)

type ContactsHandler struct {
	services *services.Services
	Handler  http.Handler
}

func NewContactsHandler(s *services.Services) *ContactsHandler {
	return &ContactsHandler{
		services: s,
	}
}

func (h *ContactsHandler) ContactsHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodPost:
		h.ContactCreateHandler(w, r)
	case http.MethodGet:
		h.ContactListHandler(w, r)
	default:
		handlers.ErrMethodNotAllowed(w, r)
	}
}
