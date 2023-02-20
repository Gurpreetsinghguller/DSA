package main

import "fmt"

func LastOccurance(arr []int, key int) int {
	fmt.Println(arr)
	if len(arr) == 0 {
		return -1
	}
	if arr[len(arr)-1] == key {
		return len(arr) - 1
	}
	return LastOccurance(arr[0:len(arr)-1], key)
}

// func LastOccr(arr []int, key, lenarr int) int {
// 	if len(arr) == 0 {
// 		return -1
// 	}
// 	if arr[len(arr)-1] == key {
// 		return
// 	}
// }
func LastOccrItr(arr []int, key, itr, index int) int {
	if itr == len(arr) && index == -1 {
		return -1
	}
	if itr == len(arr) {
		return index
	}
	if arr[itr] == key {
		index = itr
	}
	nextPossibility := LastOccrItr(arr, key, itr+1, index)
	if nextPossibility == -1 {
		return -1
	}
	return nextPossibility
}

func main() {
	arr := []int{2, 3, 4, 2, 3, 6, 5, 2, 9, 8}
	key := 2
	// fmt.Println(LastOccuranceItr(arr, key))
	// (arr []int, key, itr, index int)
	// fmt.Println(LastOccrItr(arr, key, 0, -1))
	fmt.Println(loc(arr, key))
}

func loc(arr []int, key int) int {
	if len(arr) == 0 {
		return -1
	}
	if arr[len(arr)-1] == key {
		return len(arr) - 1
	}
	return loc(arr[:len(arr)-1], key)
}
