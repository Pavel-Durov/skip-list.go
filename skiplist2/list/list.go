package list

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

const (
	maxLevels   = 16
	probability = 0.5
)

type Node struct {
	key     int
	value   interface{}
	forward []*Node
}

type SkipList struct {
	header *Node
	level  int
}

func NewNode(key int, value interface{}, level int) *Node {
	forward := make([]*Node, level)
	return &Node{
		key:     key,
		value:   value,
		forward: forward,
	}
}

func NewSkipList() *SkipList {
	rand.Seed(time.Now().UnixNano())
	header := NewNode(math.MinInt32, nil, maxLevels)
	return &SkipList{
		header: header,
		level:  1,
	}
}

func randomLevel() int {
	level := 1
	for rand.Float32() < probability && level < maxLevels {
		level++
	}
	return level
}

func (sl *SkipList) Insert(key int, value interface{}) {
	update := make([]*Node, maxLevels)
	x := sl.header
	for i := sl.level - 1; i >= 0; i-- {
		for x.forward[i] != nil && x.forward[i].key < key {
			x = x.forward[i]
		}
		update[i] = x
	}

	level := randomLevel()
	if level > sl.level {
		for i := sl.level; i < level; i++ {
			update[i] = sl.header
		}
		sl.level = level
	}

	newNode := NewNode(key, value, level)
	for i := 0; i < level; i++ {
		newNode.forward[i] = update[i].forward[i]
		update[i].forward[i] = newNode
	}
}

func (sl *SkipList) Search(key int) interface{} {
	x := sl.header
	for i := sl.level - 1; i >= 0; i-- {
		for x.forward[i] != nil && x.forward[i].key < key {
			x = x.forward[i]
		}
	}
	if x.forward[0] != nil && x.forward[0].key == key {
		return x.forward[0].value
	}
	return nil
}

func (sl *SkipList) Delete(key int) {
	update := make([]*Node, maxLevels)
	x := sl.header
	for i := sl.level - 1; i >= 0; i-- {
		for x.forward[i] != nil && x.forward[i].key < key {
			x = x.forward[i]
		}
		update[i] = x
	}

	if x.forward[0] != nil && x.forward[0].key == key {
		for i := 0; i < sl.level; i++ {
			if update[i].forward[i] != x.forward[i] {
				break
			}
			update[i].forward[i] = x.forward[i].forward[i]
		}
		for sl.level > 1 && sl.header.forward[sl.level-1] == nil {
			sl.level--
		}
	}
}

func (sl *SkipList) Print() {
	for i := sl.level - 1; i >= 0; i-- {
		fmt.Printf("Level %d: ", i)
		x := sl.header.forward[i]
		for x != nil {
			fmt.Printf("%d ", x.key)
			x = x.forward[i]
		}
		fmt.Println()
	}
}
