package services

import "fmt"

type Contact struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
	Email string `json:"email"`
}

func (c Contact) GetPK() string {
	return c.Name
}

func validate(contact Contact) error {

	if contact.Name == "" {
		return fmt.Errorf("name is required")
	}

	if contact.Phone == "" {
		return fmt.Errorf("phone is required")
	}

	if contact.Email == "" {
		return fmt.Errorf("email is required")
	}

	return nil
}

type ContactsService map[string]Contact

func NewContactsService() ContactsService {
	return make(ContactsService)
}

func (s ContactsService) Create(contact Contact) error {

	err := validate(contact)
	if err != nil {
		return err
	}

	pk := contact.GetPK()

	if _, ok := s[pk]; ok {
		return fmt.Errorf("contact with primary key %s already exists", pk)
	}

	s[pk] = contact
	return nil
}

func (s ContactsService) List() []Contact {

	contacts := make([]Contact, 0, len(s))
	for _, contact := range s {
		contacts = append(contacts, contact)
	}
	return contacts
}
