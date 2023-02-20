package main

import "fmt"

// var sum int

func main() {
	// fmt.Println(palindrome(151))
	n := palindromeRecurr(151)
	fmt.Println(n)
	if n == 151 {
		fmt.Println(true)
	}
}

func palindrome(n int) bool {
	nn := n
	total := 0
	for n != 0 {
		a := n % 10
		total = total*10 + a
		n = n / 10

	}
	if total == nn {
		return true
	}
	return false
}

// 151
func palindromeRecurr(n int) int {
	var rem, sum int

	y := func() {
		if n != 0 {
			rem = n % 10
			sum = sum*10 + rem

		}
	}
	y()
	palindromeRecurr(n / 10)
	return sum
}

// 1
