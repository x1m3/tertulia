package repos

import (
	"github.com/x1m3/Tertulia/model"
	"github.com/nu7hatch/gouuid"
)

type TopicError struct {
	topic *model.Topic
	err error
}

type ItopicsBasic interface {
	Add(*model.Topic) error
	Get(id *uuid.UUID) (*model.Topic, error)
	Update(*model.Topic) error
}

type ItopicsSearch interface {
	GetByCreatedDateDesc(from int, limit int) []*model.Topic
	GetByUpdatedDateDesc(from int, limit int) []*model.Topic
}

type ItopicsAll interface {
	All() (chan []TopicError)
}


type Itopics interface {
	ItopicsBasic
	ItopicsSearch
}

