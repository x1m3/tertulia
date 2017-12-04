package repos

import (
	"github.com/x1m3/Tertulia/model"
	"github.com/nu7hatch/gouuid"
)


type Itopics interface {
	Add(*model.Topic) error
	Get(id *uuid.UUID) (*model.Topic, error)
	Update(*model.Topic) error
	GetByCreatedDateDesc(from int, limit int) []*model.Topic
	GetByUpdatedDateDesc(from int, limit int) []*model.Topic
}

