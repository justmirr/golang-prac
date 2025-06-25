/* Channels in Go are built-in features that allow goroutines to communicate with each other and synchronize their execution
by passing data. They provide a safe and structured way to send and receive values between concurrent goroutines, eliminating
the need for explicit locks or shared memory management.

Channels are `Blocking` by Default -

1. When you send to a channel, the operation blocks until another goroutine is ready to receive.
2. When you receive from a channel, it blocks until a value is available.

In Go, channels are declared using the `chan` keyword followed by the type of data they will carry, such as `chan int`
for an integer channel. You create a channel using the `make` function, like `ch := make(chan int)`, and use the `<-` operator
to send and receive values. For example, `ch <- 5` sends the value 5 into the channel, while `val := <-ch` receives
a value from it. Channels are commonly used to enable safe communication between goroutines, ensuring data is passed
without the need for explicit locks. Their blocking behavior ensures that a send waits for a receive and vice versa,
which makes them useful for synchronizing the execution of concurrent tasks. Channels help in building efficient,
concurrent programs by enabling clear and structured interaction between goroutines. */

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func processChan(numChan chan int) {
	for num := range numChan {
		fmt.Println(num)
		time.Sleep(time.Second)
	}
}

func processSum(num1, num2 int, resultChan chan int) {
	resultChan <- num1 + num2
}

func processTask(doneOne chan bool) {
	defer func() { doneOne <- true }()
	fmt.Println("processing!")
}

func emailSender(emailChan chan string, doneTwo chan bool) {
	defer func() { doneTwo <- true }()
	for email := range emailChan {
		fmt.Println("sending email to:", email)
		time.Sleep(time.Second / 3)
	}
}

func main() {
	numChan := make(chan int)
	go processChan(numChan)
	for range 2 {
		numChan <- rand.Intn(100)
	}

	resultChan := make(chan int)
	go processSum(10, 20, resultChan)
	fmt.Println(<-resultChan)

	doneOne := make(chan bool)
	go processTask(doneOne)
	fmt.Println(<-doneOne)

	// Buffered channel
	emailChan := make(chan string, 5)
	doneTwo := make(chan bool)
	go emailSender(emailChan, doneTwo)

	for i := range 5 {
		emailChan <- fmt.Sprintf("%d@example.com", i)
	}
	close(emailChan)
	fmt.Println("done sending!")
	<-doneTwo

	// Receiving from multiple channels
	chan1 := make(chan int)
	chan2 := make(chan string)

	go func() { chan1 <- 111 }()
	go func() { chan2 <- "mir" }()

	for range 2 {
		select {
		case chan1Val := <-chan1:
			fmt.Println("received:", chan1Val)
		case chan2Val := <-chan2:
			fmt.Println("received:", chan2Val)
		}
	}
}
