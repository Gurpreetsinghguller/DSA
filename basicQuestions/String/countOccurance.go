package main

import "fmt"

func main() {
	str := "Quescol provides previous year questions and answers"
	target := "a"
	fmt.Println(occurance(str, target))
}

func occurance(str, target string) int {
	a := 0
	for _, value := range str {
		if string(value) == target {
			a++
		}
	}
	return a
}
