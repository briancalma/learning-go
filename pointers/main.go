package main

import "fmt"

func basics() {
	// A pointer is a variable that stores the memory address of another variable
	var age int = 28

	// Pointer variable that can store address of an int
	var agePointer *int

	// & operator get the address of a variable
	agePointer = &age

	fmt.Println("age:", age)
	fmt.Println("agePointer:", agePointer)

	// Pointer operators
	// & - get the address of a variable
	// * - gets the value that resides on that address

	fmt.Println("*agePointer:", *agePointer)

	// 2 behaviors of * operator
	// 1. For variable declaration - we tell go that a variable will store memory address of another variable
	// eg. var agePointer *int
	// 2. Dereference operator - tells go lang to go to the address and get / setthe value inside of it
	// eg. fmt.Println(*agePointer)
	// *agePointer = 100

	// Note: You cannot use dereference operator if the pointer address hasn't been initialized.
	// Example below will cause a panic error: `panic: runtime error: invalid memory address or nil pointer dereference`
	// var testPointer *int
	// fmt.Println(*testPointer)

	var addrPntr *string
	var sampleAddr string = "Purok Legazpi City, Albay Philippines"

	addrPntr = &sampleAddr

	fmt.Println("sampleAddr Value:", sampleAddr)
	fmt.Println("addrPntr Value:", *addrPntr)
	fmt.Println("addrPntr: ", addrPntr)

	*addrPntr = "Washington DC"

	fmt.Println("sampleAddr Value:", sampleAddr)
	fmt.Println("addrPntr Value:", *addrPntr)
	fmt.Println("addrPntr: ", addrPntr)
}

type Person struct {
	firstname string
	lastname  string
	age       int
}

func structsUsage() {
	person1 := Person{
		firstname: "Brian",
		lastname:  "Calma",
	}

	// Method 1. Using the & operator
	personPointer1 := &person1

	fmt.Println("personPointer1:", personPointer1)
	fmt.Println("personPointer1 value:", *personPointer1)

	fmt.Println("person firstname:", person1.firstname)
	fmt.Println("person lastname:", person1.lastname)

	(*personPointer1).firstname = "John"
	(*personPointer1).lastname = "Doe"

	fmt.Println("personPointer1:", personPointer1)
	fmt.Println("personPointer1 value:", *personPointer1)

	// Method 2. Using the new() function
	personPointer2 := new(Person)
	personPointer2.firstname = "Kai"
	personPointer2.lastname = "Root"

	fmt.Println("personPointer2:", personPointer2)
	fmt.Println("personPointer2 value:", *personPointer2)

	// Note: If you try to print or access  a pointer of type struct, you will find that it is going to print/get the values, which is weird
	// since there is no dereferencing operation occured. This is possible because of Go's implicit dereferencing behavior for structs.
	// But note that this is only for structs and the other types must be explicitly dereferenced.
}

func celebrateBirthdayByValue(p Person) {
	p.age++

	fmt.Printf("Inside function: %+v\n", p.age)
}

func celebrateBirthdayByPointer(p *Person) {
	p.age++

	fmt.Printf("Inside function: %+v\n", p.age)
}

func functionParameters() {
	// In go parameters are passed by value by default - which means the function gets a copy of the value.
	// When you pass a pointer you can modify the original value.

	person := Person{
		firstname: "Brian",
		lastname:  "Calma",
		age:       80,
	}

	celebrateBirthdayByValue(person)
	fmt.Println("Outside function:", person) // Age remains 80

	// celebrateBirthdayByPointer(&person)
	celebrateBirthdayByPointer(&person)
	fmt.Println("Outside function:", person) // Age is now 81
}

func main() {
	basics()
	// structsUsage()
	// functionParameters()
}

// When to use pointers with structs?
// 1. Use pointers when you need to modify the struct
// 2. Use pointers when the struct is large and you want to avoid copying
// 3. Use pointers when you want to indicate that a value might be nil

// Method receivers
// 1. Methods can have pointer receivers (s *MyStruct) or value receivers (s MyStruct)
// 2. Pointer receivers can modify the original struct
// 3. Value receivers work with copies of the struct

// Common patterns:
// - Constructor functions often return pointers to structs
// - Methods that modify state typically use pointer receivers
// - Methods that only read data can use value receivers

// Receivers
// - Receivers are use to determine "who" a method belongs to
// - Like attaching a function to a type eg. struct
