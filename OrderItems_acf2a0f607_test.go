package main

import "fmt"

var customerOrder = make(map[string]uint)
var subTotalBill float64
var menu = []struct{ itemName string; itemPrice float64 }{}

func printMenu() {
	// to be implemented
}

func orderTillNow() {
	// to be implemented
}

func orderItemsForTest() {
	// Print the menu to the user
	printMenu()
	var itemNumber uint
	var noOfPlates uint

	// Loop until user decides to break (by inputting 0)
	for {
		fmt.Println()
		fmt.Println("Enter '0' to exit.")

		// Ask user for input
		fmt.Print("Enter the serial no. of the item to order: ")

		// Scan input value
		fmt.Scan(&itemNumber)
		// Break loop if user input was 0
		if itemNumber == 0 {
			break
		}

		var choiceName string
		var itemPrice float64
		// Find chosen item name and price from the menu
		for index, item := range menu {
			if index+1 == int(itemNumber) {
				choiceName = item.itemName
				itemPrice = item.itemPrice
				break
			}
		}
		// Ask user for a number of items
		fmt.Printf("How many %v do you want: ", choiceName)
		// Scan input value
		fmt.Scan(&noOfPlates)

		// If user asked for 0 items, continue the loop
		if noOfPlates == 0 {
			continue
		} else {
			// Update order and bill for the current item
			for index /*, item*/ := range menu {
				// Add chosen item to the order
				if index+1 == int(itemNumber) {
					customerOrder[choiceName] += noOfPlates
					// Calculate price and update the bill
					subTotalBill += itemPrice * float64(noOfPlates)
					break
				}
			}

			// Print a message about the current order
			fmt.Printf("\nYou just ordered %v %v which amounts to ₹%v.\n", noOfPlates, choiceName, itemPrice*float64(noOfPlates))
			// Print the order so far
			orderTillNow()
		}
		fmt.Println()
	}
}
