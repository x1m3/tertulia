package model

import (
	"github.com/nu7hatch/gouuid"
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

func NewTopic(id *uuid.UUID) *Topic {
	now := time.Now()
	return &Topic{
		id:*id,
		creationDate: now,
		modDate:now,
	}
}

func (t *Topic) ID() *uuid.UUID {
	return &t.id
}

func (t *Topic) SetTitle(title string) {
	t.title = title
	t.modDate = time.Now()
}

func (t *Topic) SetBody(body string) {
	t.body = zstrings.NewZStringCompressed(body)
	t.modDate = time.Now()
}

func (t *Topic) SetCreationDate(tim time.Time) {
	t.creationDate = tim
}

func (t *Topic) SetModDate(tim time.Time) {
	t.modDate = tim
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