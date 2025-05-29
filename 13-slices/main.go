package main

import "fmt"

func main() {
	var one []int
	fmt.Println(one, len(one))
	var two = make([]int, 2)
	fmt.Println(two, len(two))

	two[0] = 50
	two[1] = 100
	two = append(two, 150)

	fmt.Println(two, len(two))
	var three = make([]int, len(two))
	copy(three, two)
	fmt.Println(three, len(three))
}
