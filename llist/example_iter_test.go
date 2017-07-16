package llist_test


import (
	"fmt"
	"github.com/secnot/goalgos/llist"
)


func ExampleIter() {

	list := Llist.NewLlist()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Append(4)
	list.Append(5)

	// Delete odd values, add 1000 to the rest
	iter := list.Iter()
	for value, ok := iter.Next(); ok; value, ok = iter.Next() {
		if value.(int) % 2 == 1 {
			iter.Delete()
		} else {
			iter.Set(1000+value.(int))
		}
	}
	fmt.Println(list) 
	
	// Output:
	// Llist[1002, 1004]
}
	
