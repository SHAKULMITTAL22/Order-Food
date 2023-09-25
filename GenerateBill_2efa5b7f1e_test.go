package main

import (
	"fmt"
	"math/rand"
	"strings"
	"testing"
	"time"
)

var subTotalBill float64

// This function prints the details of the food item that you've ordered.
func printOrderData() {
	for _, item := range orderList {
		fmt.Printf(" %-30s %-20.2f %-20d %-20.2f\n", item.itemName, item.price, item.quantity, item.totalPrice)
	}
}

// This function prints the sub total amount in a cool way!
func printSubTotal() {
	fmt.Printf("%47s", " ")
	for _, element := range "Sub Total(excluding GST): ₹" {
		fmt.Printf("%c", element)
		time.Sleep(time.Millisecond * 50)
	}
	fmt.Printf("%.2f\n", subTotalBill)
}

// This function generates a bill.
func generateBill() {

	fmt.Println()
	fmt.Printf("+%s+\n", strings.Repeat("-", 90))
	fmt.Printf(" %-30s %-20s %-20s %-20s\n", "Item Name", "Price(₹)", "Quantity", "Total Price(₹)")
	fmt.Printf("+%s+\n", strings.Repeat("-", 90))

	//prints the details of the food item that you've ordered.
	printOrderData()

	fmt.Printf("+%s+\n", strings.Repeat("-", 90))

	//print sub total amount in a cool way!
	printSubTotal()

}

// TestGenerateBill_2efa5b7f1e tests the generateBill function.
func TestGenerateBill_2efa5b7f1e(t *testing.T) {
	// Create a test order.
	order := Order{
		orderList: []Item{
			{itemName: "Item 1", price: 10.0, quantity: 2},
			{itemName: "Item 2", price: 20.0, quantity: 3},
		},
	}

	// Calculate the expected sub total.
	expectedSubTotal := 0.0
	for _, item := range order.orderList {
		expectedSubTotal += item.price * float64(item.quantity)
	}

	// Generate the bill.
	generateBill()

	// Check that the sub total is correct.
	if subTotalBill != expectedSubTotal {
		t.Errorf("Expected sub total to be %.2f, but got %.2f", expectedSubTotal, subTotalBill)
	} else {
		t.Log("Sub total is correct")
	}
}
