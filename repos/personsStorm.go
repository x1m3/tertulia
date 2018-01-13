package repos

import (
	"github.com/asdine/storm"
	"github.com/nu7hatch/gouuid"
	"github.com/x1m3/Tertulia/model"
	"github.com/x1m3/Tertulia/repos/dbentities"
)

type PersonsStorm struct {
	db *storm.DB
}

func NewPersonsStorm(db *storm.DB) *PersonsStorm {
	return &PersonsStorm{db: db}
}

func (r *PersonsStorm) Add(p *model.Person) error {
	// TODO: Return a repo generic error
	return r.db.Save(dbentities.NewPerson(p))
}

func (r *PersonsStorm) Get(id *uuid.UUID) (*model.Person, error) {
	p := &dbentities.Person{}
	err := r.db.One("Id", id, p)
	switch err {
	case nil:
		person := model.NewPerson(&p.Id)
		person.SetNickname(p.Nickname)
		person.SetRegistrationDate(p.RegistrationDate)
		err = person.SetEmail(p.Email)
		return person, err
	case storm.ErrNotFound:
		// TODO: Return a repo error, not a model one
		return nil, model.ErrNotFound
	default:
		return nil, err
	}
}

func (r *PersonsStorm) Update(person *model.Person) error {
	err := r.db.Update(dbentities.NewPerson(person))
	switch err {
	case nil:
		return nil
	case storm.ErrNotFound:
		// TODO: Return a repo error, not a model one
		return model.ErrNotFound
	default:
		return err
	}
}

func (r *PersonsStorm) Delete(person *model.Person) error {
	err := r.db.DeleteStruct(dbentities.NewPerson(person))
	switch err {
	case nil:
		return nil
	case storm.ErrNotFound:
		// TODO: Return a repo error, not a model one
		return model.ErrNotFound
	default:
		return err
	}
}
