package quicksort

import "fmt"

// https://rgalus.medium.com/sorting-algorithms-quick-sort-implementation-in-go-9ebfd91fe95f

func QuickSort(arr []int, start, end int) {
	fmt.Printf("\narr: %d, start: %d, end: %d\n", arr, start, end)

	if (end - start) < 1 {
		return
	}

	pivot := arr[end]
	pivotIndex := start

	for i := start; i < end; i++ {
		if arr[i] < pivot {

			temp := arr[pivotIndex]

			arr[pivotIndex] = arr[i]
			arr[i] = temp

			pivotIndex++
		}
	}

	arr[end] = arr[pivotIndex]
	arr[pivotIndex] = pivot

	QuickSort(arr, start, pivotIndex-1)
	QuickSort(arr, pivotIndex+1, end)
}
