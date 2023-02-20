package main

import "fmt"

func fibo(n int) int {
	// base condition

	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	// breaking down into smaller one
	smallerAns := fibo(n - 1)

	// bigger problem
	biggerAns := fibo(n-2) + smallerAns

	return biggerAns
}

func main() {
	// 0,1,1,2,3,5,8,13,21,34,55
	fmt.Println(fibo(9))
}
