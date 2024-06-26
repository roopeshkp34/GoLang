package main

import "fmt"

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// make new bill

func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
	}
	return b
}

// format the bill
func (b *bill) format() string {
	fs := "Bill breakdown: \n"
	var total float64 = 0

	// list items
	for k, v := range b.items {
		fs += fmt.Sprintf("%-25v ... $%v \n", k+":", v) // add 25 characters, if the word have only 2 character then it will add 23 space.
		total += v
	}

	// tip
	fs += fmt.Sprintf("%-25v ... $%v\n", "tip:", b.tip)
	// total
	fs += fmt.Sprintf("%-25v ... $%0.2f", "total:", total+b.tip)

	return fs
}

// update the tip
func (b *bill) updateTip(tip float64) {
	b.tip = tip
}

// add an item to the bill
func (b *bill) addItem(name string, price float64) {
	b.items[name] = price
}
