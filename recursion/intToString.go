package main

import "fmt"

func main() {
	num := 20481

	fmt.Println(intToString(num))
}

func intToString(num int) string {
	if num == 0 {
		return ""
	}
	a := num % 10
	switch a {
	case 0:
		return intToString(num/10) + "Zero "
	case 1:
		return intToString(num/10) + "One "
	case 2:
		return intToString(num/10) + "Two "
	case 3:
		return intToString(num/10) + "Three "
	case 4:
		return intToString(num/10) + "Four "
	case 5:
		return intToString(num/10) + "Five "
	case 6:
		return intToString(num/10) + "Six "
	case 7:
		return intToString(num/10) + "Seven "
	case 8:
		return intToString(num/10) + "Eight "
	case 9:
		return intToString(num/10) + "Nine "

	}
	return ""
}
