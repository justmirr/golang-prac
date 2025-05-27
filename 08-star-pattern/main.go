package main

import "fmt"

func main() {
	for i := 1; i <= 5; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println("")
	}

	fmt.Println("")
	for a := 5; a >= 1; a-- {
		for b := 1; b <= a; b++ {
			fmt.Print("*")
		}
		fmt.Println("")
	}
}
