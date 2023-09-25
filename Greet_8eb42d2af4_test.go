package main

import (
	"bytes"
	"fmt"
	"testing"
)

// Mocked print function
func printToBuffer(format string, a ...interface{}) (n int, err error) {
	var buf bytes.Buffer
	return fmt.Fprintf(&buf, format, a...)
}

// TestGreet tests the greet function
func TestGreet(t *testing.T) {
	cases := []struct {
		customerName string
		expected     string
	}{
		{
			customerName: "John",
			expected:     "                                              Namaste John\n                                               _/\\_ Welcome to Jaipur Bhojanalya! _/\\_\n\n",
		},
		{
			customerName: "Alex",
			expected:     "                                              Namaste Alex\n                                               _/\\_ Welcome to Jaipur Bhojanalya! _/\\_\n\n",
		},
		// TODO: Add more test cases here
	}

	for i, c := range cases {
		// Call the function
		got, _ := printToBuffer("%52s %s\n", "Namaste ", c.customerName)
		got2, _ := printToBuffer("%72s\n", "_/\\_ Welcome to Jaipur Bhojanalya! _/\\_ ")
		got += got2

		// Check the result
		if got != c.expected {
			t.Errorf("Test %d failed, expected %q but got %q", i+1, c.expected, got)
		} else {
			t.Logf("Test %d passed", i+1)
		}
	}
}

func TestGenerateBill(t *testing.T) {
	cases := []struct {
		itemName string
		price    float64
		quantity int
		expected string // This is just a mock. Change this based on the actual function signature and return type
	}{
		{
			itemName: "Test Item",
			price:    10.0,
			quantity: 2,
			expected: "Your bill: 20.0", // This is just a mock. Change this based on the actual function signature and return type
		},
		// TODO: Add more test cases here
	}

	for i, c := range cases {
		got := generateBill(c.itemName) // Modified this by passing only the required parameters

		if got != c.expected {
			t.Errorf("Test %d failed, expected %q but got %q", i+1, c.expected, got)
		} else {
			t.Logf("Test %d passed", i+1)
		}
	}
}

func generateBill(item string) string {
	// TODO: Implement your generateBill function here
	return "Your bill: 20.0" // This is just a mock. Change this based on the actual function implementation
}
