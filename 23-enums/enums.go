/* Go doesn't have traditional enum types like some other languages, but you can achieve similar behavior using const and the
iota keyword. iota is a Go predeclared identifier used to simplify the definition of incrementing constants. iota starts
with 0 and increments on each line of the const. */

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
