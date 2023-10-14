package skiplist

import (
	"math"
	"math/rand"
	"strconv"
	"time"
)

type SkipList struct {
	head             *Node
	tail             *Node
	heightOfSkipList int
}

const (
	NEG_INFINITY int32 = math.MinInt32
	POS_INFINITY int32 = math.MaxInt32
	HIGH
)

func NewSkipList() *SkipList {
	rand.Seed(time.Now().UnixNano())
	head := NewNode(NEG_INFINITY)

	tail := NewNode(POS_INFINITY)
	tail.prev = head
	head.next = tail

	return &SkipList{
		head,
		tail,
		0,
	}
}

func (self *SkipList) SkipSearch(key int32) *Node {
	result := self.head
	for result.below != nil { // go down
		result = result.below
		for key >= result.next.key { // go rigth
			result = result.next
		}
	}
	return result
}

func (self *SkipList) RemoveRefsToNode(node *Node) {
	after := node.next
	remove := node.prev

	remove.next = after
	after.prev = remove
}

func (self *SkipList) Remove(key int32) *Node {
	nodeToBeRemoved := self.SkipSearch(key)

	if nodeToBeRemoved.key != key { // not found
		return nil
	}
	self.RemoveRefsToNode(nodeToBeRemoved)

	for nodeToBeRemoved != nil {
		self.RemoveRefsToNode(nodeToBeRemoved)
		if nodeToBeRemoved.above != nil {
			nodeToBeRemoved = nodeToBeRemoved.above
		} else {
			break
		}
	}

	return nodeToBeRemoved
}

func (self *SkipList) Insert(key int32) *Node {
	position := self.SkipSearch(key)

	for position.key == key { // already exists
		return position
	}

	numOfHeads := -1
	level := -1
	var q *Node = nil
	for {

		numOfHeads++
		level++

		self.canIncreaseLevel(level)

		q = position // greatest key <= key

		for position.above == nil {
			position = position.prev
		}
		position = position.above

		q = self.insertAfterAbove(position, q, key)

		if rand.Intn(2) == 0 { // 50% chance
			break
		}
	}
	return q
}

func (self *SkipList) setBeforeAndAfterRef(q *Node, newnode *Node) { //  q <-> newnode <-> q.next
	newnode.next = q.next
	newnode.prev = q
	q.next.prev = newnode
	q.next = newnode
}

func (self *SkipList) setAboveAndBelowRef(position *Node, key int32, newnode, nodeBefore *Node) {
	if nodeBefore != nil {
		node := nodeBefore
		for {
			if node.next.key != key {
				node = nodeBefore.next
			} else {
				break
			}
		}
		newnode.below = nodeBefore.next
		nodeBefore.next.above = newnode
	}

	if position != nil {
		if position.next.key == key {
			newnode.above = position.next
		}
	}
}

func (self *SkipList) insertAfterAbove(position, q *Node, key int32) *Node {
	newnode := NewNode(key)
	nodeBefore := position.below.below
	self.setBeforeAndAfterRef(q, newnode)
	self.setAboveAndBelowRef(position, key, newnode, nodeBefore)
	return newnode
}

func (self *SkipList) canIncreaseLevel(level int) {
	if level >= self.heightOfSkipList {
		self.heightOfSkipList++
		self.addEmptyLevel()
	}
}

func (self *SkipList) addEmptyLevel() {

	newHead := NewNode(NEG_INFINITY)
	newTail := NewNode(POS_INFINITY)

	newHead.next = newTail
	newHead.below = self.head
	newTail.prev = newHead
	newTail.below = self.tail

	self.head.above = newHead
	self.tail.above = newTail

	//finally
	self.head = newHead
	self.tail = newTail
}

func (self *SkipList) Print() {
	result := ""
	starting := self.head
	highestLevel := starting

	level := self.heightOfSkipList

	for highestLevel != nil {
		result += "Level " + strconv.FormatInt(int64(level), 10) + "\n"
		for starting != nil {
			if starting.key == NEG_INFINITY {
				// result += " -∞ "
			} else if starting.key == POS_INFINITY {
				// result += " +∞ "
			} else {
				result += strconv.FormatInt(int64(starting.key), 10)
			}

			if starting.next != nil {
				result += " : "
			}
			starting = starting.next
		}
		highestLevel = highestLevel.below
		starting = highestLevel
		level--
	}
	println(result)
}
