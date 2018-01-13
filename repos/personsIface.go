package repos

import (
	"github.com/nu7hatch/gouuid"
	"github.com/x1m3/Tertulia/model"
)

type IPersonsCRUD interface {
	Add(*model.Person) error
	Get(*uuid.UUID) (*model.Person, error)
	Update(*model.Person) error
	Delete(*model.Person) error
}
