package main

import "fmt"

func sortedOrNot(arr []int) bool {
	fmt.Println(arr)
	if len(arr) <= 1 {
		return true
	}
	if arr[0] > arr[1] {
		return false
	}
	return sortedOrNot(arr[1:])
}

func sortedOrNotItr(arr []int, itr int) bool {
	if itr == len(arr)-1 {
		return true
	}
	if arr[itr] > arr[itr+1] {
		return false
	}
	return sortedOrNotItr(arr, itr+1)
}

func st(arr []int) int {
	if len(arr) == 1 {
		return 1
	}
	if arr[0] > arr[1] {
		return -1
	}
	return st(arr[1:])
}

func main() {
	arr := []int{1, 4, 6, 8, 11, 10}
	fmt.Println(st(arr))
	// fmt.Println("dummy", arr[1:])
	// fmt.Println(sortedOrNot(arr))
	// fmt.Println(sortedOrNotItr(arr, 0))
}
