package main

import (
	"fmt"
	"log"
	"multiple/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)
	names := []string{"Mir", "Sam", "Tom"}
	messages, err := greetings.Multiple(names)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(messages)
}
