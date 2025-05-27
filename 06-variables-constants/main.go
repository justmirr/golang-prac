package main

import "fmt"

const (
	port = 8080
	host = "localhost"
)

var (
	department = "technical"
)

func main() {
	var name string = "mir"
	fmt.Println(name, department)
	fmt.Println(port, host)

	var interest = "music"
	game := "valorant"
	fmt.Println(interest, game)
}
