package main

import "fmt"

func main() {
	fmt.Println(fact(5))
	fmt.Println(factRecurr(5))
}

// 8 7 6 5 4 3 2 1
func fact(n int) int {
	total := 1
	for i := n; i > 1; i-- {
		total = total * i
	}
	return total
}

// 8,7,6,5,4,3,2,1
func factRecurr(n int) int {
	if n == 1 {
		return 1
	}
	return n * factRecurr(n-1)
}
