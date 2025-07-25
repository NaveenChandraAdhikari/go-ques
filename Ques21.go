package main

/*
.Write a function that uses defer to log “Function ended”
 */

// the defer keyword schedules a function call to run after the surrounding function returns
// commonly sed fr cleanup like closing files , logginf or releasing resources
//order of deferred calls run is LIO order



func myFunction(){
	defer fmt.Println("Function ended") //schedules the printing for the end
	fmt.Println("Hello world ")
}
func main() {

	myFunction()


}

//even if we return early still defer still runs
func myFunction2() {
	defer fmt.Println("Function ended")
	fmt.Println("Doing work")
	return
}


//with error handling
func task() (err error){
	defer func(){
		fmt.Println("function ended(was error:",err,")")
	}()
	err =fmt.Errorf("this is an error")
	return
}
func main()  {
	_=task() //ignoring the error value 

}
