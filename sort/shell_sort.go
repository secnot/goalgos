package altsort

func ShellSort(data SortInterface) {
	n := data.Len()
	
	// Find starting subsequence distance
	var h int
	for h = 1; h < n/3; h=3*h+1 {}

	// Sort
	for ;h>=1; h=h/3 {
		for i := h; i < n; i++ {
			for j := i; j >= h && data.Less(j, j-h); j -= h {
				data.Swap(j, j-h)
			}
		}
	}
}
