package llist_test


import (
	"fmt"
	"github.com/secnot/goalgos/llist"
)


func ExampleQueue() {
	// Create an empty list
	list := Llist.NewLlist()

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
