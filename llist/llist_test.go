package llist

import (
	"testing"
	"fmt"
)


// Compare int slice agains linked list values
func hasValues(t *testing.T, list *Llist, values []int) {
	
	var llistValues []int
	iter := list.Iter()
	for value, ok := iter.Next(); ok; value, ok = iter.Next() {
		llistValues = append(llistValues,  value.(int))
	}


	if len(llistValues) != len(values) {
		t.Error("LikedList and reference values aren't of same lenght", 
			len(llistValues), len(values))
		return
	}
		
	for n, val := range values {
		if llistValues[n] != val {
			t.Error("Reference values and linked list contents differ:")
			t.Error("\t", values)
			t.Error("\t", llistValues)
			break
		}
	}
}

func hasLen(t *testing.T, list *Llist, length int) {
	if list.Len() != length {
		t.Error("Len(): Expected", length, "was", list.Len())
	}

	elements := 0
	iter := list.Iter()
	for _, ok := iter.Next(); ok; _, ok = iter.Next() {
		elements++
	}

	if elements != length {
		t.Error("Expected", length, "elements, but there were", elements)
	}
}

func hasLast(t *testing.T, list *Llist, value int) {
	if v, ok := list.Last(); !ok || v != value {
		t.Error("Last() expected", value, "returned", v)
	}
}

func hasFirst(t *testing.T, list *Llist, value int) {
	if v, ok := list.First(); !ok || v != value {
		t.Error("First() expected", value, "returned", v)
	}
}

func initLlist(list *Llist, size int) {

	for i := 0; i < size; i++ {
		list.Append(i)
	}
}


// Test Basic operations
func TestBase(t *testing.T) {

	list := NewLlist()
	hasLen(t, list, 0)

	// Append() and Prepend()
	list.Append(22)
	list.Append(44)
	list.Prepend(33)
	hasValues(t, list, []int{33, 22, 44})

	// PopLast() and PopFirst()
	if value, ok := list.PopLast(); !ok || value != 44 {
		t.Error("PopLast() didn't return the last element")
	}
	hasValues(t, list, []int{33, 22})

	if value, ok := list.PopFirst(); !ok || value != 33 {
		t.Error("PopFirst() didn't return the first element")
	}
	hasValues(t, list, []int{22})

	if value, ok := list.PopLast(); !ok || value != 22 {
		t.Error("PopLast() didn't return the last element")
	}
	hasValues(t, list, []int{})
	hasLen(t, list, 0)

	// Popping from empty list is not an error
	if value, ok := list.PopLast(); ok || value != nil {
		t.Error("Popping form an empty list should return (nil, false)")
	}
	if value, ok := list.PopFirst(); ok || value != nil {
		t.Error("Popping form an empty list should return (nil, false)")
	}

	// First() and Last()
	list.Prepend(100)
	hasLast(t, list, 100)
	hasFirst(t, list, 100)

	list.Append(101)
	hasLast(t, list, 101)
	hasFirst(t, list, 100)
	
	list.Prepend(102)
	hasLast(t, list, 101)
	hasFirst(t, list, 102)

	hasValues(t, list, []int{102, 100, 101})
}

// Test stringer interface
func TestString(t *testing.T) {
	list := NewLlist()
	initLlist(list, 10)

	str := "Llist[0, 1, 2, 3, 4, 5, 6, 7, 8, 9,]"
	if fmt.Sprintf("%v", list) != str {
		t.Error("Llist to String not working")
	}

	// Empty list
	list = NewLlist()
	str = "Llist[]"
	if fmt.Sprintf("%v", list) != str {
		t.Error("Llist to String not working")
	}
}

// Delete element while iterating
func TestIterDelete(t *testing.T) {

	// Delete odd values
	list := NewLlist()
	initLlist(list, 10)
	iter := list.Iter()
	for v, ok := iter.Next(); ok; v, ok = iter.Next() {
		if v.(int) % 2 == 1 {
			iter.Delete()
		}
	}
	hasValues(t, list, []int{0, 2, 4, 6, 8})

	// Delete even values
	list = NewLlist()
	initLlist(list, 10)	
	
	iter = list.Iter()
	for v, ok := iter.Next(); ok; v, ok = iter.Next() {
		if v.(int) % 2 == 0 {
			iter.Delete()
		}
	}
	hasValues(t, list, []int{1, 3, 5, 7, 9})
	list.Append(22)
	list.Prepend(11)
	hasValues(t, list, []int{11, 1, 3, 5, 7, 9, 22})

	// Delete all values
	list = NewLlist()
	initLlist(list, 10)

	iter = list.Iter()
	for _, ok := iter.Next(); ok; _, ok = iter.Next() {
		iter.Delete()
	}
	hasValues(t, list, []int{})
	list.Append(66)
	list.Prepend(44)
	list.Prepend(33)
	hasValues(t, list, []int{33, 44, 66})

	// Delete even values twice	
	list = NewLlist()
	initLlist(list, 10)

	iter = list.Iter()
	for v, ok := iter.Next(); ok; v, ok = iter.Next() {
		if v.(int) % 2 == 0 {
			iter.Delete()
			iter.Delete()
		}
	}
	hasValues(t, list, []int{1, 3, 5, 7, 9})
	list.Append(66)
	list.Prepend(44)
	list.Prepend(33)
	hasValues(t, list, []int{33, 44, 1, 3, 5, 7, 9, 66})
	
	// Delete all values twice	
	list = NewLlist()
	initLlist(list, 10)

	iter = list.Iter()
	for _, ok := iter.Next(); ok; _, ok = iter.Next() {
		iter.Delete()
		iter.Delete()
	}
	hasValues(t, list, []int{})
	list.Append(66)
	list.Prepend(44)
	list.Prepend(33)
	hasValues(t, list, []int{33, 44, 66})

}


// Test adding values after the current while iterating
func TestIterInsertAfter(t *testing.T) {
	
	// Add one value after each odd one
	list := NewLlist()
	initLlist(list, 6)

	iter := list.Iter()
	for v, ok := iter.Next(); ok; v, ok = iter.Next() {
		if v.(int) % 2 == 1 {
			iter.InsertAfter(v.(int)+101)
		}
	}
	hasValues(t, list, []int{0, 1, 102,  2, 3, 104, 4, 5, 106}) 

	// Add two values after each odd one
	list = NewLlist()
	initLlist(list, 6)

	iter = list.Iter()
	for v, ok := iter.Next(); ok; v, ok = iter.Next() {
		if v.(int) % 2 == 1 {
			iter.InsertAfter(v.(int)+101)
			iter.InsertAfter(v.(int)+201)
		}
	}
	hasValues(t, list, []int{0, 1, 202, 102, 2, 3, 204, 104, 4, 5, 206, 106}) 

	// Add values to finished iter
	iter.InsertAfter(7)
	iter.InsertAfter(8)
	hasValues(t, list, []int{0, 1, 202, 102, 2, 3, 204, 104, 4, 5, 206, 106, 7, 8}) 
}


// Test adding values before the current one while iterating
func TestIterInsertBefore(t *testing.T) {	
	// Add one value after each odd one
	list := NewLlist()
	initLlist(list, 6)

	iter := list.Iter()
	for v, ok := iter.Next(); ok; v, ok = iter.Next() {
		if v.(int) % 2 == 1 {
			iter.InsertBefore(v.(int)+101)
		}
	}
	hasValues(t, list, []int{0, 102, 1, 2, 104, 3, 4, 106, 5}) 

	// Add two values after each odd one
	list = NewLlist()
	initLlist(list, 6)

	iter = list.Iter()
	for v, ok := iter.Next(); ok; v, ok = iter.Next() {
		if v.(int) % 2 == 1 {
			iter.InsertBefore(v.(int)+101)
			iter.InsertBefore(v.(int)+201)
		}
	}
	hasValues(t, list, []int{0, 102, 202, 1, 2, 104, 204, 3, 4, 106, 206, 5}) 

	// Add values to finished iter
	iter.InsertBefore(7)
	iter.InsertBefore(8)
	hasValues(t, list, []int{0, 102, 202, 1, 2, 104, 204, 3, 4, 106, 206, 5, 7, 8}) 
}


// Test inserting and deleting elements while iterating over the linked list
func TestIterMixedInsertDelete(t *testing.T) {
	
	// InsertBefore, Delete, Delete, InsertAfter, Delete, InsertBefore 
	for length := 0; length < 100; length++ {
		list := NewLlist()
		initLlist(list, length)

		result := make([]int, 0)

		iter := list.Iter()
		for v, ok := iter.Next(); ok; v, ok = iter.Next() {
			if v.(int) % 2 == 0 {
				iter.InsertBefore(1011)
				iter.Delete()
				iter.Delete()
				iter.InsertAfter(1043)
				iter.Delete()
				iter.InsertBefore(1015)
				result = append(result, 1011)
				result = append(result, 1015)
				result = append(result, 1043)
			} else if v.(int) < 1000 {
				result = append(result, v.(int))
			}
		}

		hasValues(t, list, result)
	}

	// Same test bur for odd members
	for length := 0; length < 100; length++ {
		list := NewLlist()
		initLlist(list, length)

		result := make([]int, 0)

		iter := list.Iter()
		for v, ok := iter.Next(); ok; v, ok = iter.Next() {
			if v.(int) % 2 == 1 {
				iter.InsertBefore(1012)
				iter.Delete()
				iter.Delete()
				iter.InsertAfter(1044)
				iter.Delete()
				iter.InsertBefore(1016)
				result = append(result, 1012)
				result = append(result, 1016)
				result = append(result, 1044)
			} else if v.(int) < 1000 {
				result = append(result, v.(int))
			}
		}

		hasValues(t, list, result)
	}
}


// Test setting elements value while iterating
func TestIterSet(t *testing.T) {

	// Add 1000 to odd values, Delete even
	for length := 0; length < 100; length++ {
		list := NewLlist()
		initLlist(list, length)
	
		// Iterate and modify linked list
		iter := list.Iter()
		for v, ok := iter.Next(); ok; v, ok = iter.Next() {
			if v.(int) %2 == 1 {
				iter.Set(v.(int)+1000)
			} else {
				iter.Delete()
			}
		}

		// Construct expected result
		result := make([]int, 0)
		for i := 0; i < length; i++ {
			if i % 2 == 1 {
				result = append(result, i+1000)
			}
		}

		// Compare results
		hasValues(t, list, result)
	}
}

// Test appending and prepending to list while iterating
func TestIterMixedDeleteAppend(t *testing.T) {
	
	// Append
	for length := 0; length < 100; length++ { // Llist length
		for pos := 0; pos < length; pos++ { // Position where append is called

			list := NewLlist()
			initLlist(list, length)
			
			iter := list.Iter()
			current := 0
			for _, ok:= iter.Next(); ok; _, ok = iter.Next(){
				if current == pos {
					iter.Delete()
					list.Append(2000)
				}
				current++
			}
			
			// Expected result
			result:= make([]int, 0)
			for i:= 0; i<length; i++ {
				if i != pos {
					result = append(result, i)
				}
			}
			result = append(result, 2000)
			hasValues(t, list, result)
		}
		
	}	

	// Prepend
	for length := 0; length < 100; length++ { // Llist length
		for pos := 0; pos < length; pos++ { // Position where append is called

			list := NewLlist()
			initLlist(list, length)
			
			iter := list.Iter()
			current := 0
			for _, ok := iter.Next(); ok; _, ok = iter.Next() {
				if current == pos {
					iter.Delete()
					list.Prepend(2000)
				}
				current++
			}
			
			// Expected result
			result:= make([]int, 0)
			result = append(result, 2000)
			for i:= 0; i<length; i++ {
				if i != pos {
					result = append(result, i)
				}
			}
			hasValues(t, list, result)
		}
		
	}

}

