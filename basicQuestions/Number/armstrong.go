package main

import "fmt"

func main() {
	n := 27
	fmt.Println(isArmStrong(n))
}

func isArmStrong(n int) bool {
	originalNum := n
	total := 0
	for n != 0 {
		a := n % 10
		total = total + a*a*a
		n = n / 10
	}
	if originalNum == total {
		return true
	}
	return false
}
