package contacts

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/registsys/contacts/internal/services"
)

// TODO для того что бы можно было протестировать хендлер, нужно объявить здесь свой интефейс
// и передать его в качестве аргумента вместо *services.ContactsService
func NewContactCreateHandler(s *services.ContactsService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method != http.MethodPost {
			http.Error(w, fmt.Sprintf("Method %s not allowed", r.Method), http.StatusMethodNotAllowed)
			return
		}

		var contact services.Contact

		err := json.NewDecoder(r.Body).Decode(&contact)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if err := r.Body.Close(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		err = s.Create(contact)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		response, err := json.Marshal(contact)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(response)
	}
}
