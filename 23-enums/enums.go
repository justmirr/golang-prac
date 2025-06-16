package main

import "fmt"

type OrderStatus int

const (
	Received OrderStatus = iota
	Pending
	Completed
)

func ShowStatus(o OrderStatus) {
	fmt.Println("order status changed to:", o)
}

func main() {
	ShowStatus(Received)
}
