package main

import (
	"fmt"
	"log"
	"random/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	message, err := greetings.Hello("Mir")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
}
