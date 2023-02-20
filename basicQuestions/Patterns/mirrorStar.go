package main

import "fmt"

func main() {
	n := 4
	// for i := 0; i < n; i++ {
	// 	for j := 0; j <= n; j++ {
	// 		if j >= n-i {
	// 			fmt.Print("*")
	// 			continue
	// 		}
	// 		fmt.Print(" ")
	// 	}
	// 	fmt.Print("\n")
	// }
	// for i := 0; i < n; i++ {
	// 	for j := 0; j < n; j++ {
	// 		if j >= i {
	// 			print("*")
	// 		}
	// 	}
	// 	print("\n")
	// }
	// for i := 0; i < n; i++ {
	// 	for j := 0; j < n; j++ {
	// 		if j >= i {
	// 			print("*")
	// 			continue
	// 		}
	// 		print(" ")
	// 	}
	// 	print("\n")
	// }
	// for i := 0; i < n; i++ {
	// 	for j := n - i; j >= i; j-- {
	// 		fmt.Print("*")
	// 		if j == n {
	// 			break
	// 		}
	// 	}
	// 	print("\n")
	// }
	counter := 1
	for i := 0; i < n; i++ {
		for j := 0; j < 2*n-1; j++ {
			if j > n-2*i+1 {
				fmt.Print(counter)
				counter++
				continue
			}
			print(" ")
		}
		counter = 1
		print("\n")
	}
}
