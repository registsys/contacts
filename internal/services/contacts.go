package services

import (
	"fmt"

	"github.com/registsys/contacts/internal/storage"
)

type (
	Contact = storage.Contact

	ContactPublisher func(contact Contact) error

	Services struct {
		store            storage.StorageI
		onContactCreated ContactPublisher
	}
)

func NewServices(store storage.StorageI, onContactCreated ContactPublisher) *Services {
	return &Services{
		store:            store,
		onContactCreated: onContactCreated,
	}
}

func (s Services) ContactCreate(contact Contact) error {

	if contact.Name == "" {
		return fmt.Errorf("name is required")
	}

	if contact.Phone == "" {
		return fmt.Errorf("phone is required")
	}

	if contact.Email == "" {
		return fmt.Errorf("email is required")
	}

	err := s.store.Create(contact)
	if err != nil {
		return err
	}

	err = s.onContactCreated(contact)
	if err != nil {
		return fmt.Errorf("failed to send event: %w", err)
	}

	return nil
}

func (s Services) ContactList() []Contact {
	return s.store.List()
}
