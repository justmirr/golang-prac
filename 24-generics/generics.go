/*
Generics in Go allow you to write functions, types, and data structures that can work with multiple types
while maintaining type safety at compile time. This feature eliminates code duplication, improves type safety
by catching errors at compile time rather than runtime, and makes Go code more reusable and maintainable without
sacrificing the language's emphasis on simplicity and performance.
*/

package main

import "fmt"

func printSlice[T interface{}](slice []T) {
	for _, v := range slice {
		fmt.Printf("%v ", v)
	}
	fmt.Println()
}

type stack[T any] struct {
	elements []T
}

func main() {
	intSlice := []int{1, 10, 20, 50}
	stringSlice := []string{"mir", "mozin"}
	printSlice(intSlice)
	printSlice(stringSlice)

	myStack := stack[int]{
		elements: []int{1, 3, 6},
	}

	fmt.Println(myStack)
}
