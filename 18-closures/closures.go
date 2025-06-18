/* In Go, closures are function values that reference variables from outside their body. This means that the function
“closes over” those values, allowing it to access or manipulate them even after the outer function has returned. */

package main

import "fmt"

func increment() func() int {
	count := 0
	return func() int {
		count += 1
		return count
	}
}

func main() {
	counter := increment()
	fmt.Println(counter())
	fmt.Println(counter())
}
