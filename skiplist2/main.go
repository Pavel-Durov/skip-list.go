package main

import (
	"fmt"

	"p3ld3v.dev/skiplist/skiplist2/list"
)

func main() {
	sl := list.NewSkipList()

	sl.Insert(10, "A")
	sl.Insert(20, "B")
	sl.Insert(5, "C")
	sl.Insert(15, "D")
	sl.Insert(25, "E")
	sl.Insert(50, "Z")

	fmt.Println("Skip List:")
	sl.Print()

	key := 15
	result := sl.Search(key)
	if result != nil {
		fmt.Printf("Key %d found with value: %s\n", key, result)
	} else {
		fmt.Printf("Key %d not found\n", key)
	}

	// keyToDelete := 20
	// sl.Delete(keyToDelete)
	// fmt.Printf("Key %d deleted\n", keyToDelete)

	fmt.Println("Updated Skip List:")
	sl.Print()
}
