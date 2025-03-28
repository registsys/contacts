package main

import (
	"log"
	"net/http"

	"github.com/registsys/contacts/internal/config"
	"github.com/registsys/contacts/internal/mux"
	"github.com/registsys/contacts/internal/services"
)

func main() {

	// ВАЖНО: вызов этого конструкторв должен быть здесь
	cfg := config.New("путь к файлу")
	if cfg.PostgresDSN != "" {
		// TODO получаем подключение к БД
	}
	// если подключения к БД нет, то используюем в качестве хранилища памятьq

	s := services.NewServices()

	err := http.ListenAndServe(`:8080`, mux.New(s))
	if err != nil {
		log.Fatal(err)
	}
}
