// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func partition(arr []int, begin int, end int) int {
	var pivot int = arr[end]
	var i int = begin-1

	for j:= begin; j < end; j++ {
		if arr[j] <= pivot {
			i++

			var swapTemp int = arr[i]
			arr[i] = arr[j]
			arr[j] = swapTemp
		}
	}

	var swapTemp int = arr[i+1]
	arr[i+1] = arr[end]
	arr[end] = swapTemp

	return i+1
}

func quickSort(arr []int, begin int, end int) {
	if begin < end {
		var partitionIndex int = partition(arr, begin, end)
		quickSort(arr, begin, partitionIndex - 1)
		quickSort(arr, partitionIndex + 1, end)
	}
}

func main() {
	data := []int{100, 1, 3, 5, -1, 9, -33}
	fmt.Println(data)
	quickSort(data, 0, len(data)-1)
	fmt.Printf("After Sorting: %v\n", data) // string format example

	largeData := []int{4245, -5432, 112, -500, 23, -219993, -2934294, 400, 0, 0, 0, 110, 114, -43, 134, 9999999}
	fmt.Println(data)
	quickSort(largeData, 0, len(largeData)-1)
	fmt.Printf("After Sorting: %v\n", largeData) // string format example
}