package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// user input

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')
	return strings.TrimSpace(input), err

}

func createBell() bill {
	reader := bufio.NewReader(os.Stdin)

	// fmt.Print("Create a new bill name: ")
	// name, _ := reader.ReadString('\n')
	// name = strings.TrimSpace(name)

	name, _ := getInput("Create a new bill name: ", reader)

	b := newBill(name)
	fmt.Println("Created the bill - ", b.name)
	return b
}
func promptOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)
	opt, _ := getInput("Choose Option (a - add item, s - save bill, t - add tip) : ", reader)
	fmt.Println(opt)

}
func main() {

	myBill := createBell()
	promptOptions(myBill)

}
