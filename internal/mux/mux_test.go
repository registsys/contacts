package mux

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMux(t *testing.T) {

	mux := New() // Если в этом конструкторе будет будут собираьтся все слои, то мы не сможе нормально писать тесты
	if mux == nil {
		t.Error("mux is nil")
	}

	existsRouteTests := []struct {
		method string
		route  string
	}{
		{
			method: http.MethodPost,
			route:  "/contact",
		},
		{
			method: http.MethodGet,
			route:  "/contact/list",
		},
	}

	for _, test := range existsRouteTests {
		t.Run(fmt.Sprintf("%s %s", test.method, test.route), func(t *testing.T) {
			r := httptest.NewRequest(test.method, test.route, nil)
			w := httptest.NewRecorder()

			mux.ServeHTTP(w, r)

			if w.Result().StatusCode == http.StatusNotFound {
				t.Errorf("endpoint %s not found in mux", test.route)
			}
		})
	}
}
