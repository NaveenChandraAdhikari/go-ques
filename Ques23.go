package main

/*
Write a function that waits for a value on a channel but times out after 2 seconds
*/

//GO concurrency releaed ques that teaches how to handle timeouts and avoid deadlocks.

//channel receive in go will block until a value is sent
// sometimes we dont want to wait forrever so we use a timeout
// GO's select statement combined with time.After we use here


//waitOrTimeout waits for a string from ch, or exits after 2 seconds

func waitOrTimeout(ch chan string) {
	select {
	case msg := <-ch:
		fmt.Println("Received message:", msg)
	case <-time.After(2 * time.Second):
		fmt.Println("Timed out after 2 seconds")
	}
}

func main() {
	ch := make(chan string)
	go func() {
		// we use  next line to test successful receive:
		// time.Sleep(1 * time.Second); ch <- "hello"
		// for timeout test we can seee from here
	}()
	waitOrTimeout(ch)
}
/*
Value received---->	Receives and prints message
No value, timeout first	Prints "Timed out after 2 seconds"
Both arrive at once	Go picks one at random
 */