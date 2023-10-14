package main

import (
	skiplist "p3ld3v.dev/skiplist/skip_list"
)

func main() {
	// fmt.Println("Hello!")
	// node := list.Node{}
	// fmt.Println(node)

	slist := skiplist.NewSkipList()
	slist.Insert(6)
	slist.Insert(15)
	slist.Insert(4)
	slist.Insert(1)
	slist.Insert(2)
	slist.Insert(4)
	slist.Insert(3)
	slist.Insert(7)
	// slist.Insert(22)

	slist.Print()
	// println("@@@@@@@")
	// slist.Remove(4)
	// slist.Print()
}
