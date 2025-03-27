package store

import (
	"testing"
)

func TestStore(t *testing.T) {

	t.Run("add new contact", func(t *testing.T) {
		store := NewStore()
		contact := Contact{
			Name:  "John Doe",
			Phone: "1234567890",
			Email: "john.doe@example.com",
		}

		err := store.Add(contact)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}

		if len(store) != 1 {
			t.Errorf("expected 1 contact, got %d", len(store))
		}
	})

	t.Run("add exists contact", func(t *testing.T) {
		store := NewStore()
		contact := Contact{
			Name:  "John Doe",
			Phone: "1234567890",
			Email: "john.doe@example.com",
		}

		store.Add(contact)

		duplicated := Contact{
			Name:  "John Doe",
			Phone: "9876543210",
			Email: "duplicated@example.com",
		}

		err := store.Add(duplicated)
		if err == nil {
			t.Errorf("expected error, got nil")
		}

		if len(store) != 1 {
			t.Errorf("expected 1 contact, got %d", len(store))
		}
	})

	t.Run("new contact name required", func(t *testing.T) {
		store := NewStore()
		contact := Contact{
			Phone: "1234567890",
			Email: "john.doe@example.com",
		}

		err := store.Add(contact)
		if err == nil {
			t.Errorf("expected error, got nil")
		}

		if len(store) != 0 {
			t.Errorf("expected 0 contact, got %d", len(store))
		}
	})

	t.Run("new contact phone required", func(t *testing.T) {
		store := NewStore()
		contact := Contact{
			Name:  "John Doe",
			Email: "john.doe@example.com",
		}

		err := store.Add(contact)
		if err == nil {
			t.Errorf("expected error, got nil")
		}

		if len(store) != 0 {
			t.Errorf("expected 0 contact, got %d", len(store))
		}
	})

	t.Run("new contact email required", func(t *testing.T) {
		store := NewStore()
		contact := Contact{
			Name:  "John Doe",
			Phone: "1234567890",
		}

		err := store.Add(contact)
		if err == nil {
			t.Errorf("expected error, got nil")
		}

		if len(store) != 0 {
			t.Errorf("expected 0 contact, got %d", len(store))
		}
	})

	t.Run("contacts list", func(t *testing.T) {
		store := NewStore()
		contact1 := Contact{
			Name:  "John Doe",
			Phone: "1234567890",
			Email: "john.doe@example.com",
		}

		store.Add(contact1)

		contact2 := Contact{
			Name:  "John Doe brother",
			Phone: "9876543210",
			Email: "duplicated@example.com",
		}

		store.Add(contact2)

		selected := store.List()

		if len(selected) != 2 {
			t.Errorf("expected 2 contact, got %d", len(selected))
		}
	})
}
