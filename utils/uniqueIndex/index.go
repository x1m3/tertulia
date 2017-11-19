package uniqueIndex

import (
	"github.com/google/btree"
	"sync"
)

// A thread safe wrapper for the github.com/google/btree implementation
type Index struct {
	sync.RWMutex
	tree  *btree.BTree
}

type IdxItem btree.Item

type IdxItemList []IdxItem

func (i IdxItemList) From(pos int) IdxItemList {
	if pos>=0 && len(i)>pos {
		return i[pos:]
	}else {
		return nil
	}
}

func (i IdxItemList) Limit(lim int) IdxItemList {
	if lim<=0 {
		return nil
	}
	if len(i)<lim {
		lim = len(i)
	}
	r := make(IdxItemList, lim)
	_ = copy(r,i[:lim])
	return r
}

func New(degree int) *Index {
	return &Index{tree:btree.New(degree)}
}

func (i *Index) ReplaceOrInsert(item IdxItem) IdxItem {
	i.Lock()
	defer i.Unlock()

	return i.tree.ReplaceOrInsert(item)
}

func (i *Index) Get(key IdxItem) IdxItem {
	i.RLock()
	defer i.RUnlock()
	return i.tree.Get(key)
}

func (i *Index) Has(key IdxItem) bool {
	i.RLock()
	defer i.RUnlock()
	return i.tree.Has(key)
}

func (i *Index) Delete(item IdxItem) IdxItem {
	i.Lock()
	defer i.Unlock()
	return i.tree.Delete(item)
}

func (i *Index) Len() int {
	i.RLock()
	defer i.RUnlock()
	return i.tree.Len()
}

func (i *Index) Max() IdxItem {
	i.RLock()
	defer i.RUnlock()
	return i.tree.Max()
}

func (i *Index) Min() IdxItem {
	i.RLock()
	defer i.RUnlock()
	return i.tree.Min()
}

func (i *Index) AllAsc() IdxItemList {
	i.RLock()
	defer i.RUnlock()
	all := make([]IdxItem,0, i.tree.Len())
	i.tree.Ascend(func(a btree.Item) bool {
		all = append(all, a)
		return true
	})
	return all
}

func (i *Index) AllDesc() IdxItemList {
	i.RLock()
	defer i.RUnlock()
	all := make([]IdxItem,0, i.tree.Len())
	i.tree.Descend(func(a btree.Item) bool {
		all = append(all, a)
		return true
	})
	return all
}