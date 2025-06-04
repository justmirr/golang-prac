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
