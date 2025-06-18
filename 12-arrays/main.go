/* In Go, an array is a fixed-size, ordered collection of elements of the same type. The size of an array is part of its type,
which means arrays have a fixed length that cannot be changed after declaration. These are used to store data of different
types (a single type per array), which can be accessed using an index. */

package main

import "fmt"

func main() {
	var list [6]int
	var boolist [3]bool
	list[0] = 1
	boolist[0] = true
	fmt.Println(list, len(list))
	fmt.Println(boolist, len(boolist))

	short := [3]int{1, 2, 3}

	two := [2][3]int{{11, 22, 33}, {33, 44, 55}}

	fmt.Println(short)
	fmt.Println(two, len(two))
	fmt.Println(two[0][1], two[1][0])
}
