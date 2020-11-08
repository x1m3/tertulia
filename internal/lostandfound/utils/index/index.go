package index

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/google/btree"
	"github.com/x1m3/Tertulia/internal/lostandfound/utils/uniqueIndex"
	"sync"
)

type IdxItem struct {
	index  uniqueIndex.IdxItem
	values IdxItemList
}

func (i IdxItem) Less(than btree.Item) bool {
	fmt.Println("En el primer less")
	spew.Dump(than)
	return i.index.Less(than)
}

type Index struct {
	sync.RWMutex
	index *uniqueIndex.Index
}

type IdxItemList []uniqueIndex.IdxItem

func New(degree int) *Index {
	return &Index{index: uniqueIndex.New(degree)}
}

func (i *Index) Insert(item uniqueIndex.IdxItem) *IdxItem {
	i.Lock()
	defer i.Unlock()

	fmt.Println("lala1")
	previous := i.index.Get(item)
	fmt.Println("lala2")
	if previous != nil {
		fmt.Println("lala3")
		newItem := previous.(IdxItem)
		newItem.values = append(newItem.values, item)
		i.index.ReplaceOrInsert(newItem)

	} else {
		fmt.Println("lala4")
		newItem := IdxItem{index: item}
		newItem.values = append(newItem.values, item)
		i.index.ReplaceOrInsert(newItem)
	}
	return nil

}

func (i *Index) Get(key IdxItem) IdxItemList {
	return i.index.Get(key).(IdxItem).values
}
