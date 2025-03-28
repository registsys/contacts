package services

type Services struct {
	Contacts ContactsService
}

func NewServices() *Services {
	return &Services{
		Contacts: NewContactsService(),
	}
}
