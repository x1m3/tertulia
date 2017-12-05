package repos

import (
	"github.com/nu7hatch/gouuid"
	"github.com/x1m3/Tertulia/model"
	"github.com/x1m3/Tertulia/repos/dbentities"
)

type TopicError struct {
	topic *model.Topic
	err   error
}

type ItopicsCRUD interface {
	Add(*model.Topic) error
	Get(id *uuid.UUID) (*dbentities.Topic, error)
	Update(*model.Topic) error
	Delete(*model.Topic) error
}

type ItopicsAll interface {
	All() chan []TopicError
}
