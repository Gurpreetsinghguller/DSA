package main

import "fmt"

func printKeyOccurance(arr []int, key int) {
	if len(arr) == 0 {
		return
	}
	if arr[0] == key {
		fmt.Println(key)
	}
	printKeyOccurance(arr[1:], key)
}

func printIndexOccurance(arr []int, key, Itr int) {
	if Itr == len(arr) {
		return
	}
	if arr[Itr] == key {
		fmt.Println(Itr)
	}
	printIndexOccurance(arr, key, Itr+1)
}

func main() {
	arr := []int{1, 2, 3, 2, 4, 2}
	// printKeyOccurance(arr, 2)
	// printIndexOccurance(arr, 2, 0)
	pac(arr, 2, 0)
}

func pac(arr []int, key, itr int) {
	if len(arr) == 0 {
		return
	}
	if arr[0] == key {
		fmt.Println(itr)
	}
	pac(arr[1:], key, itr+1)
}
