package main

import "fmt"

/*
Declare variables of all basic types (int,float,bool,string)
*/

func main() {

	var i int
	var f float64
	var b bool
	var s string
	fmt.Println(i, f, b, s)
}

/*
Initialized variables
*/
func main2() {
	var i int = 10
	var f float64 = 3.14
	var b bool = true
	var s string = "Hello, Go!"

	fmt.Println(i, f, b, s)
}

// IDIOMATIC GO
func main3() {
	i := 10
	f := 3.14
	b := true
	s := "Hello, Go!"
	fmt.Println(i, f, b, s)
}

/*
USE CONST TO DEFINE PI and print the area of circle
*/

func main4() {
	const pi = 3.14
	var radius float64 = 5
	var area float64 = pi * radius * radius
	fmt.Println("Area of ciecle is", area)
}

/*
with user input for area of circle
*/
func main5() {
	const pi = 3.14
	var radius float64
	fmt.Println("Enter radius")
	fmt.Scan(&radius)
	area := pi * radius * radius
	fmt.Println("Area of circle is", area)
}

/*
Write a function that takes two ints and returns the sum, difference, and product
*/
func calculate(a int, b int) (int, int, int) {

	sum := a + b
	diff := a - b
	prod := a * b

	return sum, diff, prod
}

func main6() {
	s, d, p := calculate(1, 20)
	fmt.Printf("Sum:%d Differece:%d Product :%d", s, d, p)
}
