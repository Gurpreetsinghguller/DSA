package main

import "fmt"

func main() {
	str := 'b'
	fmt.Println(giveType(str))
}

func giveType(str rune) string {
	if str == 'a' || str == 'e' || str == 'i' || str == 'o' || str == 'u' {
		return "String"
	} else if str == 'A' || str == 'E' || str == 'I' || str == 'O' || str == 'U' {
		return "String"
	} else {
		return "consonant"
	}
}
