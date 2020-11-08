package dbentities

import (
	"github.com/nu7hatch/gouuid"
	"github.com/x1m3/Tertulia/backend/internal/lostandfound/model"
	"time"
)

type Person struct {
	Id               uuid.UUID `storm:"id"`
	Nickname         string    `storm:"unique"`
	Email            string
	RegistrationDate time.Time
}

func NewPerson(person *model.Person) *Person {
	return &Person{
		Id:               *person.ID(),
		Nickname:         person.Nickname(),
		Email:            person.Email().String(),
		RegistrationDate: person.RegistrationDate(),
	}
}
