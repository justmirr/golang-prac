/* Variadic functions in Go allow you to pass a variable number of arguments of the same type to a function.
This is useful when you don't know in advance how many values will be passed. You define a variadic function
using ... before the type, like func variadic(value ...int). Inside the function, values acts like a slice,
so you can loop through it and perform operations as needed. */

package main

import (
	"fmt"
)

func variadic(values ...int) int {
	total := 0
	for _, value := range values {
		total += value
	}
	return total
}

func main() {
	fmt.Println(variadic(10, 20, 30, 40))
}
