package storage

import (
	"database/sql"
	"fmt"
)

type Contact struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
	Email string `json:"email"`
}

type StorageI interface {
	Create(contact Contact) error
	List() []Contact
}

type pgStore struct {
	db *sql.DB
}

type inMemoryStorage map[string]Contact

func New(dsn string) StorageI {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return make(inMemoryStorage)
	}

	return &pgStore{db}
}
func (c Contact) GetPK() string {
	return c.Name
}

func (s inMemoryStorage) Create(contact Contact) error {
	pk := contact.GetPK()

	if _, ok := s[pk]; ok {
		return fmt.Errorf("contact with primary key %s already exists", pk)
	}

	s[pk] = contact
	return nil
}

func (s inMemoryStorage) List() []Contact {
	contacts := make([]Contact, 0, len(s))
	for _, contact := range s {
		contacts = append(contacts, contact)
	}
	return contacts
}

func (s pgStore) Create(contact Contact) error {

	return nil
}

func (s pgStore) List() []Contact {
	return []Contact{}
}
