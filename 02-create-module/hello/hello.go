package main

import (
	"fmt"
	"the-second/greetings"
)

func main() {
	message := greetings.Hello("Mir")
	fmt.Println(message)
}
