package skiplist

type Node struct {
	above *Node
	below *Node
	next  *Node
	prev  *Node
	key   int32
}

func NewNode(key int32) *Node {
	return &Node{
		nil,
		nil,
		nil,
		nil,
		key,
	}
}
