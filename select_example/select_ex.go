package selections

import (
	"fmt"
	"time"
)

// SelectTest use this to test select usage with goroutines in Golang
func SelectTest() {

	// For our example we'll select across two channels.
	c1 := make(chan string)
	c2 := make(chan string)

	// Each channel will receive a value after some amount
	// of time, to simulate e.g. blocking RPC operations
	// executing in concurrent goroutines.
	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one.1"
		time.Sleep(1 * time.Second)
		c1 <- "one.2"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two.1"
		time.Sleep(2 * time.Second)
		c2 <- "two.2"
	}()

	// We'll use `select` to await both of these values
	// simultaneously, printing each one as it arrives.
	for i := 0; i < 4; {
		time.Sleep(100 * time.Millisecond)
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
			i++
		case msg2 := <-c2:
			fmt.Println("received", msg2)
			i++
		default:
			fmt.Println("no activity")
		}
	}
}
