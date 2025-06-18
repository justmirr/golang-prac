/* Usually, when we pass a variable to a function in Go, the function receives a copy of the value, i.e., by the
‘call by value’ method. This means any changes made inside the function do not affect the original variable.
However, when we want to modify the original value, we use pointers. A pointer holds the memory address (reference) of a
variable instead of its value. By passing a pointer to a function using the & operator, the function can access
and change the original variable using dereferencing (with the * operator). This allows the function to work
directly with the original data, not just a copy. */

package main

import "fmt"

func change(value *int) {
	*value = 100
}

func main() {
	value := 10
	fmt.Println("value before change:", value)
	change(&value)
	fmt.Println("value after change:", value)
}
