package main

import (
	"fmt"
	"example.com/myapp/mathutils"
)

func main() {
	x := 5
	result := mathutils.Square(x)
	fmt.Printf("SQAURE OF %d IS %d\n", x, result)
}

/*
Module Init	go mod init example.com/myapp	Sets up project module
Create package	mkdir mathutils + .go	New reusable function in its own package
Use in main	import "example.com/myapp/mathutils"	Cleanly imports your own code
 */
