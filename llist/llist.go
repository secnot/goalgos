// A doubly linked list for easy implementation of LIFO/FIFO, stacks, and queues.
package llist

import "fmt"

// Element in the linked list
type element struct {
	Value interface{}

	next *element
	prev *element

}

// Linked list
type Llist struct {
	// Sentinel node
	root element

	// element in the linked list
	length int
}


type Iterator struct {
	// element currently being iterated over
	current *element

	// 
	list *Llist
}


// insert new element after another one
func (l *Llist) insert(elm, newElm *element) {
	newElm.next = elm.next
	newElm.prev = elm
	elm.next.prev = newElm
	elm.next = newElm
	l.length++
}

// remove element and return its value
func (l *Llist) remove(elm *element) interface{} {
	elm.prev.next = elm.next
	elm.next.prev = elm.prev
	l.length--
	return elm.Value
}



func NewLlist() *Llist{
	list := &Llist{}
	list.root.next = &list.root
	list.root.prev = &list.root
	return list
}


func (l *Llist) Len() int{
	return l.length
}


// firstElement returns the first element in the list
func (l *Llist) firstElement() *element {
	if l.root.next != &l.root {
		return l.root.next
	}
	return nil
}

// lastelement returns the last element in the list
func (l *Llist) lastElement() *element {
	if l.root.prev != &l.root {
		return l.root.prev
	}
	return nil
}

// First returns the value of the first element in the list
func (l *Llist) First() (value interface{}, ok bool) {
	first := l.firstElement()
	if first != nil {
		return first.Value, true
	} 
	return nil, false
}

// Last returns the value of the last element in the list
func (l *Llist) Last() (value interface{}, ok bool) {
	last := l.lastElement()
	if last != nil {
		return last.Value, true
	}
	return nil, false
}

// PopLast remove the last element in the list and return its value
func (l *Llist) PopLast() (value interface{}, ok bool) {
	last := l.lastElement()
	if last != nil {
		return l.remove(last), true
	}
	return nil, false
}

// PopFirst remove the first element in the list and return its value
func (l *Llist) PopFirst() (value interface{}, ok bool) {
	first := l.firstElement()
	if first != nil {
		return l.remove(first), true
	}
	return nil, false
}

// Add a value at the end of the list
func (l *Llist) Append(value interface{}) {
	l.insert(l.root.prev, &element{Value: value})
}

// Add a value at the beginnig of the list
func (l *Llist) Prepend(value interface{}) {
	l.insert(&l.root, &element{Value: value})
}

// String interface
func (l *Llist) String() string {
	buffer := make([]string, l.Len())
	iter := l.Iter()
	
	current := 0
	for v, ok := iter.Next(); ok; v, ok = iter.Next() {
		buffer[current] = fmt.Sprintf("%v,", v)
		current++
	}

	return fmt.Sprintf("Llist%v", buffer)
}

// Return iterator to iterate over Linked List
func (l *Llist) Iter() *Iterator {
	iter := &Iterator{list:l, current:&l.root}
	return iter
}

// Next iteration value
func (i *Iterator) Next() (value interface{}, ok bool) {
	if i.current == nil {
		return nil, false
	}

	i.current = i.current.next
	if i.current == &i.list.root { // Finished iteration
		i.current = nil
		return nil, false
	}
	return i.current.Value, true
}

// Delete the element currently being iterated 
func (i *Iterator) Delete() {
	if i.current == nil {
		return
	} else {
		i.list.remove(i.current)
	}
}

// InsertAfter adds a value after the element being iterated
func (i *Iterator) InsertAfter(value interface{}) {
	if i.current == nil {
		i.list.Append(value)
	} else {
		elm := &element{Value: value}
		i.list.insert(i.current.next.prev, elm)
		i.current.next = elm // In the case the current was deleted
	}
}


// InsertBefore adds a value before the element being iterated
func (i *Iterator) InsertBefore(value interface{}) {
	if i.current == nil {
		i.list.Append(value)
	} else {
		elm := &element{Value: value}
		i.list.insert(i.current.prev, elm)
		i.current.prev = elm // In the case the current was deleted
	}
}


// Set current element value while iterating
func (i *Iterator) Set(value interface{}) {
	if i.current == nil && i.current == &i.list.root{
		return
	}
	i.current.Value = value
}
