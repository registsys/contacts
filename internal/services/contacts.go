package services

import (
	"fmt"

	"github.com/registsys/contacts/internal/storage"
)

type (
	Contact = storage.Contact

	ContactPublisher func(contact Contact) error

	ContactsService struct {
		storage   storage.StorageI
		sendEvent ContactPublisher
	}
)

func NewContactsService(s storage.StorageI, sendEvent ContactPublisher) ContactsService {
	return ContactsService{
		storage:   s,
		sendEvent: sendEvent,
	}
}

func (s ContactsService) Create(contact Contact) error {

	if contact.Name == "" {
		return fmt.Errorf("name is required")
	}

	if contact.Phone == "" {
		return fmt.Errorf("phone is required")
	}

	if contact.Email == "" {
		return fmt.Errorf("email is required")
	}

	err := s.storage.Create(contact)
	if err != nil {
		return fmt.Errorf("failed to create contact: %w", err)
	}

	err = s.sendEvent(contact)
	if err != nil {
		return fmt.Errorf("failed to send event: %w", err)
	}

	return nil
}

func (s ContactsService) List() []Contact {
	return s.storage.List()
}
