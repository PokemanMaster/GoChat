package main

import (
	"fmt"
)

// quickSort sorts an array using the quicksort algorithm
func quickSort(arr []int, low, high int) {
	if low < high {
		// Partition the array and get the pivot index
		p := partition(arr, low, high)
		// Recursively sort the elements before and after partition
		quickSort(arr, low, p-1)
		quickSort(arr, p+1, high)
	}
}

// partition rearranges the elements in the array such that all elements
// less than the pivot are to its left and all elements greater are to its right
func partition(arr []int, low, high int) int {
	pivot := arr[low] // 选取最左边作为基点
	i := low + 1      // Index of greater element

	for j := low + 1; j <= high; j++ {
		// If current element is smaller than or equal to pivot
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}

	// Swap the pivot element with the element at i-1 position
	// 为什么不是 arr[low], arr[i] = arr[i], arr[low]，因为你只要做一次交换，就证明i-1必须小于或者等于pivot
	arr[low], arr[i-1] = arr[i-1], arr[low]
	return i - 1
}

func main() {
	arr := []int{3, 3, 2, 1, 4}
	n := len(arr)
	quickSort(arr, 0, n-1)
	fmt.Println("Sorted array:", arr)
}
