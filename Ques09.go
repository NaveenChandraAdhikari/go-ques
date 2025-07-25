package main
import "fmt"
/*
Define a struct Person with fields Name, Age


Write a method Greet() for Person that prints “Hello, my name is X”

 */

/*
With &:
p is a pointer (*Person), holding the memory address of the struct.
fmt.Println(p) shows the pointer notation (&{Naveen 21}).
If you called Greet(p), you’d get a compilation error because Greet expects a Person, not a *Person. You’d need to pass Greet(*p) to dereference the pointer.
Without &:
p is a Person struct.
fmt.Println(p) shows just the struct’s data ({Naveen 21}).
You can call Greet(p) directly since Greet expects a Person.
 */
type Person struct {
	Name string
	Age int
}

//its a function not a method yet
func Greet(p Person){
	fmt.Printf("Hello,my name is %s \n", p.Name)
}

//Idiomatic go  ,attach greet as a method ,

type Person1 struct {
	Name string
	Age int
}
// Method for Person1
func(p Person1)Greet1() {
	fmt.Printf("Hello,my name is %s \n", p.Name)
}

// Method 4 if we need or use a pointer receiver for the method
func (p *Person) Greet2() {
	fmt.Printf("Hello, my name is %s\n", p.Name)
}


func main()  {
	//p:= &Person{Name:"Naveen",Age:21}
	//p is like making an instance of the Person struct, like creating an object in Java:
	p:= Person{Name:"Naveen",Age:21}
	fmt.Println(p)
	Greet(p)
	//Greet(*p)

	person:=Person1{Name:"Nitin",Age:21}
	person.Greet1()
	p.Greet2()


}



/*

In Go, methods can use either value receivers ((p Person)) or pointer receivers ((p *Person)).

Use a pointer if you might modify the struct or want to avoid copying large structs.

For this Greet method (which only prints), a value receiver is fine, but the pointer receiver is best practice when designing larger or more complex types.



The (p Person1) makes Greet1 a method for Person1. The p here is the receiver, representing the Person1 instance the method is called on.
In Java, this is like an instance method:
java



class Person1 {
    String name;
    int age;
    void greet1() {
        System.out.printf("Hello, my name is %s\n", name);
    }
}
The p in (p Person1) is like this in Java—it refers to the instance of Person1 that calls Greet1.
You call it as person.Greet1() in Go, or person.greet1() in Java.

In func (p Person1) Greet1(), the p is the receiver, a parameter that represents the Person1 instance the method is called on.
When you write person.Greet1(), Go passes person as p to the method.
In Java, this is implicit with this:
java



person.greet1(); // In Java, 'this' inside greet1() refers to person
In Go, the receiver (p) is explicit, and you choose whether it’s a value (Person1) or pointer (*Person1).
In your code, p is a value receiver (Person1), so Greet1 gets a copy of person. If it were (p *Person1), it would get a pointer to person.
 */