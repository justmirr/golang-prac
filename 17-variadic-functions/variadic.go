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
