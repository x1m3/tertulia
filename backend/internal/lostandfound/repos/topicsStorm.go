package repos

import (
	"github.com/asdine/storm"
	"github.com/nu7hatch/gouuid"
	"github.com/x1m3/Tertulia/backend/internal/lostandfound/model"
	"github.com/x1m3/Tertulia/backend/internal/lostandfound/repos/dbentities"
)

type TopicsStorm struct {
	db *storm.DB
}

func NewTopicsStorm(db *storm.DB) *TopicsStorm {
	return &TopicsStorm{db: db}
}

func (r *TopicsStorm) Add(topic *model.Topic) error {
	// TODO: Return a repo generic error
	return r.db.Save(dbentities.NewTopicDTO(topic))
}

func (r *TopicsStorm) Get(id *uuid.UUID) (*dbentities.TopicDTO, error) {
	topic := &dbentities.TopicDTO{}
	err := r.db.One("Id", id, topic)
	switch err {
	case nil:
		return topic, nil
	case storm.ErrNotFound:
		// TODO: Return a repo error, not a model one
		return nil, model.ErrNotFound
	default:
		return nil, err
	}
}

func (r *TopicsStorm) Update(topic *model.Topic) error {
	err := r.db.Update(dbentities.NewTopicDTO(topic))
	switch err {
	case nil:
		return nil
	case storm.ErrNotFound:
		// TODO: Return a repo error, not a model one
		return model.ErrNotFound
	default:
		return err
	}
}

func (r *TopicsStorm) Delete(topic *model.Topic) error {
	err := r.db.DeleteStruct(dbentities.NewTopicDTO(topic))
	switch err {
	case nil:
		return nil
	case storm.ErrNotFound:
		// TODO: Return a repo error, not a model one
		return model.ErrNotFound
	default:
		return err
	}
}

func (r *TopicsStorm) All(responseChan chan model.TopicError) {
	buffersize := 1 + 2*cap(responseChan)
	count := 0

	for {
		topics := make([]dbentities.TopicDTO, 0, buffersize)
		err := r.db.All(&topics, storm.Skip(count), storm.Limit(buffersize))
		if err != nil {
			responseChan <- model.TopicError{Topic: nil, Err: err}
		}
		if len(topics) == 0 {
			close(responseChan)
			return
		}
		count += buffersize
		for _, topic := range topics {
			responseChan <- model.TopicError{Topic: topic.ToTopic(), Err: nil}
		}
	}
}
