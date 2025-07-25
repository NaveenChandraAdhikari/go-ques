package main

import (
	"errors"
	"fmt"
)

/*
Write a function which takes a slice of size n and:
( also handle the potential errors in each of the functions )

a
Appends an element


Removes the second element


Removes third last element

*/

func main() {



	s := []int{12, 3124, 4, 124, 12}

	//APPEND EXAMPLE
	s = AppendElement(s, 100)
	fmt.Println("after append:", s)

	//remove second element

	updated, err := RemoveSecondElementSafe(s)
	if err != nil {
		fmt.Println("ERROR", err)
	} else {
		fmt.Println("after remove:", updated)
		s = updated
	}

	// Remove third-last element
	updated, err = RemoveThirdLastSafe(s)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("After removing third-last:", updated)
	}

}

func AppendElement(slice []int, value int) []int {
	return append(slice, value)
}

// with error handling

func AppendElementLimit(slice []int, value int, maxLen int) ([]int, error) {
	if len(slice) >= maxLen {
		return slice, errors.New("cannot append : slice has reached maximm allowed size")

	}
	return append(slice, value), nil
}

/*
remove the second element
*/
func RemoveSecondElement(slice []int) []int {
	//removes element at indxe 1 (the second element)
	return append(slice[:1], slice[2:]...)

}

/*
with error handling,prevents index out of range issue
*/

func RemoveSecondElementSafe(slice []int) ([]int, error) {
	if len(slice) < 2 {
		return slice, errors.New("cannot remove second element :slice too small ")
	}

	//remove index 1
	return append(slice[:1], slice[2:]...), nil
}

/*
remove the third last element
*/
func RemoveThirdLast(slice []int) []int {
	idx := len(slice) - 3
	return append(slice[:idx], slice[idx+1:]...)
}

//with errors

func RemoveThirdLastSafe(slice []int) ([]int, error) {
	if len(slice) < 3 {
		return slice, errors.New("slice too short to remove third-last element")
	}
	idx := len(slice) - 3
	return append(slice[:idx], slice[idx+1:]...), nil
}
