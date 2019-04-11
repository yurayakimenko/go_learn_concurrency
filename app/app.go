// Go's _select_ lets you wait on multiple channel
// operations. Combining goroutines and channels with
// select is a powerful feature of Go.

package main

import "github.com/yy/routines/chanutils"

func main() {
	// selections.SelectTest()
	// chanutils.ChanTest()
	// chanutils.ClosingChannels()
	chanutils.WorkerPools()
}
