package main

import "fmt"

type customer struct {
	customerId   int
	customerName string
}

type order struct {
	orderId     int
	orderItem   string
	orderStatus string
	customer
}

func main() {
	oneCustomer := customer{
		customerId:   1,
		customerName: "mir",
	}

	oneOrder := order{
		orderId:     1,
		orderItem:   "chipotle",
		orderStatus: "received",
		customer:    oneCustomer,
	}

	fmt.Println(oneOrder)

	twoOrder := order{
		orderId:     2,
		orderItem:   "fries",
		orderStatus: "in progress",
		customer: customer{
			customerId:   2,
			customerName: "shaukat",
		},
	}

	fmt.Println(twoOrder)
}
