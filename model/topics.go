package model

import (
	"github.com/google/btree"
	"github.com/nu7hatch/gouuid"
	"sync"
	"time"
)

type Topics struct {
	sync.RWMutex
	indexCreatedDate *btree.BTree
	indexUpdatedDate *btree.BTree
	memory           map[uuid.UUID]*Topic
}

func NewTopics() *Topics {
	return &Topics{
		indexUpdatedDate: btree.New(4),
		indexCreatedDate: btree.New(4),
		memory:           make(map[uuid.UUID]*Topic),
	}
}

func (r *Topics) Add(topic *Topic) error {
	r.Lock()
	defer r.Unlock()

	if _, found := r.memory[*topic.ID()]; found {
		return ErrDuplicatedValue
	}

	err := r.addToIndexes(topic)
	if err == ErrDuplicatedIndexValue {
		return err
	}

	r.memory[*topic.ID()] = topic
	return nil
}

func (r *Topics) Get(id *uuid.UUID) (*Topic, error) {
	var t *Topic
	var found bool

	r.RLock()
	defer r.RUnlock()

	if t, found = r.memory[*id]; !found {
		return nil, ErrNotFound
	}
	return t, nil

}

func (r *Topics) Update(t *Topic) error {

	r.Lock()
	defer r.Unlock()

	if _, found := r.memory[*t.ID()]; !found {
		return ErrNotFound
	}

	err := r.updateIndexes(t)
	if err != nil {
		return err
	}

	r.memory[*t.ID()] = t
	return nil
}

func (r *Topics) GetByCreatedDateDesc(from int, limit int) []*Topic {
	return r.getByIndexDescend(r.indexCreatedDate, from, limit)
}

func (r *Topics) GetByUpdatedDateDesc(from int, limit int) []*Topic {
	return r.getByIndexDescend(r.indexUpdatedDate, from, limit)
}

type dateIdxItem struct {
	time  time.Time
	topic *Topic
}

func (t dateIdxItem) Less(than btree.Item) bool {
	return t.time.Before(than.(dateIdxItem).time)
}

func (r *Topics) addToIndexes(t *Topic) error {
	var p btree.Item
	p = r.indexCreatedDate.ReplaceOrInsert(dateIdxItem{time: t.CreationDate(), topic: t})
	if p != nil { // Repeated value.. Let's abort
		r.indexCreatedDate.ReplaceOrInsert(p)
		return ErrDuplicatedIndexValue
	}

	p = r.indexUpdatedDate.ReplaceOrInsert(dateIdxItem{time: t.ModDate(), topic: t})
	if p != nil { // Repeated value.. Let's abort
		r.indexUpdatedDate.ReplaceOrInsert(p)
		return ErrDuplicatedIndexValue
	}
	return nil
}

func (r *Topics) updateIndexes(t *Topic) error {
	var p btree.Item
	var creationDateIdx, updatedDateIdx dateIdxItem

	creationDateIdx = dateIdxItem{time: t.CreationDate(), topic: t}
	updatedDateIdx = dateIdxItem{time: t.ModDate(), topic: t}

	p = r.indexCreatedDate.Get(creationDateIdx)
	if p == nil {
		return ErrIndexNotFound
	}

	p = r.indexUpdatedDate.Get(updatedDateIdx)
	if p == nil {
		return ErrIndexNotFound
	}
	r.indexCreatedDate.ReplaceOrInsert(creationDateIdx)
	r.indexUpdatedDate.ReplaceOrInsert(updatedDateIdx)

	return nil
}

func (r *Topics) getByIndexDescend(index *btree.BTree, from int, limit int) []*Topic {
	var all []*Topic
	var c, stored int

	if limit == 0 {
		all = make([]*Topic, 0)
	} else {
		all = make([]*Topic, 0, limit)
	}
	c = 0
	stored = 0

	index.Descend(func(a btree.Item) bool {
		if c < from { // Skip unneeded values
			c++
			return true
		}

		all = append(all, a.(dateIdxItem).topic)
		stored++
		if limit != 0 && stored == limit {
			return false
		}
		return true
	})
	return all
}
