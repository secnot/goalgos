# Linked List [![Build Status](https://travis-ci.org/secnot/goalgos/llist.svg?branch=master)](https://travis-ci.org/secnot/goalgos/llist) [![GoDoc](https://godoc.org/github.com/secnot/goalgos/llist?status.svg)](http://godoc.org/github.com/secnot/goalgos/llist)

A doubly linked list written in Go that makes easier to implement LIFO/FIFO, 
stacks, and queues than the builtin package "container/list", while still
allowing to loop over the list.


## Usage

Basic operations:

```go
package main

import (
	"fmt"
	"github.com/secnot/goalgos/llist"
)

func main() {

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
}
```

And to iterate over the linked list to delete, modify, and insert elements:


```go
package main

import (
	"fmt"
	"github.com/secnot/goalgos/llist"
)

func main() {

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
	fmt.Println(list) //> Llist[1002, 1004]
}
```
