package services

import (
	"fmt"

	"github.com/registsys/contacts/internal/storage"
)

type Contact = storage.Contact

type ContactsService struct {
	storage storage.StorageI
}

func NewContactsService(s storage.StorageI) ContactsService {
	return ContactsService{s}
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

	return s.storage.Create(contact)
}

func (s ContactsService) List() []Contact {
	return s.storage.List()
}
