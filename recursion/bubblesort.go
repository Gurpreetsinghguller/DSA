package main

import "fmt"

func main() {
	arr := []int{5, 4, 3, 2, 1}
	// for i := 0; i < len(arr); i++ {
	// 	for j := 0; j < len(arr)-1; j++ {
	// 		if arr[j] > arr[j+1] {
	// 			temp := arr[j]
	// 			arr[j] = arr[j+1]
	// 			arr[j+1] = temp
	// 		}
	// 	}
	// }

	// fmt.Println(arr)
	fmt.Println(bubblerecurr(arr, 0, 0))
}

func bubblerecurr(arr []int, i, j int) []int {
	if i == len(arr)-1 {
		return arr
	}
	if j < len(arr)-1 && arr[j] > arr[j+1] {
		temp := arr[j]
		arr[j] = arr[j+1]
		arr[j+1] = temp
		return bubblerecurr(arr, i, j+1)
	}
	return bubblerecurr(arr, i+1, 0)
}
