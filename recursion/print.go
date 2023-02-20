package main

import "fmt"

func main() {
	print(5)
}

func print(n int) {
	if n == 1 {
		fmt.Println(n)
		return
	}
	fmt.Println(n)
	print(n - 1)
}
