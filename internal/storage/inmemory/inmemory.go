package inmemory

import (
	"fmt"

	"github.com/registsys/contacts/internal/errors"
	"github.com/registsys/contacts/internal/storage"
)

type (
	Contact = storage.Contact

	inMemoryStorage map[string]Contact
)

func New() inMemoryStorage {
	return make(inMemoryStorage)
}

func (s inMemoryStorage) Create(contact Contact) error {
	pk := contact.GetPK()

	if _, ok := s[pk]; ok {
		return fmt.Errorf("%w: contact with primary key %q already exists", errors.ErrObjectExists, pk)
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
