package main

import "fmt"

func main() {
	fmt.Println(isPrime(37))
	fmt.Println(20 / 10)
}

func isPrime(n int) bool {
	c := 2
	for c*c <= n {
		if n%c == 0 {
			return false
		}
		c++
	}
	return true
}
