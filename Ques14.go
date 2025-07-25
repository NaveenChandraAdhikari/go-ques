package main
/*

Define an interface Shape with method Area() float64


Implement Shape interface for Circle and Rectangle structs


Write a function printArea(s Shape) that prints the area
 */

// full function
// 1. Define interface
type Shape interface {
	Area() float64
}

// 2. Circle struct & method
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14159 * c.Radius * c.Radius
}

// 3. Rectangle struct & method
type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// 4. Function taking interface
func printArea(s Shape) {
	fmt.Printf("Area: %.2f\n", s.Area())
}

func main() {
	c := Circle{Radius: 5}
	r := Rectangle{Width: 4, Height: 6}
	printArea(c) 
	printArea(r)
}


