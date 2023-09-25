package main

import (
	"fmt"
	"strings"
	"testing"
)

// The targeted function "orderTillNow"

func orderTillNowTest() {
	fmt.Println("\nYour order till now: ")
	fmt.Printf("%s\n", strings.Repeat("-", 32))
	fmt.Printf(" %-12s %s\n", "Quantity", "Item")
	fmt.Printf("%s\n", strings.Repeat("-", 32))

	for i := range customerOrderTest {
		fmt.Printf(" %-12v %v\n", customerOrderTest[i], i)
	}

	fmt.Printf("%s\n", strings.Repeat("-", 32))
}

//TestMethodName to validate functionality of orderTillNow
func TestOrderTillNow(t *testing.T) {
    var customerOrderTest = make(map[string]int)

	//Testing orderTillNow function when order is not placed
	customerOrderTest = make(map[string]int)
	orderTillNowTest()

	if len(customerOrderTest) != 0 {
		t.Error("Expected order quantity to be 0, got ", len(customerOrderTest))
	} else {
		t.Log("Test passed when order was not placed")
	}

	//Testing orderTillNow function after placing an order
	customerOrderTest = map[string]int{
		"Item1": 2,
		"Item2": 3,
	}

	orderTillNowTest()

	if len(customerOrderTest) == 0 {
		t.Error("Expected order quantity to be more than 0, got ", len(customerOrderTest))
	} else {
		t.Log("Test passed when order was placed")
	}

	// TODO: Add more test cases here for other scenarios if required
}
