package main

// pointers
import "fmt"

func updateName(x *string) {
	*x = "wedge"
}

func main() {
	name := "tifa"

	fmt.Println("memory address of name is: ", &name)

	m := &name

	fmt.Println("memory address", m)
	fmt.Println("value at memory address:", *m)
	fmt.Println("before", name)

	updateName(m)
	fmt.Println("after ", name)

}
