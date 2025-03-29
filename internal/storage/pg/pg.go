package pg

import (
	"database/sql"

	"github.com/registsys/contacts/internal/storage"
)

type (
	Contact = storage.Contact

	pgStore struct {
		db *sql.DB
	}
)

func New(dsn string) (*pgStore, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	return &pgStore{db: db}, nil
}

func (s pgStore) Create(contact Contact) error {
	return nil
}

func (s pgStore) List() []Contact {
	return []Contact{}
}
