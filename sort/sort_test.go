package altsort


import (
	"math"
	"sort"
	"testing"
)


// Test arrays
var ints = [...]int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}
var float64s = [...]float64{74.3, 59.0, math.Inf(1), 238.2, -784.0, 2.3, math.NaN(), math.NaN(), math.Inf(-1), 9845.768, -959.7485, 905, 7.8, 7.8}
var strings = [...]string{"", "Hello", "foo", "bar", "foo", "f00", "%*&^*&^&", "***"}


type SortFunc func(data SortInterface)




func tSortBase(t *testing.T, sortf SortFunc) {
	
	// Copy test arrays
	iunsorted1 := ints
	iunsorted2 := ints
	funsorted1 := float64s
	funsorted2 := float64s
	sunsorted1 := strings
	sunsorted2 := strings

	sortf(IntSlice(iunsorted1[0:]))
	sort.Sort(IntSlice(iunsorted2[0:]))
	if iunsorted1 != iunsorted2 {
		t.Error("Ints didn't sort")
		t.Error("\t", iunsorted1)
		t.Error("\t", iunsorted2)
	}	
	
	sortf(Float64Slice(funsorted1[0:]))
	sort.Sort(Float64Slice(funsorted2[0:]))
	if !IsSorted(Float64Slice(funsorted1[0:])) {
		t.Error("Float64s didn't sort")
		t.Error("\t", funsorted1)
		t.Error("\t", funsorted2)
	}
	
	sortf(StringSlice(sunsorted1[0:]))
	sort.Sort(StringSlice(sunsorted2[0:]))
	if sunsorted1 != sunsorted2 {
		t.Error("Strings didn't sort")
		t.Error("\t", sunsorted1)
		t.Error("\t", sunsorted2)
	}
}


func tSortFunc(t *testing.T, sortf SortFunc) {
	tSortBase(t, sortf)
	//TODO: tSortEmpty(t, sortf)
	//TODO: tSortRandInt(t, sortf)
	//TODO: tSortRandFloat64(t, sortf)
	//TODO: tSortRandString(t, sortf)
}
