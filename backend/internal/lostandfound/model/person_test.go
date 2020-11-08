package model

import (
	"github.com/nu7hatch/gouuid"
	"strconv"
	"testing"
	"time"
)

func TestPerson(t *testing.T) {

	id, _ := uuid.NewV4()
	p := NewPerson(id)
	p.SetNickname("xime")
	p.SetEmail("me@domain.com")
	p.SetRegistrationDate(time.Now())

	if p.Nickname() != "xime" {
		t.Errorf("Bad nickname. Expecting <xime>, got <%s>", p.Nickname())
	}

	now := time.Now()
	if now.Sub(p.RegistrationDate()) > 1*time.Millisecond {
		t.Errorf("Creation date is wrong (Or your server is a potato...)Expecting <%v>, got <%v>", now, p.RegistrationDate())
	}
}

func TestAddTopicsToPerson(t *testing.T) {

	id, _ := uuid.NewV4()
	p := NewPerson(id)
	p.SetNickname("xime")

	for i := 0; i < 1000; i++ {
		id, _ := uuid.NewV4()
		t := NewTopic(id)
		t.SetTitle("Title" + strconv.Itoa(i))
		t.SetBody("Body" + strconv.Itoa(i))
		t.SetAuthor(p)
		p.AddTopic(t)
	}

	topics := p.Topics()
	if len(topics) != 1000 {
		t.Errorf("Bad topics count related to Person. Expecting <1000>, got <%d>", len(topics))
	}

	for i := 0; i < 1000; i++ {
		title := topics[i].Title()
		body := topics[i].Body()
		author := topics[i].Author()

		if title != "Title"+strconv.Itoa(i) {
			t.Errorf("Bad topics title related to Person. Expecting <%s>, got <%s>", "Title"+strconv.Itoa(i), title)
		}
		if body != "Body"+strconv.Itoa(i) {
			t.Errorf("Bad topics body related to Person. Expecting <%s>, got <%s>", "Body"+strconv.Itoa(i), body)
		}
		if p != author {
			t.Errorf("Bad topics author related to Person. Expecting <%s>, got <%s>", p.Nickname(), author.Nickname())
		}
	}
}
