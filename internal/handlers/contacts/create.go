package contacts

import (
	"encoding/json"
	"net/http"

	"github.com/registsys/contacts/internal/handlers"
	"github.com/registsys/contacts/internal/services"
)

// TODO для того что бы можно было протестировать хендлер, нужно объявить здесь свой интефейс
// и передать его в качестве аргумента вместо *services.ContactsService
func (h *ContactsHandler) ContactCreateHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		handlers.ErrMethodNotAllowed(w, r)
		return
	}

	var contact services.Contact

	err := json.NewDecoder(r.Body).Decode(&contact)
	if err != nil {
		handlers.ErrBadRequest(w, r, err.Error())
		return
	}
	if err := r.Body.Close(); err != nil {
		handlers.ErrInternalServerError(w, r, err.Error())
		return
	}

	err = h.Services.ContactCreate(contact)
	if err != nil {
		handlers.ErrBadRequest(w, r, err.Error())
		return
	}

	response, err := json.Marshal(contact)
	if err != nil {
		handlers.ErrInternalServerError(w, r, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}
