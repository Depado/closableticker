# closableticker
A simple closable ticker

## Usage

A simple usage for this package is when you need to perform a task periodically
and need to be able to cancel that task at some point (or restart it).

Basically a closableticker is composed with two channels, one that ticks
periodically and the other one used to notify that the ticker is closed 
(stopped). The closableticker also embeds a sync.WaitGroup in case you need to
synchronize. 

### Without WaitGroup

```go
package main

import (
	"fmt"
	"time"

	"github.com/Depado/closableticker"
)

func work(ct *closableticker.Ticker) {
	for {
		select {
		case <-ct.Ticker.C:
			fmt.Println("Working !")
		case <-ct.StopC:
			fmt.Println("Returning !")
			return
		}
	}
}

func main() {
	ct := closableticker.New(1 * time.Second)
	go work(ct)
	time.Sleep(5 * time.Second)
	ct.Stop()
	time.Sleep(1 * time.Second)
}
```

###  With WaitGroup

```go
package main

import (
	"fmt"
	"time"

	"github.com/Depado/closableticker"
)

func work(ct *closableticker.Ticker) {
	ct.WG.Add(1)
	defer ct.WG.Done()
	for {
		select {
		case <-ct.Ticker.C:
			fmt.Println("Working !")
		case <-ct.StopC:
			fmt.Println("Returning !")
			return
		}
	}
}

func main() {
	ct := closableticker.New(1 * time.Second)
	go work(ct)
	time.Sleep(5 * time.Second)
	ct.StopWait()
}
```