/* Functions in Go are blocks of reusable code that perform specific tasks. They help organize code, make it more readable,
and avoid repetition. A function in Go is defined using the func keyword, followed by a name, parameters,
and a return type (if any). Go also supports functions with multiple return values and allows functions
to be passed as arguments, making them flexible and powerful tools for building modular programs. */

package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func multiple() (int, string) {
	return 1, "Mir"
}

func one(func(a, b int) int) {
	result := two(20, 20)
	fmt.Println(result)
}

func two(a, b int) int {
	return a + b
}

func three() func(a, b int) int {
	return func(a, b int) int {
		return a + b
	}
}

func main() {
	fmt.Println(add(10, 20))
	id, name := multiple()
	fmt.Println(id, name)
	one(two)
	fmt.Println(three()(19, 21))
	value := three()
	fmt.Println(value(100, 200))
}
