package main

import "fmt"

/*
Declare an array of 5 integers, initialize and print them
*/

func main() {
	var arr [5]int //declaration
	arr[0] = 10    // initialization
	arr[1] = 20
	arr[2] = 30
	arr[3] = 40
	arr[4] = 50
	fmt.Println(arr[0])
	fmt.Println(arr[1])
	fmt.Println(arr[2])
}

func method2() {
	var arr = [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr)
}

func method3() {
	arr := [5]int{1, 2, 3, 4, 5}
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

}

func method4() {
	arr := [5]int{1, 2, 3, 4, 5}
	for i, v := range arr {
		fmt.Println("arr[%d]=%d\n", i, v)
	}
}

func method5() {
	// Compiler infers the array length
	arr := [...]int{10, 20, 30, 40, 50}

	// Print each element
	for i, v := range arr {
		fmt.Printf("arr[%d] = %d\n", i, v)
	}
}
