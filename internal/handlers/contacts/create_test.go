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
func TestNewContactCreateHandler_201Created(t *testing.T) {

	// TODO мокаем поведение метода Create таким образом что бы мы получили успешний результат

}

// TestNewContactCreateHandler_MethodNotAllowed тестирует получение ошибки
// в случае когда используется не валидный http метод
func TestNewContactCreateHandler_MethodNotAllowed(t *testing.T) {

	// TODO делаем запрос с невалидным http методом и ожидаем получение соотвествующей ошибки

}

// TestNewContactCreateHandler_InternalServerError тестирует обработку 500 статуса
func TestNewContactCreateHandler_InternalServerError(t *testing.T) {

	// TODO мокаем поведение метода Create таким образом что он вернул ошбику, в данном случае без разницы какую
	// и ожидаетм получить 500 ошибку

}
