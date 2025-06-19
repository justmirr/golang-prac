/* A WaitGroup in Go is a synchronization tool used to wait for multiple goroutines to complete, offering a more reliable
and efficient alternative to using time.Sleep(). In real-world applications, we can't predict exactly how long each
goroutine will take to finish, so instead of hardcoding delays, we use a WaitGroup to dynamically manage the execution.
We create a WaitGroup variable, add a counter for the number of goroutines we want to wait for using Add(),
and then call Done() inside each goroutine once it finishes. Finally, the main function calls Wait(),
which blocks further execution until all the registered goroutines have completed. */

package main

import (
	"fmt"
	"sync"
)

func printID(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(id)
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i <= 5; i++ {
		wg.Add(1)
		go printID(i, &wg)
	}
	wg.Wait()
}
