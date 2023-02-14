package main

import (
	"fmt"
	"os"
)

type Bill struct {
	name  string
	items map[string]float64
	tip   float64
}

func newBill(name string) Bill {
	bill := Bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
	}
	return bill
}

// creating struct functions or methods (reciever function)

func (b Bill) breakDown() string {
	fs := "Bill breakdown: \n"

	var total float64 = 0

	for key, value := range b.items {
		fs += fmt.Sprintf("%-20v ...$%v \n", key+":", value)
		total += value
	}
	fs += fmt.Sprintf("%-20v ...$%0.2f \n", "tip", b.tip)
	fs += fmt.Sprintf("%-20v ...$%0.2f", "total:", total+b.tip)

	return fs

}

// update tip
func (b *Bill) updateTip(tip float64) {
	b.tip = tip

}

// add item
func (b *Bill) addItem(item string, price float64) {
	if len(item) >= 0 {
		b.items[item] = price
	}

}

// save
func (b *Bill) save() {
	info := []byte(b.breakDown())
	err := os.WriteFile("bills/"+b.name+".txt", info, 0644)
	if err != nil {
		panic(err)
	}
	fmt.Println("bill saved")
}
