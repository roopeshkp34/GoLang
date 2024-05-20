package main

// package scope
import "fmt"

var score = 99.5

func main() {
	sayHello("Roopesh")

	for _, v := range points {
		fmt.Println(v)
	}
	showScore()
}
