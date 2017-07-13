package altsort


func doPivot(data SortInterface, lo int, hi int) int {
	pivot, lo_idx, hi_idx := lo, lo+1, hi

	for {
		for lo_idx < hi && data.Less(lo_idx, pivot) { 
			lo_idx++
		}
		for hi_idx > lo && data.Less(pivot, hi_idx) {
			hi_idx--
		}
		if lo_idx >= hi_idx {
			break
		}
		data.Swap(lo_idx, hi_idx)
	}

	data.Swap(lo, hi_idx)
	return hi_idx
}


func quickSort(data SortInterface, lo int, hi int) {
	if hi <= lo { 
		return
	}
	pivot := doPivot(data, lo, hi)
	quickSort(data, lo, pivot-1)
	quickSort(data, pivot+1, hi)
}


func QuickSort(data SortInterface) {
	quickSort(data, 0, data.Len() - 1) 
}
