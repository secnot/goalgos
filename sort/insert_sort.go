package altsort


func InsertSort(data SortInterface) {
	n := data.Len()

	for i := 0; i < n; i++{
		for j := i; j > 0 && data.Less(j, j-1); j--{
			data.Swap(j, j-1)
		}
	}
}
