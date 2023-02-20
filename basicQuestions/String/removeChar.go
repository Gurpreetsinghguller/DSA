package main

import (
	"fmt"
)

func main() {
	str := "quescol is best educational website"
	target := "e"
	func(stt, tar string) {
		for key, value := range stt {
			if string(value) == tar {
				string(stt[key]) = "a"
			}
		}
	}(str, target)
	fmt.Println(str)
}
