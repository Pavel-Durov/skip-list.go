package list_test

import (
	"fmt"
	"testing"

	"p3ld3v.dev/skiplist/list"
)

func TestNode(t *testing.T) {
	fmt.Println("Hello!")
	node := list.NewList()
	if node == nil {
		t.Errorf("Node is nil")
		t.Fail()
	}
}

func TestAppend(t *testing.T) {
	list := list.NewList()
	list.Append(1)
	list.Append(2)
	if list.GetLength() != 2 {
		t.Errorf("Expected legth to be 1")
		t.Fail()
	}
	list.Append(1)
	if list.GetLength() != 3 {
		t.Errorf("Expected legth to be 1")
		t.Fail()
	}
}
