package main

import "fmt"

func main() {
	fmt.Println(isPerfect(6))
}

func isPerfect(n int) bool {
	total := 0
	for i := 1; i < n; i++ {
		if n%i == 0 {
			total = total + i
		}
	}
	if total == n {
		return true
	}
	return false
}
