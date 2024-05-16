package main

// For loop
import "fmt"

func main() {
	x := 0
	for x < 5 {
		fmt.Println("The value of x is", x)
		x++
	}
	println("\nNormal for loop")
	for i := 0; i < 5; i++ {
		fmt.Println("The value of I is", i)
	}

	println("\niterate through slices")
	names := []string{"mario", "luigi", "yoshi", "peach"}
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	print("\n loop with rang\ne")
	for index, value := range names {
		fmt.Printf("The value at index %v is %v \n", index, value)
	}

	print("\n understand usage of _\n")
	for _, value := range names {
		fmt.Printf("The value at is %v \n", value)
	}
}
