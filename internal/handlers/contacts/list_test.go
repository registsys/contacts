package contacts

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/registsys/contacts/internal/services"
)

func TestListHandler(t *testing.T) {
	expected := `[{"name":"John Doe","phone":"1234567890","email":"john.doe@example.com"}]`
	var contacts []services.Contact
	json.Unmarshal([]byte(expected), &contacts)

	req := httptest.NewRequest(http.MethodGet, "/contact/list", nil)
	w := httptest.NewRecorder()

	s := services.NewServices()
	s.Contacts.Create(contacts[0])

	NewContactListHandler(&s.Contacts)(w, req)

	resp := w.Result()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal("Error reading response body:", err)
	}

	if resp.Body != nil {
		resp.Body.Close()
	}

	if string(body) != expected {
		t.Errorf("Expected body %q, got %q", expected, string(body))
	}
}
