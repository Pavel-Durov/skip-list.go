package list

type Node struct {
	value int32
	next  *Node
}

func (self *Node) getNext() *Node {
	return self.next
}

func (self *Node) setNext(next *Node) {
	self.next = next
}

type List struct {
	head   *Node
	length int32
}

func NewList() *List {
	return &List{
		head:   nil,
		length: 0,
	}
}

func (self *List) GetHead() *Node {
	return self.head
}

func (self *List) GetLength() int32 {
	return self.length
}

func (self *List) Append(value int32) {
	if self.head == nil {
		self.head = &Node{
			value: value,
			next:  nil,
		}
		self.length = 1
	} else {
		lastNode := self.head
		for {
			if lastNode.getNext() == nil {
				break
			}
			lastNode = lastNode.getNext()
		}
		lastNode.setNext(&Node{
			value: value,
			next:  nil,
		})
		self.length++
	}
}
