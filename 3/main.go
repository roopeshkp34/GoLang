package main

import "fmt"

func main() {
	//strings
	var nameOne string = "Roopesh"
	var nameTwo = "new name"
	var nameThree string

	fmt.Println(nameOne, nameTwo, nameThree)
	nameOne = "Roopesh2"
	nameThree = "Spiderman"
	fmt.Println(nameOne, nameTwo, nameThree)

	nameFour := "batman" // in this case no need to add var keyword, its only possible inside a function.

	fmt.Println("\n", nameFour)

	// ints
	var ageOne int = 20
	var ageTwo = 30
	ageThree := 40

	fmt.Println("\n", ageOne, ageTwo, ageThree)

	// bits & memory
	var numOne int8 = 25
	var numTwo int8 = -128
	var numThree uint16 = 256 //uint not allowed negative numbers.

	fmt.Println("\n bites", numOne, numTwo, numThree)

	// floats
	var scoreOne float32 = 25.98
	var scoreTwo float64 = 89898989898983.12232323232
	scoreThree := 123.33

	fmt.Println("\n floats", scoreOne, scoreTwo, scoreThree)

}
