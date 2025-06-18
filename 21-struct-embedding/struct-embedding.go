/* In Go, struct embedding allows us to include one struct inside another. This provides a way to reuse fields and methods
from one struct in another, which is similar to *inheritance* in other languages (though Go doesn't support inheritance directly).

By embedding a struct, the outer struct gets access to the fields and methods of the embedded struct as if they were its own. */

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
