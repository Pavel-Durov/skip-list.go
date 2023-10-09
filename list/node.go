package list

type Node struct {
	value int32
	next  *Node
}

func (self *Node) getNext() *Node {
	return self.next
}

func (self *Node) getValue() int32 {
	return self.value
}

func (self *Node) setNext(next *Node) {
	self.next = next
}
