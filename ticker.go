package closableticker

import (
	"sync"
	"time"
)

// Ticker is a simple struct that allows to return from a stopped ticker
// by propagating the stop signal to another channel
type Ticker struct {
	Ticker *time.Ticker
	StopC  chan bool
	WG     sync.WaitGroup
}

// Stop stops the ticker and closes the Halt channel so the function listening
// on both channels can return
func (t *Ticker) Stop() {
	t.Ticker.Stop()
	close(t.StopC)
}

// StopWait stops the ticker and awaits for the WaitGroup to return
func (t *Ticker) StopWait() {
	t.Stop()
	t.WG.Wait()
}

// New returns a new ClosableTicker
func New(d time.Duration) *Ticker {
	return &Ticker{
		Ticker: time.NewTicker(d),
		StopC:  make(chan bool),
	}
}
