package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// fn returns 2 things (string, error)
func getUserInput(prompt string, read *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	userInput, err := read.ReadString('\n')

	return strings.TrimSpace(userInput), err

}

func initiateBill() Bill {
	reader := bufio.NewReader(os.Stdin)

	name, _ := getUserInput("Create a new bill: ", reader)

	bill := newBill(name)
	fmt.Println(bill.name, "bill created")

	return bill
}

func billOptions(createdBill Bill) {
	reader := bufio.NewReader(os.Stdin)
	options, _ := getUserInput("type a - to add an item \n s - to save \n t - to add tip: \n", reader)

	switch options {
	case "a":
		//add item
		itemName, _ := getUserInput("Item name: ", reader)
		itemPrice, _ := getUserInput("Price: ", reader)

		price, err := strconv.ParseFloat(itemPrice, 64)

		if err != nil {
			fmt.Println("Price must be a number")
			billOptions(createdBill)
		}
		createdBill.addItem(itemName, price)
		fmt.Println(itemName, "and", price, "has been added")
		billOptions(createdBill)
	case "s":
		//save item
		createdBill.save()
		fmt.Println(createdBill.name, "Stored in a text file")

	case "t":
		//add tip
		tip, _ := getUserInput("Add a tip: ", reader)
		t, err := strconv.ParseFloat(tip, 64)
		if err != nil {
			fmt.Println("tip must be a number")
			billOptions(createdBill)
		}
		createdBill.updateTip(t)
		billOptions(createdBill)
	default:
		fmt.Println("invalid option")
		billOptions(createdBill)
	}
}

func main() {
	newBill := initiateBill()
	billOptions(newBill)
}
