package altsort


func SelectSort(data SortInterface) {
	n := data.Len()

	for i := 0; i < n; i++ {
		min := i
		for j := i+1; j < n; j++ {
			if data.Less(j, min) {
				min = j
			}
		}
		data.Swap(i, min)
	}
}
