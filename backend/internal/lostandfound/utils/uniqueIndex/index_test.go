package uniqueIndex

import (
	"github.com/google/btree"
	"math"
	"math/rand"
	"strconv"
	"testing"
)

type testItem struct {
	key   int
	value string
}

func NewTestItem(i int) *testItem {
	return &testItem{key: i, value: strconv.Itoa(i)}
}

func (a testItem) Less(b btree.Item) bool {
	return a.key < b.(testItem).key
}

func TestNew(t *testing.T) {
	index := New(2)
	if index.tree.Len() != 0 {
		t.Error("Error creating empty btree")
	}
}

func TestAddRemove(t *testing.T) {
	index := New(2)

	// Inserting some fields
	for i := 0; i < 10000; i++ {
		o := NewTestItem(i)
		j := index.ReplaceOrInsert(*o)
		if j != nil {
			t.Error("Returned value is not the same we inserted before")
		}
		if index.Len() != i+1 {
			t.Error("Error creating empty btree")
		}
		if index.Max().(testItem).key != i {
			t.Error("Error getting max value")
		}
	}

	// Are all data stored?
	for i := 0; i < 10000; i++ {
		o := NewTestItem(i)
		if !index.Has(*o) {
			t.Errorf("Value not found <%v>", o)
		}
		o2 := index.Get(*o).(testItem)
		if o.key != o2.key || o.value != o2.value {
			t.Errorf("Wrong value reading from index. Expecting <%v>, Got <%v>", o, o2)
		}
	}

	// get ordered list DESC is working?
	all := index.AllDesc()
	i := 10000 - 1
	for _, item := range all {
		o := item.(testItem)
		if o.key != i || o.value != strconv.Itoa(i) {
			t.Errorf("index.GetAll() is not working. Expecting <%d> item, got <%v>", i, o)
		}
		i--
	}

	//Removing some fields
	for i := 0; i < 10000; i++ {
		o := NewTestItem(i)
		j := index.Delete(*o)
		if j == nil {
			t.Error("Returned value is not the same we wanted to delete")
		}
		if index.Len() != 10000-i-1 {
			t.Error("Error creating empty btree")
		}
		if i < 10000-1 {
			if index.Min().(testItem).key != i+1 {
				t.Errorf("Error getting min value, expecting <%d>, Got <%v>", i, index.Min().(testItem))
			}
		}
	}
}

func TestIndex_AllAsc(t *testing.T) {
	index := New(2)

	// Inserting some fields
	for i := 0; i < 100000; i++ {
		o := NewTestItem(rand.Int())
		_ = index.ReplaceOrInsert(*o)
	}

	// get ordered list ASC is working?
	all := index.AllAsc()
	previous := math.MinInt32
	for _, item := range all {
		o := item.(testItem)
		if o.key < previous {
			t.Errorf("index.AllAsc() is not working. got items in wrong order")
		}
		previous = o.key
	}
}

func TestIndex_AllDesc(t *testing.T) {
	index := New(2)
	// Inserting some fields
	for i := 0; i < 100000; i++ {
		o := NewTestItem(rand.Intn(10000))
		_ = index.ReplaceOrInsert(*o)
	}

	// get ordered list DESC is working?
	all := index.AllDesc()
	previous := math.MaxInt32
	for _, item := range all {
		o := item.(testItem)
		if o.key > previous {
			t.Errorf("index.AllDesc() is not working. got items in wrong order <%d> should be greater than <%d>", o.key, previous)
		}
		previous = o.key
	}
}

func TestIdxItemList_From(t *testing.T) {
	index := New(2)
	// Inserting some fields
	for i := 0; i < 10000; i++ {
		o := NewTestItem(i)
		_ = index.ReplaceOrInsert(*o)
	}

	half := index.AllAsc().From(5000)
	if len(half) != 5000 {
		t.Errorf("Error in from. Expecting a len of <%d>, got <%d>", 5000, len(half))
	}

	half = index.AllAsc().From(9975)
	if len(half) != 25 {
		t.Errorf("Error in from. Expecting a len of <%d>, got <%d>", 25, len(half))
	}

	half = index.AllAsc().From(10001)
	if len(half) != 0 {
		t.Errorf("Error in from. Expecting a len of <%d>, got <%d>", 0, len(half))
	}
}

func TestIdxItemList_Limit(t *testing.T) {
	index := New(2)
	// Inserting some fields
	for i := 0; i < 100; i++ {
		o := NewTestItem(i)
		_ = index.ReplaceOrInsert(*o)
	}

	half := index.AllAsc().Limit(50)
	if len(half) != 50 {
		t.Errorf("Error in limit. Expecting a len of <%d>, got <%d>", 50, len(half))
	}

	// Limit must copy the slice and return a new one only with the minimum capacity.
	// This could help us saving memory if the underliying array is dereferenced
	if cap(half) != 50 {
		t.Errorf("Error in limit capacity. Expecting a cap of <%d>, got <%d>", 50, cap(half))
	}
}
