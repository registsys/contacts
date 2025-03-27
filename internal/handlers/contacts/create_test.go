package contacts

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/registsys/contacts/internal/store"
)

func TestCreateHandler(t *testing.T) {

	requestBody := `{"name":"John Doe","phone":"1234567890","email":"john.doe@example.com"}`

	s := store.NewStore()

	req := httptest.NewRequest(http.MethodPost, "/contact", bytes.NewReader([]byte(requestBody)))
	w := httptest.NewRecorder()

	NewContactCreateHandler(s)(w, req)

	resp := w.Result()

	if resp.StatusCode != http.StatusCreated {
		t.Errorf("Expected status code %d, got %d", http.StatusCreated, resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal("Error reading response body:", err)
	}

	if resp.Body != nil {
		resp.Body.Close()
	}

	if string(body) != requestBody {
		t.Errorf("Expected body %q, got %q", requestBody, string(body))
	}

	contacts := s.List()
	if len(contacts) != 1 {
		t.Errorf("Expected 1 contact, got %d", len(contacts))
	}

}
