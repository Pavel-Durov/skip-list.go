package list

type SortedList struct {
	head   *Node
	length int32
}

func NewList() *SortedList {
	return &SortedList{
		head:   nil,
		length: 0,
	}
}

func (self *SortedList) GetHead() *Node {
	return self.head
}

func (self *SortedList) GetLength() int32 {
	return self.length
}

func (self *SortedList) ToList() []int32 {
	node := self.head
	list := []int32{}
	for {
		if node == nil {
			break
		}
		list = append(list, node.getValue())
		node = node.getNext()
	}
	return list
}

func (self *SortedList) Insert(value int32) {

	if self.head == nil {
		self.head = &Node{
			value: value,
			next:  nil,
		}
		self.length = 1
	} else {
		current := self.head
		for {
			if current.getNext() == nil {
				current.setNext(&Node{
					value: value,
					next:  nil,
				})
				break
			} else if value >= current.getValue() && value <= current.getNext().getValue() {
				temp := current.getNext()
				current.setNext(&Node{
					value: value,
					next:  temp,
				})
				break
			}
		}

		self.length++
	}
}
