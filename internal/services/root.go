package services

import "github.com/registsys/contacts/internal/storage"

type Services struct {
	Contacts ContactsService
}

func NewServices(storage storage.StorageI, sendEvent ContactPublisher) *Services {
	return &Services{
		Contacts: NewContactsService(storage, sendEvent),
	}
}
