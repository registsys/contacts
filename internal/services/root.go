package services

import "github.com/registsys/contacts/internal/storage"

type Services struct {
	Contacts ContactsService
}

func NewServices(storage storage.StorageI) *Services {
	return &Services{
		Contacts: NewContactsService(storage),
	}
}
