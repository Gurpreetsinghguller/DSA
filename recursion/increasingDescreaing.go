package main

import "fmt"

func decreasing(n int) {
	if n == 0 {
		return
	}
	fmt.Print(n, " ")
	decreasing(n - 1)
}

func increasing(n int) {
	if n == 0 {
		return
	}
	increasing(n - 1)
	fmt.Print(n, " ")
}

func incWithItr(n, i int) {
	if n == 0 {
		return
	}
	fmt.Print(i, " ")
	// incWithItr(n-1, i+1)
	incWithItr(n-1, i+1)
}

func main() {
	// decreasing(10)
	// increasing(10)
	incWithItr(10, 1)
}
