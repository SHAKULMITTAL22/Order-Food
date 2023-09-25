package main

import (
	"fmt"
	"math/rand"
	"strings"
	"testing"
	"time"
)

var (
	menu = []MenuItem{
		{"Dosa", 100.0},
		{"Idly", 50.0},
		{"Vada", 20.0},
		{"Uttapam", 150.0},
		{"Poori", 120.0},
	}

	customerOrder = map[string]int{
		"Dosa":   2,
		"Idly":   3,
		"Vada":   1,
		"Uttapam": 1,
		"Poori":  2,
	}
)

func printOrderData() {
	for key := range customerOrder {
		//key -> it is the key values
		for _, element := range menu {
			if key == element.itemName {
				//customerOrder[key] -> will provide the no. of plates of that item
				//float64(customerOrder[key])*element.itemPrice -> this will result in the cost of each item
				totalCostOfItem := float64(customerOrder[key]) * element.itemPrice
				fmt.Printf(" %-30s %-20.2f %-20v %-20.2f\n", key, element.itemPrice, customerOrder[key], totalCostOfItem)
			}
		}
	}
	fmt.Println()
}

// TestPrintOrderData_47d77efb9c tests the printOrderData function.
func TestPrintOrderData_47d77efb9c(t *testing.T) {
	t.Log("Given customer order and menu data")
	{
		expectedOutput := `Item Name                 Item Price  Quantity  Total Cost
Dosa                       100.00       2         200.00
Idly                        50.00       3         150.00
Vada                        20.00       1          20.00
Uttapam                    150.00       1         150.00
Poori                      120.00       2         240.00

`
		printOrderData()
		if diff := compareOutput(expectedOutput, t); diff != "" {
			t.Errorf("printOrderData() failed.\nExpected:\n%s\nActual:\n%s\nDiff:\n%s", expectedOutput, output, diff)
		} else {
			t.Log("printOrderData() passed.")
		}
	}
}

func compareOutput(expected, t *testing.T) string {
	t.Helper()
	expectedLines := strings.Split(expected, "\n")
	actualLines := strings.Split(output, "\n")
	if len(expectedLines) != len(actualLines) {
		return fmt.Sprintf("Expected %d lines, got %d", len(expectedLines), len(actualLines))
	}
	for i := range expectedLines {
		if expectedLines[i] != actualLines[i] {
			return fmt.Sprintf("Line %d: Expected %q, got %q", i+1, expectedLines[i], actualLines[i])
		}
	}
	return ""
}

var output string

func init() {
	rand.Seed(time.Now().UnixNano())
	buf := make([]byte, 1024)
	rand.Read(buf)
	output = fmt.Sprintf("%s", buf)
}
