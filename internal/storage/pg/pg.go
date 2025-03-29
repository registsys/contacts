package pg

import (
	"database/sql"
	"fmt"

	"github.com/registsys/contacts/internal/storage"
)

var (
	ErrNotConfigured       = fmt.Errorf("dsn not provided")
	ErrImpoperlyConfigured = fmt.Errorf("could not connect to db")
)

type (
	Contact = storage.Contact

	pgStore struct {
		db *sql.DB
	}
)

func New(dsn string) (*pgStore, error) {
	if dsn == "" {
		return nil, ErrNotConfigured
	}

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("%w\n\t%v", ErrImpoperlyConfigured, err)
	}
	return &pgStore{db: db}, nil
}

func (s pgStore) Create(contact Contact) error {
	return nil
}

func (s pgStore) List() []Contact {
	return []Contact{}
}
