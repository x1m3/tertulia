package repos_test

import (
	"github.com/asdine/storm"
	"github.com/nu7hatch/gouuid"
	"github.com/x1m3/Tertulia/model"
	"github.com/x1m3/Tertulia/repos"
	"os"
	"testing"
	"time"
)

func TestPersonsRepo(t *testing.T) {
	db, err := storm.Open("test.db")
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		db.Close()
		os.Remove("test.db")
	}()

	repo := repos.NewPersonsStorm(db)

	persons := getSomeAuthors(50)

	// Add all persons to repo
	for _, person := range persons {
		err := repo.Add(person)
		if err != nil {
			t.Error(err)
		}
	}
	// Let's check if all persons have been saved
	for _, expected := range persons {
		got, err := repo.Get(expected.ID())
		if err != nil {
			t.Error(err)
		}
		if got.ID().String() != expected.ID().String() {
			t.Error("ID differs")
		}
		if got.Nickname() != expected.Nickname() {
			t.Error("Nickname differs")
		}
		if got.Email().String() != expected.Email().String() {
			t.Error("Email differs")
		}
		if got.RegistrationDate().Sub(expected.RegistrationDate()) != 0 {
			t.Errorf("Registration date differs. Got <%v>, expected <%v>", got.RegistrationDate(), expected.RegistrationDate())
		}
	}

	// Let's update all records
	for _, p := range persons {
		p.SetNickname(p.Nickname() + "lala")
		p.SetEmail("lala" + p.Email().String())
		err := repo.Update(p)
		if err != nil {
			t.Errorf("Error updating record <%s>", err)
		}
	}

	// Let's check if all persons are updated
	for _, expected := range persons {
		got, err := repo.Get(expected.ID())
		if err != nil {
			t.Error(err)
		}
		if got.ID().String() != expected.ID().String() {
			t.Error("ID differs")
		}
		if got.Nickname() != expected.Nickname() {
			t.Error("Nickname differs")
		}
		if got.Email().String() != expected.Email().String() {
			t.Error("Email differs")
		}
		if got.RegistrationDate().Sub(expected.RegistrationDate()) != 0 {
			t.Errorf("Registration date differs. Got <%v>, expected <%v>", got.RegistrationDate(), expected.RegistrationDate())
		}
	}
	// Let's remove all
	for _, p := range persons {
		err := repo.Delete(p)
		if err != nil {
			t.Errorf("Error removing record <%s>", err)
		}
	}
	// Let's check if all persons are removed
	for _, person := range persons {
		got, err := repo.Get(person.ID())
		if err == nil {
			t.Error(err)
		}
		if got != nil {
			t.Error("Expecting a nil value")
		}
	}
}

func getSomeAuthors(n int) []*model.Person {
	persons := make([]*model.Person, 0, n)
	for i := 0; i < n; i++ {
		id, _ := uuid.NewV4()
		person := model.NewPerson(id)
		person.SetNickname(id.String())
		person.SetEmail(id.String() + "@tertulia.com")
		person.SetRegistrationDate(time.Now())

		persons = append(persons, person)
	}
	return persons
}
