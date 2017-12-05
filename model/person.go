package model

import (
	"github.com/nu7hatch/gouuid"
	"time"
	"net/mail"
	"github.com/Sirupsen/logrus"
)

type Person struct {
	id               uuid.UUID
	nickname         string
	email            *mail.Address
	registrationDate time.Time
	topics           TopicList
}

func NewPerson(nickname string) *Person{
	uuid, err := uuid.NewV4()
	if err != nil {
		logrus.Errorf("Error generating UUID for Topic : <%s>", err)
	}
	now := time.Now()
	return &Person{
		id:*uuid,
		nickname: nickname,
		registrationDate: now,
	}
}

func (p *Person) ID() *uuid.UUID {
	return &p.id
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

func (p *Person) RegistrationDate() time.Time {
	return p.registrationDate
}

func (p *Person) Topics() TopicList {
	return p.topics
}

func (p *Person) AddTopic(topic *Topic) {
	p.topics = append(p.topics, topic)
}

