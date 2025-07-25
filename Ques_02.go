package main

import "fmt"

/*
Write a function using if-else to check if a number is odd or even
*/

/*
if else odd even function
*/

func oddOrEven(n int) string {
	if n%2 == 0 {
		return "even"
	} else {
		return "odd"
	}
}

func oddOrEven2(n int) string {
	if n%2 == 0 {
		return "even"
	}
	return "odd"
}

func oddOrEven3(n int) bool {
	return n%2 == 0
}

func main() {
	number := 5
	result := oddOrEven(number)
	fmt.Println(result)

}
