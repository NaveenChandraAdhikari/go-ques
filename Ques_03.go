package main

import "fmt"

/*
Use switch to print weekday name given an integer (1–7)
*/

func printWeekday(day int) {
	switch day {

	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	default:
		fmt.Println("Invalid day")
	}

}

func getWeekday1(day int) string {
	switch day {
	case 1:
		return "Sunday"
	case 2:
		return "Monday"
	case 3:
		return "Tuesday"
	case 4:
		return "Wednesday"
	case 5:
		return "Thursday"
	case 6:
		return "Friday"
	case 7:
		return "Saturday"
	default:
		return "Invalid day"
	}
}

func getWeekday5(day int) string {
	weekdays := []string{
		"", // to make index 1 map to Sunday
		"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday",
	}
	if day >= 1 && day <= 7 {
		return weekdays[day]
	}
	return "Invalid day"
}

func main() {
	printWeekday(7)
	printWeekday(8)
	getWeekday1(3)
}

/*
is actually allowed and idiomatic in Go. In Go, lists in composite literals (like slices and arrays, maps, or structs) may optionally end with a trailing comma—even after the final element.

Why is this comma used?

It makes adding or removing items easier and reduces “diff noise” in version control.

It helps avoid syntax errors if you later add or reorder lines.

It's required if your literal spans multiple lines.

So, in Go, this is normal and even recommended style, and the trailing comma does not cause any error.

Composite Types: Arrays, slices, maps, and structs are data structures initialized with a list of values or key-value pairs. The commas separate these values, and the trailing comma is a convenience for editing multi-line lists.
Switch Statements: These are control flow constructs where case clauses define conditions, not a list of values. The syntax uses colons (:) to delimit the start of each case’s code block, and no separators (like commas) are needed between cases.
*/
