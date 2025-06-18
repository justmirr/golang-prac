/* Goroutines are Go's lightweight threads that make it incredibly easy to run multiple tasks at the same time
without blocking your main program. Think of them as mini-workers that can handle different jobs simultaneously -
while one goroutine is waiting for a file to load, another can be processing data, and a third can be handling user input. */

/* Creating a goroutine is remarkably simple - you just put the word go in front of any function call,
and that function will run concurrently with the rest of your code. */

package main

import (
	"fmt"
	"time"
)

func main() {
	for id := 1; id <= 5; id++ {
		go func(id int) {
			fmt.Println(id)
		}(id)
	}
	time.Sleep(time.Second * 1)
}
