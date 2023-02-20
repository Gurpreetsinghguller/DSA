package main

import "fmt"

func main() {
	// printFibo(8)
	fmt.Println(fiboRecur(3))
}

// 0, 1, 1, 2, 3, 5, 8 ,13,21……
func printFibo(n int) {
	a, b := 0, 1
	// fibo := []int{}
	for i := 0; i <= n; i++ {
		fmt.Println(a)
		c := a + b
		a = b
		b = c
	}
}

// 0, 1, 1, 2, 3, 5, 8, 13
func fiboRecur(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fiboRecur(n-1) + fiboRecur(n-2)
}
