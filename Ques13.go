package main
/*
Create an int variable, get its pointer, print the address & value via pointer


Write a function increment(n *int) that increments the value by 1

 */

func main() {
	var x int=10

	//get its pointer
	ptr:=&x

	//print the address and the value via pointer
	fmt.Println("pointer address:",ptr)
	fmt.Println("value at address",*ptr) //dereferncing
}

/*
for increament functipn

 */

func increament(n *int){
	*n=*n+1 //derefernce and update the value
}

func main() {
	x:=10
	fmt.Println("before increamener",x)
	//pasing the address to the function, *n refers to the original x
	increament(&x)
	fmt.Println("after increamener",x)
}

/*
idomatic slution
 */
// increment increments the value pointed to by n
func increment(n *int) {
	// n now hols teh address of x or any variable you pass eg:increamenet(ptr)
	*n = *n + 1 // change value at location via pointer in func
}

func main() {
	x := 10
	ptr := &x //get address of a variable, holds the location of x

	fmt.Println("Pointer address:", ptr) //show mem address
	fmt.Println("Value at address:", *ptr) //derefernce pointer to get/set actual value

	fmt.Println("Before increment:", x)
	increment(ptr)
	fmt.Println("After increment:", x)
}



/*

ILLUSTRATE :
x is at address 0x1234, value 10.
ptr is at address 0x5678, value 0x1234 (points to x).
When we call increment(ptr), n is a new variable (say, at address 0x9abc) that also holds the value 0x1234.
n and ptr both point to 0x1234 (the address of x).
*n = *n + 1 changes the value at 0x1234 from 10 to 11
 */