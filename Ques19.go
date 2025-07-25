package main
/*
Create an unbuffered channel of type string:


Send a message from a goroutine and receive it in main
 */

func main() {
	// creating an unbuffered channel of string type

	ch :=make(chan string)

	// sstart a goroutine that sends a message into the channel (from goroutine)
	go func() {
		ch <- "Hello World from Go!!"
	}()

	// receive the message in main
	msg := <-ch //this will wait until something is sent
	fmt.Println("Received",msg)
}


//function to send a message using a channel
func sendMessage(ch chan string){

	ch <-"Hellow from goroutine"

}

func main (){
	ch :=make(chan string)
	go sendMessage(ch)  //run the function in a gorotuine
	msg := <-ch // wait for the message
	fmt.Println("Received",msg)
}

/*
sending or receiving on an unbuffered channel will block until the other side is ready.

Channels let different parts of our code (goroutines) communicate safely and easily.

we don’t need to worry about manual locking or sharing data.
 */