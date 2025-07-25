package main
/*
Write a function sayHello() that prints “Hello”


Call it as a goroutine from main
 */

/*

func sayHellow(){
	fmt.Println("Hello World")
}


func main() {
sayHellow()
}


 */
//now witht go routine we can run it concurrently in its own lightweight thread

func sayHellow(){
	fmt.Println("Hello World")
}
func main(){
	go sayHellow() //run sayhellow in goroitine


	//give the gorotuine a moment to run before main exits
	// otherwise main may exit before hellow prints

	var input string
	fmt.Scanln(&input)
}

//the idiomatic code is to use sync.WaitGroup to robustly wait for gorouties to finish

func sayHellow2(wg *sync.WaitGroup) {
	fmt.Println("Hello World")
	wg.Done() //notify that this goroutine is done
}

func main(){

	var wg sync.WaitGroup
	wg.Add(1) // set number of goroutines to wait for

	go sayHellow2(&wg) //pass waitgroup by pointer

	wg.Wait() //wait for all goroutines to finish
}
/*
You add to the WaitGroup before starting the goroutine.

Each goroutine calls wg.Done() when finished.

main calls wg.Wait() to block until all have finished.
 */

/*

Giving Time for Goroutines Before main Exits
When you start a goroutine in Go, it runs in the background. The main function ends the program as soon as it finishes—regardless of whether your goroutines have finished their work. To let your goroutines finish, you must find a way to "give them a moment" before main exits.

Why Give Goroutines Time?
If you don't, your goroutine may never run or complete, because the program ends as soon as main() returns.

Ways to Let Goroutines Complete Before Exit
Here are the most common and practical methods:

1. Using time.Sleep()
Temporarily pauses the main goroutine (the main program) for a set duration.

Example:

go
go doSomething()
time.Sleep(2 * time.Second) // Waits 2 seconds, lets goroutines finish
Limitation: Not reliable for real programs—goroutines may take longer or less time.

2. Using Channels
Create a channel; have the goroutine signal on the channel when it’s done.

Main waits for the signal before exiting.

go
done := make(chan bool)
go func() {
    // work...
    done <- true // Signal completion
}()
<-done          // Wait for goroutine
Works for single or multiple goroutines (with more logic).

3. Using sync.WaitGroup (Most Idiomatic for Multiple Goroutines)
Add a counter for each goroutine; call Done() when finished.

Main waits on the WaitGroup until all are done.

go
var wg sync.WaitGroup
wg.Add(1)
go func() {
    defer wg.Done()
    // work...
}()
wg.Wait()        // Wait for all goroutines to finish
Prefer this for real projects with several concurrent actions.

4. Using Timers
Similar to sleep, but more flexible for time-based waiting.

go
timer := time.NewTimer(2 * time.Second)
<-timer.C // Blocks until timer fires
What Does "Giving a Moment" Mean?
It means ensuring main stays active just long enough for your goroutines to make progress or finish before the program ends—otherwise, their work will be lost.

Comparison Table
Method	When to Use	Reliable?	Idiomatic?
time.Sleep	Demos, quick tests	No	No
Channel	One/few goroutines	Yes	Yes
WaitGroup	Many goroutines	Yes	Best
Timer	Timed waits	Rarely	Sometimes
Key Takeaways
You CAN use time.Sleep to pause, but it’s not reliable in real programs with unpredictable execution time13.

The idiomatic way is to use channels or WaitGroup to deterministically wait for goroutines to finish.

Always use syncing tools for correct, robust concurrent programs.
 */