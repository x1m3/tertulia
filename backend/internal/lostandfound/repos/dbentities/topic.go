package dbentities

import (
	"github.com/nu7hatch/gouuid"
	"github.com/x1m3/Tertulia/backend/internal/lostandfound/model"
	"time"
)

type TopicDTO struct {
	Id           uuid.UUID `storm:"id"`
	Title        string
	Body         string
	Author       uuid.UUID
	CreationDate time.Time `storm:"index"`
	ModDate      time.Time `storm:"index"`
}

func (dto *TopicDTO) ToTopic() *model.Topic {
	t := model.NewTopic(&dto.Id)
	t.SetTitle(dto.Title)
	t.SetBody(dto.Body)
	t.SetCreationDate(t.CreationDate())
	t.SetModDate(t.ModDate())
	a := model.NewPerson(&dto.Author) // Only author id is set
	t.SetAuthor(a)
	return t
}

func NewTopicDTO(topic *model.Topic) *TopicDTO {
	return &TopicDTO{
		Id:           *topic.ID(),
		Title:        topic.Title(),
		Body:         topic.Body(),
		Author:       *topic.Author().ID(),
		CreationDate: topic.CreationDate(),
		ModDate:      topic.ModDate(),
	}
}
