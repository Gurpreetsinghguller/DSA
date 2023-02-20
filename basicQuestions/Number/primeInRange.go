package main

import "fmt"

func main() {
	fmt.Println(firstNprime(4))
}

func firstNprime(n int) []int {
	arr := []int{}
	if n == 1 {
		return arr
	}
	// arr = append(arr, 2)
	for i := 2; i <= n; i++ {
		result := isPrime(i)
		if result {
			arr = append(arr, i)
		}
	}
	return arr
}

func isPrime(n int) bool {
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
