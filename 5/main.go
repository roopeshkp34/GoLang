package main

import "fmt"

// Arrays & Slices
func main() {
	// Array
	var ages [3]int = [3]int{20, 25, 30} // an array with 3 integer , in go the arrays have fixed length.
	var number = [3]int{20}

	names := [4]string{"roopesh", "roopesh", "me", "then me"}
	names[1] = "new"
	fmt.Println(ages, len(ages), number)
	fmt.Println(names, len(names))

	// Slices (use arrays under the hood)
	var scores = []int{100, 50, 60}
	scores[1] = 23
	scores = append(scores, 85) // it just return a new slice. if we need to update a slice assign to a new variable.
	fmt.Println(scores)

	//slice ranges
	rangeOne := names[1:3]
	rangeTwo := names[2:]
	rangeThree := names[:3]

	fmt.Println(rangeOne, rangeTwo, rangeThree)
}
