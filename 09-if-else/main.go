/* These are used to control the flow based on conditions. The condition must not be enclosed in parentheses,
but braces {} are required. You can also declare variables inside the if statement, scoped only to that block. */

package main

import "fmt"

func main() {
	number := 18
	if number > 18 {
		fmt.Println("Greater than 18")
	} else if number == 18 {
		fmt.Println("Equal to 18")
	} else {
		fmt.Println("Less than 18")
	}
}
