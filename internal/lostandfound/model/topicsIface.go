package model

import (
	"github.com/nu7hatch/gouuid"
)

type TopicError struct {
	Topic *Topic
	Err   error
}

type ItopicsCRUD interface {
	Add(*Topic) error
	Get(id *uuid.UUID) (*Topic, error)
	Update(*Topic) error
	Delete(*Topic) error
}

type ItopicsAll interface {
	All(chan TopicError)
}
