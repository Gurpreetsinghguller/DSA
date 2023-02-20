package main

import "fmt"

func main() {
	a := 7
	b := 4
	a = a - b
	b = a + b
	a = b - a // a+b-a+b = 2b = 2(a+b)
	fmt.Println(a, b)
}
