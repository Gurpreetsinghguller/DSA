package main

import "fmt"

func main() {
	arr := []int{5, 4, 3, 2, 1}
	for i := 0; i < len(arr)-1; i++ {
		min := arr[i]
		for j := i + 1; j < len(arr); j++ {
			if min > arr[j] {
				temp := min
				min = arr[j]
				arr[j] = temp
			}
		}
		arr[i] = min
	}
	fmt.Println(arr)
}
