package storage

type Contact struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
	Email string `json:"email"`
}

type StorageI interface {
	Create(contact Contact) error
	List() []Contact
}

func (c Contact) GetPK() string {
	return c.Name
}
