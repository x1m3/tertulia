package repos

import (
	"github.com/asdine/storm"
	"github.com/nu7hatch/gouuid"
	"github.com/x1m3/Tertulia/model"
	"github.com/x1m3/Tertulia/repos/dbentities"
)

type TopicsStorm struct {
	db *storm.DB
}

func NewTopicsStorm(db *storm.DB) *TopicsStorm {
	return &TopicsStorm{db: db}
}

func (r *TopicsStorm) Add(topic *model.Topic) error {
	return r.db.Save(dbentities.NewTopic(topic))
}

func (r *TopicsStorm) Get(id *uuid.UUID) (*model.Topic, error) {
	var dto dbentities.Topic

	err := r.db.One("Id", id, &dto)
	if err != nil {
		return nil, err
	}
	return dto.Topic(), nil
}

func (r *TopicsStorm) Update(topic *model.Topic) error {
	return r.db.Update(dbentities.NewTopic(topic))
}

func (r *TopicsStorm) GetByCreatedDateDesc(from int, limit int) []*model.Topic {
	return r.getByIndexDescend("CreationDate", from, limit)
}

func (r *TopicsStorm) GetByUpdatedDateDesc(from int, limit int) []*model.Topic {
	return r.getByIndexDescend("ModDate", from, limit)
}

func (r *TopicsStorm) getByIndexDescend(index string, from int, limit int) []*model.Topic {
	var topicsDTO []dbentities.Topic
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
