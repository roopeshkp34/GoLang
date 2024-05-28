package main

import "fmt"

// receiver functions with pointers

func main() {
	myBill := newBill("marios bill")
	fmt.Println(myBill)
	myBill.updateTip(10)

	myBill.addItem("onion soup", 4.50)
	myBill.addItem("veg pie", 8.95)
	myBill.addItem("toffee pudding", 4.95)
	myBill.addItem("coffee", 3.25)

	fmt.Println(myBill.format())

}
