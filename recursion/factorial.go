package main

import "fmt"

func fact(n int) int {
	// cover the base case
	if n == 0 {
		return 1
	}

	smallerAns := fact(n - 1)

	biggerAns := n * smallerAns

	return biggerAns
}

func main() {
	fmt.Println(fact(8))
}

// // 8,7,6,5,4,3,2,1
// func fact(n int) int {
// 	if n == 0 {
// 		return 0
// 	}
// 	if n == 1 {
// 		return 1
// 	}
// 	return n * fact(n-1)
// }
