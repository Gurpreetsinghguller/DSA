package main

import "fmt"

func main() {
	a := 1
	for i := 0; i <= 5; i++ {
		for j := 0; j <= 5; j++ {
			if j >= 5-i {
				fmt.Print(a)
				a++
				continue
			}
			fmt.Print(" ")
		}

		fmt.Printf("\n")
	}
}

//    *
//   **
//  ***
// ****
//*****

//       *
//      **
//     ***
//    ****
//   *****
//  ******
