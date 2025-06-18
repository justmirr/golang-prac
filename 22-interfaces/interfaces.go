/* Interfaces are a powerful feature that allows you to define a set of method signatures without implementing them.
Any type that implements those methods automatically satisfies the interface, without needing explicit declarations.
This enables polymorphism - different types can be treated the same if they implement the same interface.
Interfaces are widely used in Go for writing flexible, modular, and testable code. */

package main

import "fmt"

type payer interface {
	pay(amount float64)
}

type payment struct {
	gateway payer
}

func (p payment) makePayment(amount float64) {
	p.gateway.pay(amount)
}

type billdesk struct{}

func (b billdesk) pay(amount float64) {
	fmt.Printf("amount %v paid using billdesk\n", amount)
}

type razorpay struct{}

func (r razorpay) pay(amount float64) {
	fmt.Printf("amount %v paid using razorpay\n", amount)
}

func main() {
	razorpayGw := razorpay{}
	billdeskGw := billdesk{}
	var input string
	fmt.Scanf(input)
	rPayment := payment{
		gateway: razorpayGw,
	}
	rPayment.makePayment(100)

	bPayment := payment{
		gateway: billdeskGw,
	}
	bPayment.makePayment(120)
}
