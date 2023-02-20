package main

import "fmt"

func power(a, n int) int {
	if n == 0 {
		return 1
	}
	return a * power(a, n-1)
}

func main() {
	// fmt.Println(power(2, 3))
	fmt.Println(pow(2, 3))
}

func pow(a, b int) int {
	if b == 0 {
		return 1
	}
	return a * pow(a, b-1)
}
