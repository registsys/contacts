package mux

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/registsys/contacts/internal/services"
)

func TestMux(t *testing.T) {

	s := services.NewServices()
	mux := New(s)
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
