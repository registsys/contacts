package contacts_test

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/registsys/contacts/internal/handlers/contacts"
	"github.com/registsys/contacts/internal/services"
	"github.com/stretchr/testify/mock"
)

type MockedServices struct {
	mock.Mock
}

func (m *MockedServices) ContactCreate(c services.Contact) error {
	m.Called(c)
	return nil
}

func (m *MockedServices) ContactList() []services.Contact {
	m.Called()
	return nil
}

// TODO
// Общие услвия для всех тестов:
// - имя текущего пакета(contacts_test) остается не изменным
// - пакеты config и storage сюда импортироваться не должны
// - моки можно написать свои, но лучше заюзать готовый toolkit, посмотри как работают с https://github.com/stretchr/testify

// TestNewContactCreateHandler_201Created тестирует успешный сценарий
func TestContactCreateHandler_201Created(t *testing.T) {

	contact := `{"name":"John Doe","phone":"1234567890","email":"john.doe@example.com"}`

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/contacts", bytes.NewReader([]byte(contact)))

	// TODO мокаем поведение метода Create таким образом что бы мы получили успешний результат
	s := new(MockedServices)

	h := &contacts.ContactsHandler{
		Services: s,
	}

	s.On("ContactCreate", services.Contact{
		Name:  "John Doe",
		Phone: "1234567890",
		Email: "john.doe@example.com",
	}).Return(nil)

	h.ContactCreateHandler(w, r)

	s.AssertExpectations(t)

	result := w.Result()

	if result.StatusCode != http.StatusCreated {
		t.Errorf("got status code %d, want status code %d", result.StatusCode, http.StatusCreated)
	}

	body, err := io.ReadAll(result.Body)
	if err != nil {
		t.Fatal("Error reading response body:", err)
	}

	if result.Body != nil {
		result.Body.Close()
	}

	if string(body) != contact {
		t.Errorf("got body %s, want body %s", string(body), contact)
	}
}

// TestNewContactCreateHandler_MethodNotAllowed тестирует получение ошибки
// в случае когда используется не валидный http метод
func TestNewContactCreateHandler_MethodNotAllowed(t *testing.T) {

	// for _, method := range []string{http.MethodPut, http.MethodPatch, http.MethodDelete} {
	// 	request := httptest.NewRequest(method, "/contacts", nil)
	// 	response := httptest.NewRecorder()
	// 	contacts.ContactCreateHandler(response, request)

	// 	result := response.Result()

	// 	if result.StatusCode != http.StatusMethodNotAllowed {
	// 		t.Errorf("got stattus code %d, want stattus code %d", result.StatusCode, http.StatusMethodNotAllowed)
	// 	}
	// }

}

// TestNewContactCreateHandler_InternalServerError тестирует обработку 500 статуса
func TestNewContactCreateHandler_InternalServerError(t *testing.T) {

	// TODO мокаем поведение метода Create таким образом что он вернул ошбику, в данном случае без разницы какую
	// и ожидаетм получить 500 ошибку

}
