package main

import "fmt"

// receiver functions

func main() {
	myBill := newBill("marios bill")
	fmt.Println(myBill)

	fmt.Println(myBill.format())

}
