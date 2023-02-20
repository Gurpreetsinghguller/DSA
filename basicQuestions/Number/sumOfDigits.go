package main

import "fmt"

func main() {
	fmt.Println(sumOfDigits(1547))
}

func sumOfDigits(n int) int {
	if n == 0 {
		return 0
	}
	return n%10 + sumOfDigits(n/10)
}
