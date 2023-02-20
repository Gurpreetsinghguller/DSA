package main

import "fmt"

// func firstOccuranceIndex(arr []int, key, i int) int {
// 	if arr[i] == key {
// 		return i
// 	}
// 	if i == len(arr)-1 {
// 		return -1
// 	}
// 	return firstOccurance(arr, key, i+1)
// }

func firstOccurance(arr []int, key int) int {
	if len(arr) == 1 && arr[0] != key {
		return -1
	}
	if arr[0] == key {
		return key
	}
	return firstOccurance(arr[1:], key)
}

func iiRstOccurance(arr []int, key int) int {
	if len(arr) == 1 && arr[0] != key {
		return -1
	}
	if arr[0] == key {
		return arr[0]
	}
	return iiRstOccurance(arr[1:], key)
}

func iiRstOccuranceIndex(arr []int, key, index int) int {
	if len(arr) == 1 && arr[0] != key {
		return -1
	}
	if arr[0] == key {
		return index
	}
	return iiRstOccuranceIndex(arr[1:], key, index+1)
}

func firstOccuranceIndexAgain(arr []int, key, len int) int {
	if len == 0 {
		return -1
	}
	if arr[0] == key {
		return 0
	}
	i := firstOccuranceIndexAgain(arr[1:], key, len-1)
	if i == -1 {
		return -1
	}
	return i + 1
}

// the data will going to be return, not the index
func main() {
	arr := []int{3, 4, 3, 6, 2}
	key := 2
	// fmt.Println(firstOccurance(arr, key))
	// fmt.Println(firstOccuranceIndexAgain(arr, key, len(arr)))
	fmt.Println(ftoc(arr, key))
}

func ftoc(arr []int, key int) int {
	// 3, 4, 3, 6, 2
	if len(arr) == 0 {
		return -1
	}
	if arr[0] == key {
		return 0
	}
	i := ftoc(arr[1:], key)
	if i == -1 {
		return -1
	}
	return i + 1
}
