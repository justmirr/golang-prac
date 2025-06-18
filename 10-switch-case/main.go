/* The switch statement is a clean and powerful way to perform multiple conditional checks. It evaluates expressions and
executes the matching case block. Unlike some other languages, Go’s switch does not require explicit break statements -
each case breaks automatically after execution, preventing fallthrough unless explicitly stated. */

package main

import (
	"fmt"
	"time"
)

func main() {
	var zero int = 0
	switch zero {
	case 0:
		fmt.Println("Zero")
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	default:
		fmt.Println("More than three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Weekends")
	default:
		fmt.Println("Weekdays")
	}

	myType := func(value interface{}) {
		switch val := value.(type) {
		case int:
			fmt.Println("Integer")
		case string:
			fmt.Println("String")
		case bool:
			fmt.Println("Boolean")
		default:
			fmt.Println("Other", val)
		}
	}
	myType("String")
}
