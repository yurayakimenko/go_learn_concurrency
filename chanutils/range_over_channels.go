// In a [previous](range) example we saw how `for` and
// `range` provide iteration over basic data structures.
// We can also use this syntax to iterate over
// values received from a channel.

package chanutils

import (
	"fmt"
	"time"
)

// RangeOverChannels use this syntax to iterate over
// values received from a channel.
func RangeOverChannels() {

	// We'll iterate over 2 values in the `queue` channel.
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	// This `range` iterates over each element as it's
	// received from `queue`. Because we `close`d the
	// channel above, the iteration terminates after
	// receiving the 2 elements.
	for elem := range queue {
		fmt.Println(elem)
	}
}

func Timers() {
	// Timers represent a single event in the future. You
	// tell the timer how long you want to wait, and it
	// provides a channel that will be notified at that
	// time. This timer will wait 2 seconds.
	timer1 := time.NewTimer(2 * time.Second)

	// The `<-timer1.C` blocks on the timer's channel `C`
	// until it sends a value indicating that the timer
	// expired.
	<-timer1.C
	fmt.Println("Timer 1 expired")

	// If you just wanted to wait, you could have used
	// `time.Sleep`. One reason a timer may be useful is
	// that you can cancel the timer before it expires.
	// Here's an example of that.
	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 expired")
	}()
	time.Sleep(time.Second * 3)
	if timer2.Stop() {
		fmt.Println("Timer 2 stopped")
	}
}

func Tickers() {
	// Tickers use a similar mechanism to timers: a
	// channel that is sent values. Here we'll use the
	// `range` builtin on the channel to iterate over
	// the values as they arrive every 500ms.
	ticker := time.NewTicker(500 * time.Millisecond)
	go func() {
		for {
			fmt.Println("Tick at", <-ticker.C)
		}
	}()

	// Tickers can be stopped like timers. Once a ticker
	// is stopped it won't receive any more values on its
	// channel. We'll stop ours after 1600ms.
	time.Sleep(1501 * time.Millisecond)
	ticker.Stop()
	fmt.Println("Ticker stopped")
}
