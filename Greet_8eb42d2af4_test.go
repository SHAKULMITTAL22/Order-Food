package main

import (
	"fmt"
	"testing"
)

func greet(customerName string) {
	fmt.Printf("%52s %s\n", "Namaste ", customerName)
	fmt.Printf("%72s\n", "_/\\_ Welcome to Jaipur Bhojanalya! _/\\_ ")
	fmt.Println()
}

// TestGreet_8eb42d2af4 is a test function for greet function.
// It checks if the function prints the correct output for a given customer name.
func TestGreet_8eb42d2af4(t *testing.T) {
	// Test case 1: Valid customer name
	customerName := "John Doe"
	expectedOutput := `                                 Namaste John Doe
________________________________________________________________________________
                               _/\_ Welcome to Jaipur Bhojanalya! _/\_ 
`
	greet(customerName)
	if gotOutput := fmt.Sprintf("%s", greet(customerName)); gotOutput != expectedOutput {
		t.Errorf("greet(%q) = %q, want %q", customerName, gotOutput, expectedOutput)
	} else {
		t.Log("Test case 1 passed: greet() prints the correct output for a valid customer name.")
	}

	// Test case 2: Empty customer name
	customerName = ""
	expectedOutput = `                                 Namaste 
________________________________________________________________________________
                               _/\_ Welcome to Jaipur Bhojanalya! _/\\_ 
`
	greet(customerName)
	if gotOutput := fmt.Sprintf("%s", greet(customerName)); gotOutput != expectedOutput {
		t.Errorf("greet(%q) = %q, want %q", customerName, gotOutput, expectedOutput)
	} else {
		t.Log("Test case 2 passed: greet() prints the correct output for an empty customer name.")
	}

	// Test case 3: Customer name with special characters
	customerName = "J@hn D*e"
	expectedOutput = `                                 Namaste J@hn D*e
________________________________________________________________________________
                               _/\_ Welcome to Jaipur Bhojanalya! _/\\_ 
`
	greet(customerName)
	if gotOutput := fmt.Sprintf("%s", greet(customerName)); gotOutput != expectedOutput {
		t.Errorf("greet(%q) = %q, want %q", customerName, gotOutput, expectedOutput)
	} else {
		t.Log("Test case 3 passed: greet() prints the correct output for a customer name with special characters.")
	}
}
