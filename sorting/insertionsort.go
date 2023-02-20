package main

import "fmt"

func main() {
	arr := []int{5, 2, 6, 9, 4}
	for i := 1; i < len(arr); i++ {
		temp := arr[i]
		j := i - 1
		for j >= 0 {

			if arr[j] > temp {
				arr[j+1] = arr[j]
			} else {
				break
			}
			j--
		}
		arr[j+1] = temp

	}

	fmt.Println(arr)
}
