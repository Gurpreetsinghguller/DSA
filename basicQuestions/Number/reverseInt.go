package main

import "fmt"

func main() {
	fmt.Println(ReverseInt(125))
}

func ReverseInt(n int) int {
	reverse := 0
	for n != 0 {
		a := n % 10
		reverse = reverse*10 + a
		n = n / 10
	}
	return reverse
}
