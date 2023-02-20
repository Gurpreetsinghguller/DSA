package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	key := 0
	// fmt.Println(binarySearch(arr, key))
	fmt.Println(binarySearchRecurr(arr, key, 0, len(arr)-1))
}

// func binarySearch(arr []int, key int) int {
// 	start := 0
// 	end := len(arr) - 1
// 	for {
// 		mid := (start + end) / 2
// 		if arr[mid] == key {
// 			return mid
// 		}
// 		if arr[mid] > key {
// 			end = mid - 1
// 		}
// 		if arr[mid] < key {
// 			start = mid + 1
// 		}
// 	}
// }

func binarySearchRecurr(arr []int, key, start, end int) int {
	if start > end {
		return -1
	}
	mid := (start + end) / 2
	if arr[mid] == key {
		return mid
	}
	if arr[mid] > key {
		return binarySearchRecurr(arr, key, start, mid-1)
	}
	return binarySearchRecurr(arr, key, mid+1, end)
}
