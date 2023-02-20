package main

import "fmt"

func main() {
	fmt.Println(primeFactor(10))
}

// 1 2 3 4 5 6 7 8 9 10
// 10, 11 12 13 14 15 16 17 18 19 20
func primeFactor(n int) []int {
	aa := []int{}
	for n%2 == 0 {
		aa = append(aa, 2)
		n = n / 2
	}
	for i := 3; i < n; i = i + 2 {
		if n%i == 0 {
			aa = append(aa, i)
			n = n / i
		}
	}
	if n > 2 {
		aa = append(aa, n)
	}
	return aa
}
