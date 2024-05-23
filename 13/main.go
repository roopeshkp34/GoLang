package main

// pass	by value

import "fmt"

func updateName(x string) {
	x = "wedge"
}
func updateMenu(y map[string]float64) {
	y["coffee"] = 2.99
}
func main() {
	name := "tifa"
	updateName(name)
	fmt.Println(name)

	menu := map[string]float64{
		"pie":       3.14,
		"ice cream": 3.99,
	}
	updateMenu(menu)
	fmt.Println(menu)

}
