package dbentities

import (
	"github.com/nu7hatch/gouuid"
	"github.com/x1m3/Tertulia/model"
	"time"
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
