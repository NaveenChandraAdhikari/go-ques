package main
/*
Write a function that panics, and recover from panic gracefully
*/

//panic immediately stops normal function execution, unqindinding the stack .usually signals an unrecoverable error
// recover can only be used inisde a deferred function. if a panic occurs , recover() stops the panic and returns the error value so your program doesnt crash
//for critical errors or to gracefully revocer from unexpected failure like failed dISK IO or bad user input in a server



/*
//panic without recovery
func causePanic(){
	panic("Something went wrong")


}
func main(){
	causePanic()
	fmt.Println("this will not print")
}
*/

//idiomatic way panic and recover gracefully
func safeFunction(){
	defer func() {
		if r:=recover(); r!=nil{
			fmt.Println("recoverd from panic")
		}
	}()

	fmt.Println("About to panic")
	panic("serious error happend")
	fmt.Println("this will not print")

}
func main(){
	safeFunction()
	fmt.Println("program continues running after panic")
}


/*
The defer function is set up before the panic might occur.
When panic() happens, normal flow stops and Go unwraps the stack.
Deferred function runs, recover() catches the panic and resumes normal flow.

If you don’t recover, the program would crash.

When is this used in real code?
Servers: To avoid crashing the whole server if one request handler panics.

Plugins/Extensions: To sandbox untrusted code.

Resource/Cleanup: To ensure files/connections close even if things go wrong.
 */