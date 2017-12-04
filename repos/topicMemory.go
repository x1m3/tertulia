package repos

import (
	"github.com/google/btree"
	"github.com/x1m3/Tertulia/model"
	"github.com/nu7hatch/gouuid"
	"time"
	"sync"
)

type TopicsMemory struct {
	sync.RWMutex
	indexCreatedDate *btree.BTree
	indexUpdatedDate *btree.BTree
	memory           map[string]*model.Topic
}

func NewTopicsMemory () *TopicsMemory{
	return &TopicsMemory{
		indexUpdatedDate: btree.New(4),
		indexCreatedDate: btree.New(4),
		memory : make(map[string]*model.Topic),
	}
}

func (r *TopicsMemory) Add(topic *model.Topic) error {
	r.Lock()
	defer r.Unlock()

	if _, found := r.memory[topic.ID().String()]; found {
		return ErrDuplicatedValue
	}

	err := r.addToIndexes(topic)
	if err == ErrDuplicatedIndexValue {
		return err
	}

	r.memory[topic.ID().String()] = topic
	return nil
}

func (r *TopicsMemory) Get(id *uuid.UUID) (*model.Topic, error) {
	var t *model.Topic
	var found bool

	r.RLock()
	defer r.RUnlock()

	if t, found = r.memory[id.String()]; !found {
		return nil, ErrNotFound
	}
	return t, nil

}

func (r *TopicsMemory) Update(t *model.Topic) error {

	r.Lock()
	defer r.Unlock()

	if _, found := r.memory[t.ID().String()]; !found {
		return ErrNotFound
	}
	r.memory[t.ID().String()] = t
	return nil
}

func (r *TopicsMemory) GetByCreatedDateDesc(from int, limit int) []*model.Topic {
	return r.getByIndexDescend(r.indexCreatedDate, from, limit)
}

func (r *TopicsMemory) GetByUpdatedDateDesc(from int, limit int) []*model.Topic {
	return r.getByIndexDescend(r.indexUpdatedDate, from, limit)
}

type dateIdxItem struct {
	time  time.Time
	topic *model.Topic
}

func (t dateIdxItem) Less(than btree.Item) bool {
	return t.time.Before(than.(dateIdxItem).time)
}

func (r *TopicsMemory) addToIndexes(t *model.Topic) error {
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

func (r *TopicsMemory) updateIndexes(t *model.Topic) error {
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

func (r *TopicsMemory) getByIndexDescend(index *btree.BTree, from int, limit int) []*model.Topic {
	var all []*model.Topic
	var c, stored int

	if limit == 0 {
		all = make([]*model.Topic, 0)
	} else {
		all = make([]*model.Topic, 0, limit)
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
