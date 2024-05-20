package main

// functions
import (
	"fmt"
	"math"
)

func sayGreeting(n string) {
	fmt.Println("Good morning", n)
}
func sayBye(n string) {
	fmt.Println("Good Bye", n)
}

func cycleNames(n []string, f func(string)) {
	for _, value := range n {
		f(value)
	}
}

func circleArea(r float64) float64 {
	return math.Pi * r * r
}
func main() {
	sayGreeting("Mario")
	sayGreeting("Roopesh")
	sayBye("Mario")

	cycleNames([]string{"cloud", "tifa", "bharat"}, sayGreeting)
	cycleNames([]string{"cloud", "tifa", "bharat"}, sayBye)

	a1 := circleArea(10.5)
	a2 := circleArea(15)

	fmt.Println("area of a1", a1)
	fmt.Println("area of a2", a2)

}
