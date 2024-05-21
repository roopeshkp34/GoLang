package main

import "fmt"

// maps
func main() {
	menu := map[string]float64{
		"soup":           4.99,
		"pie":            7.99,
		"salad":          6.99,
		"toffee pudding": 3.55,
	}
	fmt.Println(menu)
	fmt.Println(menu["pie"])

	// looping maps
	for k, v := range menu {
		fmt.Println(k, "-", v)
	}

	// ints as key type
	fmt.Println("\n\n\n ")
	phonebook := map[int]string{
		2632323:    "mario",
		123123123:  "roopesh",
		2423232232: "You",
	}
	fmt.Println(phonebook)
	fmt.Println(phonebook[2423232232])

	// update
	phonebook[2423232232] = "new user"
	fmt.Println(phonebook)

}
