package index

import (
	"fmt"
	"github.com/google/btree"
)

type testItem struct {
	key   int
	value string
}

func NewTestItem(key int, val string) *testItem {
	return &testItem{key: key, value: val}
}

func (a testItem) Less(b btree.Item) bool {
	fmt.Println("En el less")

	return a.key < b.(testItem).key
}

/*
func TestAddRepeated(t *testing.T) {
	t.Skipped()


	index := New(2)
	index.Insert(NewTestItem(1,"1"))
	index.Insert(NewTestItem(1,"1.1"))
	/*
	index.Insert(NewTestItem(1,"1.1.1"))
	index.Insert(NewTestItem(1,"1.1.1.1"))
	index.Insert(NewTestItem(1,"1.1.1.1.1"))

}
*/
