package main

import (
	"addition"
	"division"
	"fmt"
	"multiplication"
	"subtraction"
)

func main() {
	var first, second float64
	var operator string
	fmt.Print("Enter an expression (e.g. 23 + 43): ")
	fmt.Scanf("%g %s %g", &first, &operator, &second)
	switch operator {
	case "+":
		fmt.Println("Result:", addition.Plus(first, second))
	case "-":
		fmt.Println("Result:", subtraction.Minus(first, second))
	case "*":
		fmt.Println("Result:", multiplication.Multiply(first, second))
	case "/":
		fmt.Println("Result:", division.Divide(first, second))
	default:
		fmt.Println("Incorrect Operator")
	}
}
