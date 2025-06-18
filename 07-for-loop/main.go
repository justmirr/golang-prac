/* the for loop is the only looping construct provided by the language, but it is highly flexible and can be used in a
variety of ways to fit different scenarios. Unlike other programming languages that have separate while or do-while loops,
Go uses different forms of the for loop to replicate their behavior.  */

package main

import "fmt"

func main() {
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i++
	}

	for j := 1; j <= 3; j++ {
		fmt.Println(j)
	}

	for k := range 4 {
		fmt.Println(k)
	}
}
