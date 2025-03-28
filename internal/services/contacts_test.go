package services

import (
	"testing"

	"github.com/registsys/contacts/internal/config"
	"github.com/registsys/contacts/internal/storage"
)

func TestStore(t *testing.T) {

	cfg := config.New("")

	t.Run("add new contact", func(t *testing.T) {
		store := storage.New(cfg.PostgresDSN)
		s := NewContactsService(store)
		contact := Contact{
			Name:  "John Doe",
			Phone: "1234567890",
			Email: "john.doe@example.com",
		}

		err := s.Create(contact)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}

		items := s.List()

		if len(items) != 1 {
			t.Errorf("expected 1 contact, got %d", len(items))
		}
	})

	t.Run("add exists contact", func(t *testing.T) {
		store := storage.New(cfg.PostgresDSN)
		s := NewContactsService(store)
		contact := Contact{
			Name:  "John Doe",
			Phone: "1234567890",
			Email: "john.doe@example.com",
		}

		s.Create(contact)

		duplicated := Contact{
			Name:  "John Doe",
			Phone: "9876543210",
			Email: "duplicated@example.com",
		}

		err := s.Create(duplicated)
		if err == nil {
			t.Errorf("expected error, got nil")
		}

		items := s.List()

		if len(items) != 1 {
			t.Errorf("expected 1 contact, got %d", len(items))
		}
	})

	t.Run("new contact name required", func(t *testing.T) {
		store := storage.New(cfg.PostgresDSN)
		s := NewContactsService(store)
		contact := Contact{
			Phone: "1234567890",
			Email: "john.doe@example.com",
		}

		err := s.Create(contact)
		if err == nil {
			t.Errorf("expected error, got nil")
		}

		items := s.List()

		if len(items) != 0 {
			t.Errorf("expected 0 contact, got %d", len(items))
		}
	})

	t.Run("new contact phone required", func(t *testing.T) {
		store := storage.New(cfg.PostgresDSN)
		s := NewContactsService(store)
		contact := Contact{
			Name:  "John Doe",
			Email: "john.doe@example.com",
		}

		err := s.Create(contact)
		if err == nil {
			t.Errorf("expected error, got nil")
		}

		items := s.List()

		if len(items) != 0 {
			t.Errorf("expected 0 contact, got %d", len(items))
		}
	})

	t.Run("new contact email required", func(t *testing.T) {
		store := storage.New(cfg.PostgresDSN)
		s := NewContactsService(store)
		contact := Contact{
			Name:  "John Doe",
			Phone: "1234567890",
		}

		err := s.Create(contact)
		if err == nil {
			t.Errorf("expected error, got nil")
		}

		items := s.List()

		if len(items) != 0 {
			t.Errorf("expected 0 contact, got %d", len(items))
		}
	})

	t.Run("contacts list", func(t *testing.T) {
		store := storage.New(cfg.PostgresDSN)
		s := NewContactsService(store)
		contact1 := Contact{
			Name:  "John Doe",
			Phone: "1234567890",
			Email: "john.doe@example.com",
		}

		s.Create(contact1)

		contact2 := Contact{
			Name:  "John Doe brother",
			Phone: "9876543210",
			Email: "duplicated@example.com",
		}

		s.Create(contact2)

		selected := s.List()

		if len(selected) != 2 {
			t.Errorf("expected 2 contact, got %d", len(selected))
		}
	})
}
