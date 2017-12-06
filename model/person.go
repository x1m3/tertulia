package model

import (
	"github.com/nu7hatch/gouuid"
	"net/mail"
	"time"
)

type Person struct {
	id               uuid.UUID
	nickname         string
	email            *mail.Address
	registrationDate time.Time
	topics           TopicList
}

func NewPerson(id *uuid.UUID) *Person {
	return &Person{id: *id}
}

func (p *Person) ID() *uuid.UUID {
	return &p.id
}

func (p *Person) SetNickname(nickname string) {
	p.nickname = nickname
}

func (p *Person) Nickname() string {
	return p.nickname
}

func (p *Person) SetEmail(email string) error {
	var err error
	p.email, err = mail.ParseAddress(email)
	return err
}

func (p *Person) Email() *mail.Address {
	return p.email
}

func (p *Person) SetRegistrationDate(t time.Time) {
	p.registrationDate = t
}

func (p *Person) RegistrationDate() time.Time {
	return p.registrationDate
}

func (p *Person) Topics() TopicList {
	return p.topics
}

func (p *Person) AddTopic(topic *Topic) {
	p.topics = append(p.topics, topic)
}
