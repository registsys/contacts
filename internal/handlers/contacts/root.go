package contacts

import (
	"net/http"

	"github.com/registsys/contacts/internal/handlers"
	"github.com/registsys/contacts/internal/services"
)

type ContactsHandler struct {
	Services ServicesI
	Handler  http.Handler
}

type ServicesI interface {
	ContactCreate(c services.Contact) error
	ContactList() []services.Contact
}

func NewContactsHandler(s ServicesI) *ContactsHandler {
	return &ContactsHandler{
		Services: s,
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
