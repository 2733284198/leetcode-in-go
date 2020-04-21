package problems0419

func getLeastNumbers(arr []int, k int) []int {
	recursiveSort(arr, 0, len(arr)-1)

	return arr[:k]
}

func recursiveSort(arr []int, start, end int) {
	if end < start {
		return
	}

	pivot := arr[end]
	splitIndex := start

	// Iterate sub array to find values less than pivot
	//   and move them to the beginning of the array
	//   keeping splitIndex denoting less-value array size
	for i := start; i < end; i++ {
		if arr[i] < pivot {
			if splitIndex != i {
				arr[splitIndex], arr[i] = arr[i], arr[splitIndex]
			}
			splitIndex++
		}
	}

	arr[end], arr[splitIndex] = arr[splitIndex], pivot

	recursiveSort(arr, start, splitIndex-1)
	recursiveSort(arr, splitIndex+1, end)
}
