package main

import (
	"fmt"
	"testing"
)

func TestMain_a2e85e6415(t *testing.T) {
	// Test case 1: Valid customer name
	customerName := "John"
	fmt.Println("What is your first name?")
	fmt.Scan(&customerName)

	greet(customerName)
	orderItems()
	displayGeneratingBill() //just displays that "generating bill" in a fancy manner.
	generateBill()
	modifyOrder()
	printFinalBill()
	sayTata(customerName)

	// Expected output:
	// What is your first name?
	// John
	// Hello John, welcome to our restaurant!
	// What would you like to order?
	// Generating bill...
	// Your bill is $100.
	// Would you like to modify your order? (y/n)
	// n
	// Thank you for dining with us, John!
	// Have a great day!

	// Test case 2: Invalid customer name
	customerName = ""
	fmt.Println("What is your first name?")
	fmt.Scan(&customerName)

	greet(customerName)
	orderItems()
	displayGeneratingBill() //just displays that "generating bill" in a fancy manner.
	generateBill()
	modifyOrder()
	printFinalBill()
	sayTata(customerName)

	// Expected output:
	// What is your first name?

	// Error: customer name cannot be empty
}
