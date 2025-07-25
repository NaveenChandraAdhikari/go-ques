package main
/*
Write two goroutines that each send a string to a channel after different delays


Use select to print whichever message arrives first

 */

// gorotuines lightweight concurrent functions
// channels are pipe for communication ; we use make(chan string) for unbuffered string channel
// delays : for delays we use time.sleep() inside gorotuines to simulate work or delay
//select statement lets us wait fro multiple channle operation proceeds witht he one thats ready first

func main() {
	ch :=make(chan string)

	//first goroutine : short delay
	go func() {
		time.Sleep(1*time.Second)
		ch <- "message from gorotuine 1"
	}()


	//second gorotuine :longer delay
	go func() {
		time.Sleep(2*time.Second)
		ch <- "message from gorotuine 2"
	}()

	//select to get whichever arrives first
	/*
	select checks multiple channel cases. The first message to arrive is printed.
	Whichever goroutine sleeps less, its message arrives first.
	 */
	select {
	case msg := <-ch:
		fmt.Println("Received",msg)
	}
}

//IDIOMATIC WAY

import (
"fmt"
"time"
)

// sendWithDelay sends a message to ch after a delay
func sendWithDelay(ch chan string, msg string, delay time.Duration) {
	time.Sleep(delay)
	ch <- msg
}

func main() {
	ch := make(chan string)

	go sendWithDelay(ch, "Fast goroutine!", 1*time.Second)
	go sendWithDelay(ch, "Slow goroutine!", 3*time.Second)

	select {
	case m := <-ch:
		fmt.Println("First received:", m)
	}
}

/*
if we want to print both message as they arrive we use a loop with select
for i := 0; i < 2; i++ {
    select {
    case msg := <-ch:
        fmt.Println("Received:", msg)
    }
}
 */

/*
SOME NOTES:
In Go, a delay means intentionally pausing the execution of a goroutine (a lightweight thread) for a certain amount of time.
Delays are commonly used to:

Simulate work or waiting (like waiting for a resource or external system)

Control timing (rate limiting API calls, pacing operations)

Retry operations after failures

Synchronize events or animations
///////////////////////////////////////////////\\\\\\\\\\\\\\\\\\\
Using time.Sleep() for Delays
Go provides the time.Sleep(duration) function in the time package.
It suspends the current goroutine (not the whole program) for the specified duration.


Go offers better tools than just time.Sleep():

Alternative	When to Use	Example Use Case
time.Ticker	Run code at regular intervals (events every N seconds)	Heartbeat, regular polling
time.Timer	Wait for a single event after N seconds, or trigger something after a delay	One-off timeout, single retry
time.After	Used with select for waiting but can unblock or abort if needed (interruptible wait)	Graceful shutdowns, timeouts
sync.WaitGroup	Wait for one or more goroutines to finish (not a delay, but for waiting)	Concurrency synchronization
Channels + Select	Combine channels with timeouts to wait or abort early	Responsive, cancelable waits





Use Cases
API Rate Limiting: Add a delay between requests to stay within allowed limits (e.g., sleep 1 second inside a loop when sending requests to avoid exceeding 60 requests/minute).

Load Shedding: Space out expensive operations to prevent server overload.

Animations/UI Throttling: Add delays between updating frames in a CLI spinner, game loop, or other timed tasks.

Polling: Wait a fixed interval between checks for new messages, jobs, data, etc..

Poor or Bad Use Cases
Responding to events or signals: Don’t use time.Sleep(). Instead, react to actual events with channels and select.

Handling shutdowns: Use context, channels, or timers, not sleep, so programs can respond quickly.

Waiting for other goroutines: Use sync.WaitGroup or channels for precise control, rather than guessing with sleep

When to Use Which
Tool	Use For	Pro/Con
time.Sleep	Simple, demo/test delays	Easy, not interruptible, not precise
time.Ticker	Regular intervals	Can stop/start, great for loops
time.Timer	One-time future trigger	Can cancel, more control
select + time.After	Responsive timeouts	Can abort, best for production waits
WaitGroup, Channels	Synchronization/waiting	Safe, idiomatic for goroutine control
*/