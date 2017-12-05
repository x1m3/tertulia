package repos

import (
	"github.com/asdine/storm"
	"github.com/nu7hatch/gouuid"
	"time"
	"github.com/x1m3/Tertulia/model"
)

type topicsStormDTO struct {
	Id    *uuid.UUID `storm:"id"`
	Title string
	Body  string
	//author       uuid.UUID
	CreationDate time.Time `storm:"index"`
	ModDate      time.Time `storm:"index"`
}

func NewTopicStormDTO(topic *model.Topic) *topicsStormDTO {
	return &topicsStormDTO{
		Id:           topic.ID(),
		Title:        topic.Title(),
		Body:         topic.Body(),
		CreationDate: topic.CreationDate(),
		ModDate:      topic.ModDate(),
	}
}

func (t *topicsStormDTO) Topic() *model.Topic {
	topic := model.NewTopic(t.Id)
	topic.SetTitle(t.Title)
	topic.SetBody(t.Body)
	topic.SetCreationDate(t.CreationDate)
	topic.SetModDate(t.ModDate)
	topic.SetAuthor(nil)
	return topic
}

type TopicsStorm struct {
	db *storm.DB
}

func NewTopicsStorm(db *storm.DB) *TopicsStorm {
	return &TopicsStorm{db: db}
}

func (r *TopicsStorm) Add(topic *model.Topic) error {
	return r.db.Save(NewTopicStormDTO(topic))
}

func (r *TopicsStorm) Get(id *uuid.UUID) (*model.Topic, error) {
	var dto topicsStormDTO

	err := r.db.One("Id", id, &dto)
	if err != nil {
		return nil, err
	}
	return dto.Topic(), nil
}

func (r *TopicsStorm) Update(topic *model.Topic) error {
	return r.db.Update(NewTopicStormDTO(topic))
}

func (r *TopicsStorm) GetByCreatedDateDesc(from int, limit int) []*model.Topic {
	return r.getByIndexDescend("CreationDate", from, limit)
}

func (r *TopicsStorm) GetByUpdatedDateDesc(from int, limit int) []*model.Topic {
	return r.getByIndexDescend("ModDate", from, limit)
}

func (r *TopicsStorm) getByIndexDescend(index string, from int, limit int) []*model.Topic {
	var topicsDTO []topicsStormDTO
	var err error

	if limit == 0 {
		err = r.db.AllByIndex(index, &topicsDTO, storm.Reverse(), storm.Skip(from))
	} else {
		err = r.db.AllByIndex(index, &topicsDTO, storm.Reverse(), storm.Skip(from), storm.Limit(limit))
	}
	if err != nil {
		return nil
	}

	topics := make([]*model.Topic, 0, len(topicsDTO))

	for _, t := range topicsDTO {
		topics = append(topics, t.Topic())
	}

	return topics
}
