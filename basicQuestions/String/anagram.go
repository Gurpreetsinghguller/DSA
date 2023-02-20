package main

import "fmt"

func main() {
	st1, st2 := "quescol", "colsque"
	fmt.Println(isAnagram(st1, st2))
}

func isAnagram(st1, st2 string) bool {
	isanagram := false
	if len(st1) != len(st2) {
		return false
	}
	for _, value := range st1 {
		isanagram = false
		for _, val := range st2 {
			if value == val {
				isanagram = true
			}
		}
	}
	return isanagram
}
