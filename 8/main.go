package main

import "fmt"

// conditions and booleans

func main() {
	age := 45
	fmt.Println(age <= 50)
	fmt.Println(age >= 50)
	fmt.Println(age == 45)
	fmt.Println(age != 50)

	if age < 30 {
		fmt.Println("Age is less than 30")
	} else if age < 40 {
		fmt.Println("age is less than 40")
	} else {
		fmt.Println("Age is not less than 45")
	}
	print("\n")

	names := []string{"mario", "luigi", "yoshi", "peach", "bowser"}

	for index, value := range names {
		if index == 1 {
			fmt.Println("\nContinuing at pos", index)
			continue
		}
		if index > 2 {
			fmt.Println("\n breaking at pos", index)
			break
		}
		fmt.Printf("\nthe value at position %v is %v", index, value)
	}
}
