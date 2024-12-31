package id

import "github.com/google/uuid"

// Nil is a zero value for the ID type
var Nil = ID(uuid.Nil)

// ID is a type that represents a unique identifier
type ID uuid.UUID

// New generates a new unique identifier
func New() ID {
	return ID(uuid.New())
}

// String returns the string representation of the identifier
func (id ID) String() string {
	return uuid.UUID(id).String()
}

// ParseID parses a string into an ID
func ParseID(s string) (ID, error) {
	u, err := uuid.Parse(s)
	if err != nil {
		return ID{}, err
	}

	return ID(u), nil
}
