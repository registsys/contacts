package contacts

import (
	"encoding/json"
	"net/http"

	"github.com/registsys/contacts/internal/handlers"
)

func (h *ContactsHandler) ContactListHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		handlers.ErrMethodNotAllowed(w, r)
		return
	}

	contacts := h.services.ContactList()

	response, err := json.Marshal(contacts)
	if err != nil {
		handlers.ErrInternalServerError(w, r, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
