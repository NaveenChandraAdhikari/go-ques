package main



type Rectangle struct {
	Width,Height float64
}
func Area(r Rectangle) float64 {
	return r.Width * r.Height
}

func Perimeter(r Rectangle) float64 {
	return 2 * (r.Width + r.Height)
}



// Idiomatic way is to use a value receiver (no mutation needed)
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}


// we can aslo use pointer reveiver, if we have to mutate the fields in methods
/ Pointer receiver method for Area (doesn't need pointer, but for uniformity)
func (r *Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r *Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func (r *Rectangle) SetSize(width, height float64) {
	r.Width = width
	r.Height = height
}

func main() {
	rect:=Rectangle{Width:10, Height:20}
	fmt.Println("Area:",Area(rect))
	fmt.Println("Perimeter:", Perimeter(rect))


	//////methods use a value receiver because no fileds are modified
	rect:=Rectangle{Width:10, Height:20}
	fmt.Println("Area:",rect.Area())
	fmt.Println("Perimeter:",rect.Perimeter())
/*
   When we call methods like rect.Area(), Go automatically takes the address if needed, so both pointer and non-pointer variables work.
 */
	//  change size using pointer receiver method
	rect.SetSize(10, 4.5)
	fmt.Printf("After resize: width=%.2f, height=%.2f\n", rect.Width, rect.Height)
	fmt.Printf("Area: %.2f\n", rect.Area())
	fmt.Printf("Perimeter: %.2f\n", rect.Perimeter())
}

