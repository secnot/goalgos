package altsort


import (
	"math"
	"sort"
	"time"
	"math/rand"
	"testing"
)


// Test arrays
var ints = [...]int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}
var float64s = [...]float64{74.3, 59.0, math.Inf(1), 238.2, -784.0, 2.3, math.NaN(), math.NaN(), math.Inf(-1), 9845.768, -959.7485, 905, 7.8, 7.8}
var strings = [...]string{"", "Hello", "foo", "bar", "foo", "f00", "%*&^*&^&", "***"}


type SortFunc func(data SortInterface)




func testSortBase(t *testing.T, sortf SortFunc) {
	
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

func testSortEmpty(t *testing.T, sortf SortFunc) {
	sortf(IntSlice([]int{}))
	sortf(Float64Slice([]float64{}))
	sortf(StringSlice([]string{}))
}


func testSortRandInt(t *testing.T, sortf SortFunc) {
	
	randomIntSlice := func (length int) []int {
		s := rand.NewSource(time.Now().UnixNano())
		r := rand.New(s)

		slice := []int{}
		for i := 0; i < length; i++ {
			slice = append(slice, r.Intn(1000000))
		}
		return slice
	}

	for l := 1; l < 100; l++ {
		random := randomIntSlice(l)
		sortf(IntSlice(random))
		if !IsSorted(IntSlice(random)) {
			t.Error("Random ints weren't sorted")
			t.Error("\t", random)
		}
	}
}


func testSortRandFloat64(t *testing.T, sortf SortFunc) {

	randomFloat64Slice := func (length int) []float64 {
		s := rand.NewSource(time.Now().UnixNano())
		r := rand.New(s)

		slice := []float64{}
		for i := 0; i < length; i++ {
			slice = append(slice, r.Float64()*1000000)
		}
		return slice
	}

	for l := 1; l < 100; l++ {
		random := randomFloat64Slice(l)
		sortf(Float64Slice(random))
		if !IsSorted(Float64Slice(random)) {
			t.Error("Random floats weren't sorted")
			t.Error("\t", random)
		}
	}
}


func testSortFunc(t *testing.T, sortf SortFunc) {
	testSortBase(t, sortf)
	testSortEmpty(t, sortf)
	for i := 0; i < 10; i++ {
		testSortRandInt(t, sortf)
		testSortRandFloat64(t, sortf)
	}
}
