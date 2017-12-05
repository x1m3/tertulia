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

func (r *TopicsStorm) Get(id *uuid.UUID) (*dbentities.Topic, error) {
	topic := &dbentities.Topic{}
	err := r.db.One("Id", id, topic)
	switch err {
	case nil:
		return topic, nil
	case storm.ErrNotFound:
		return nil, model.ErrNotFound
	default:
		return nil, err
	}
}

func (r *TopicsStorm) Update(topic *model.Topic) error {
	err := r.db.Update(dbentities.NewTopic(topic))
	switch err {
	case nil:
		return nil
	case storm.ErrNotFound:
		return model.ErrNotFound
	default:
		return err
	}
}

func (r *TopicsStorm) Delete(topic *model.Topic) error {
	err := r.db.DeleteStruct(dbentities.NewTopic(topic))
	switch err {
	case nil:
		return nil
	case storm.ErrNotFound:
		return model.ErrNotFound
	default:
		return err
	}
}
