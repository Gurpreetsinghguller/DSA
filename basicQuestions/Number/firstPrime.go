package main

import "fmt"

func main() {
	fmt.Println(FirstPrime(5))
}

// 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16
func FirstPrime(n int) []int {
	arr, i := []int{}, 3
	arr = append(arr, 2)
	if n == 1 {
		return arr
	}
	for {
		result := isPrime(i)
		if result {
			arr = append(arr, i)
		}
		i = i + 2
		if len(arr) == n {
			return arr
		}
	}
}

func isPrime(n int) bool {
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
