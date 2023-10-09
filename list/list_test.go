package list_test

import (
	"testing"

	"p3ld3v.dev/skiplist/list"
)

func TestInsert(t *testing.T) {
	list := list.NewList()
	list.Insert(1)
	list.Insert(3)
	if list.GetLength() != 2 {
		t.Errorf("Expected legth to be 1")
		t.Fail()
	}
	list.Insert(2)
	list.Insert(2)
	list.Insert(1)

	if list.GetLength() != 5 {
		t.Errorf("Expected legth to be 1")
		t.Fail()
	}
	arr := list.ToList()
	if arr[0] != 1 {
		t.Errorf("Expected first value to be 1")
		t.Fail()
	}
	if arr[1] != 1 {
		t.Errorf("Expected second value to be 1")
		t.Fail()
	}
	if arr[2] != 2 {
		t.Errorf("Expected third value to be 3")
		t.Fail()
	}
	if arr[3] != 2 {
		t.Errorf("Expected fourth value to be 2")
		t.Fail()
	}
	if arr[4] != 3 {
		t.Errorf("Expected fith value to be 3")
		t.Fail()
	}
}
