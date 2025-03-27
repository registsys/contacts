package mux

import (
	"net/http"
)

type Deps struct {
	ContactCreateHandler http.HandlerFunc
	ContactListHandler   http.HandlerFunc
}

func New(d *Deps) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/contact", d.ContactCreateHandler)
	mux.HandleFunc("/contact/list", d.ContactListHandler)
	return mux
}
