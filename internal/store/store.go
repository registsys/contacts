package store

import "fmt"

type Storage interface {
	Add(contact Contact) error
	List() []Contact
}

type Contact struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
	Email string `json:"email"`
}

type Store map[string]Contact

func NewStore() Store {
	return Store{}
}

func (s Store) Add(contact Contact) error {

	if contact.Name == "" {
		return fmt.Errorf("name is required")
	}

	if contact.Phone == "" {
		return fmt.Errorf("phone is required")
	}

	if contact.Email == "" {
		return fmt.Errorf("email is required")
	}

	if _, ok := s[contact.Name]; ok {
		return fmt.Errorf("contact with name %s already exists", contact.Name)
	}
	s[contact.Name] = contact
	return nil
}

func (s Store) List() []Contact {
	contacts := make([]Contact, 0, len(s))
	for _, contact := range s {
		contacts = append(contacts, contact)
	}
	return contacts
}
