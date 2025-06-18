/* Slices are a powerful and flexible alternative to arrays. Unlike arrays, slices are dynamically sized and are used in
most real-world applications. A slice provides a more convenient and efficient way to work with sequences of data.

Slices have a capacity that dynamically increases as values are appended to them. */

package main

import (
	"fmt"
	"slices"
)

func main() {
	var one []int
	fmt.Println(one, len(one))
	var two = make([]int, 2)
	fmt.Println(two, len(two))

	two[0] = 1
	two[1] = 2
	two = append(two, 150)

	fmt.Println(two, len(two))
	var three = make([]int, len(two))
	copy(three, two)
	fmt.Println(three, len(three))

	var four = []int{1, 2}
	fmt.Println(four[:])
	fmt.Println(slices.Equal(two, three))

	var five = [][]int{{1, 2, 3}, {1, 2, 3}}
	fmt.Println(five)
}
