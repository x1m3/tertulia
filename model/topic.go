package model

import (
	"github.com/nu7hatch/gouuid"
	"github.com/Sirupsen/logrus"
	"time"
	"github.com/x1m3/Tertulia/utils/zstrings"
)

type Topic struct {
	id           uuid.UUID
	title        string
	body         *zstrings.ZString
	author       *Person
	creationDate time.Time
	modDate      time.Time
}

type TopicList []*Topic

func NewTopic() *Topic {
	uuid, err := uuid.NewV4()
	if err != nil {
		logrus.Errorf("Error generating UUID for Topic : <%s>", err)
	}
	now := time.Now()
	return &Topic{
		id:*uuid,
		creationDate: now,
		modDate:now,
	}
}

func (t *Topic) ID() uuid.UUID {
	return t.id
}

func (t *Topic) SetTitle(title string) {
	t.title = title
	t.modDate = time.Now()
}

func (t *Topic) SetBody(body string) {
	t.body = zstrings.NewZStringCompressed(body)
	t.modDate = time.Now()
}

func (t *Topic) Title() string {
	return t.title
}

func (t *Topic) Body() string {
	return t.body.Value()
}

func (t *Topic) CreationDate() time.Time {
	return t.creationDate
}

func (t *Topic) ModDate() time.Time {
	return t.modDate
}

func (t *Topic) SetAuthor(p *Person) {
	t.author = p
}

func (t *Topic) Author() *Person {
	return t.author
}