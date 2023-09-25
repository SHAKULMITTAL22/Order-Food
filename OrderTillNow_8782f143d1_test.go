package main

import (
	"fmt"
	"strings"
	"testing"
)

var customerOrder = []string{"Dosa", "Idly", "Vada"}

func orderTillNow() {
	//Print what you've ordered till now
	fmt.Println("\nYour order till now: ")
	fmt.Printf("%s\n", strings.Repeat("-", 32))
	fmt.Printf(" %-12s %s\n", "Quantity", "Item")
	fmt.Printf("%s\n", strings.Repeat("-", 32))

	for i := range customerOrder {
		fmt.Printf(" %-12v %v\n", customerOrder[i], i)
	}

	fmt.Printf("%s\n", strings.Repeat("-", 32))
}

// TestOrderTillNow_8782f143d1 tests the orderTillNow function.
func TestOrderTillNow_8782f143d1(t *testing.T) {
	t.Log("Test case: orderTillNow() with empty customerOrder")
	customerOrder = []string{}
	orderTillNow()
	if len(customerOrder) != 0 {
		t.Errorf("Expected customerOrder to be empty, got %v", customerOrder)
	}
	t.Log("Test case passed")
}

// TestOrderTillNow_8782f143d2 tests the orderTillNow function.
func TestOrderTillNow_8782f143d2(t *testing.T) {
	t.Log("Test case: orderTillNow() with non-empty customerOrder")
	customerOrder = []string{"Dosa", "Idly", "Vada"}
	orderTillNow()
	if len(customerOrder) != 3 {
		t.Errorf("Expected customerOrder to have 3 items, got %v", customerOrder)
	}
	t.Log("Test case passed")
}
