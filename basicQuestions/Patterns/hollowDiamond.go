package main

import "fmt"

func main() {
	n := 5
	for i := 0; i < n*2-1; i++ {
		for j := 0; j < n*2-1; j++ {
			if i >= 1 && i < n*n-2 {
			}
			fmt.Print("*")
		}

		fmt.Print("\n")
	}
}
