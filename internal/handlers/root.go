package handlers

import (
	"fmt"
	"net/http"
)

func ErrMethodNotAllowed(w http.ResponseWriter, r *http.Request) {
	http.Error(w, fmt.Sprintf("Method %s not allowed", r.Method), http.StatusMethodNotAllowed)
}

func ErrInternalServerError(w http.ResponseWriter, r *http.Request, err string) {
	if err == "" {
		err = "Internal server error"
	}
	http.Error(w, "Internal server error", http.StatusInternalServerError)
}

func ErrBadRequest(w http.ResponseWriter, r *http.Request, err string) {
	if err == "" {
		err = "Bad request"
	}
	http.Error(w, err, http.StatusBadRequest)
}

func ErrNotFound(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Not found", http.StatusNotFound)
}

func ErrObjectExists(w http.ResponseWriter, r *http.Request, err string) {
	if err == "" {
		err = "Object already exists"
	}
	http.Error(w, err, http.StatusConflict)
}
