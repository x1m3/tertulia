package model

import (
	"github.com/nu7hatch/gouuid"
)

type IPersonsCRUD interface {
	Add(*Person) error
	Get(*uuid.UUID) (*Person, error)
	Update(*Person) error
	Delete(*Person) error
}
