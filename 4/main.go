package main

import "fmt"

func main() {
	age := 24
	name := "roopesh"
	//Print
	fmt.Print("Hello, ")
	fmt.Print("Word, \n") // Print method will not add a new line.

	//Println
	fmt.Println("Hello Roopesh")
	fmt.Println("By roopesh") // Println automatically add a \n at the end

	fmt.Println("my age is", age, "and my name is", name)

	// Printf (formatted strings) - %_ = format specifier
	fmt.Printf("my age is %v and my name is %v", age, name)
	fmt.Printf("\nmy age is %q and my name is %q", age, name) // it will add quotes
	fmt.Printf("\n age is type of %T", age)                   // type of the variable
	fmt.Printf("\n you scored %f points", 22.55)              // print as float
	fmt.Printf("\n you scored %0.2f points", 22.55)

	//Sprintf (Save formatted strings)
	var str = fmt.Sprintf("my age is %v and my name is %v", age, name)
	fmt.Println("\n saved string is", str)
}
