package dbentities

import (
	"github.com/nu7hatch/gouuid"
	"time"
	"github.com/x1m3/Tertulia/model"
)

type Topic struct {
	Id           uuid.UUID `storm:"id"`
	Title        string
	Body         string
	Author       uuid.UUID
	CreationDate time.Time `storm:"index"`
	ModDate      time.Time `storm:"index"`
}

func NewTopic(topic *model.Topic) *Topic {
	return &Topic{
		Id:           *topic.ID(),
		Title:        topic.Title(),
		Body:         topic.Body(),
		Author:       *topic.Author().ID(),
		CreationDate: topic.CreationDate(),
		ModDate:      topic.ModDate(),
	}
}

func (t *Topic) Topic() *model.Topic {
	topic := model.NewTopic(&t.Id)
	topic.SetTitle(t.Title)
	topic.SetBody(t.Body)
	topic.SetCreationDate(t.CreationDate)
	topic.SetModDate(t.ModDate)
	topic.SetAuthor(nil)
	return topic
}
