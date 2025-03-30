package contacts_test

import (
	"testing"
)

// TODO
// Общие услвия для всех тестов:
// - имя текущего пакета(contacts_test) остается не изменным
// - пакеты config и storage сюда импортироваться не должны
// - моки можно написать свои, но лучше заюзать готовый toolkit, посмотри как работают с https://github.com/stretchr/testify

// TestNewContactCreateHandler_201Created тестирует успешный сценарий
func TestContactCreateHandler_201Created(t *testing.T) {

	// TODO мокаем поведение метода Create таким образом что бы мы получили успешний результат

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
