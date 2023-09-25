package main

import (
	"fmt"
	"testing"
)

func sayTata(customerName string) {
	fmt.Println()
	fmt.Printf("%17s", " ")
	fmt.Printf("_/\\_ Thank you %v for visiting Jaipur Bhojanalya! _/\\_\n", customerName)
	fmt.Printf("%55s\n", "We hope to see you again!")
	fmt.Printf("%51s\n", "Have a nice day! ")
}

// TestSayTata_9d58da9153 is a test function for sayTata function.
// It checks if the function prints the correct output for a given customer name.
func TestSayTata_9d58da9153(t *testing.T) {
	// Test case 1: Valid customer name
	customerName := "John Doe"
	expectedOutput := `

               _/\_ Thank you John Doe for visiting Jaipur Bhojanalya! _/\_
We hope to see you again!
Have a nice day! `
	sayTata(customerName)
	if sayTata(customerName) != expectedOutput {
		t.Errorf("Expected output:\n%s\nActual output:\n%s", expectedOutput, sayTata(customerName))
	} else {
		t.Log("Test case 1 passed")
	}

	// Test case 2: Empty customer name
	customerName = ""
	expectedOutput = `

               _/\_ Thank you  for visiting Jaipur Bhojanalya! _/\_
We hope to see you again!
Have a nice day! `
	sayTata(customerName)
	if sayTata(customerName) != expectedOutput {
		t.Errorf("Expected output:\n%s\nActual output:\n%s", expectedOutput, sayTata(customerName))
	} else {
		t.Log("Test case 2 passed")
	}

	// Test case 3: Customer name with special characters
	customerName = "J@hn D*e"
	expectedOutput = `

               _/\_ Thank you J@hn D*e for visiting Jaipur Bhojanalya! _/\_
We hope to see you again!
Have a nice day! `
	sayTata(customerName)
	if sayTata(customerName) != expectedOutput {
		t.Errorf("Expected output:\n%s\nActual output:\n%s", expectedOutput, sayTata(customerName))
	} else {
		t.Log("Test case 3 passed")
	}
}
