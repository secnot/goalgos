package llist


import (
	"fmt"
)


func ExampleIter() {

	list := NewLlist()
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
	// Llist[1002, 1004,]
}


func ExampleQueue() {
	// Create an empty list
	list := NewLlist()

	// Append elements at the end and start or the list
	list.Append(1)
	list.Append(4)
	list.Prepend(5)
	fmt.Println(list) // > Llist[5, 1, 4,]

	// Return the values of First and Last elements (peek)
	list.Last() // > 4, true
	list.First() // > 5, true

	// Pop last and first elements returning their value
	list.PopLast()    //> 4, true
	list.PopFirst()   //> 5, true
	fmt.Println(list) //> Llist[1,]

	// Empty the list
	list.PopLast()    //> 1, true
	fmt.Println(list) //> Llist[]

	// Poping from and empty linked list returns nil
	list.PopLast()  //> nil, false
	list.PopFirst() //> nil, false
	list.Last()     //> nil, false
	list.First()    //> nil, false

	// Output:
	// Llist[5, 1, 4,]
	// Llist[1,]
	// Llist[]
}
