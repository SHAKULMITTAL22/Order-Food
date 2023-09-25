package main

import (
	"fmt"
	"strings"
	"testing"
	"time"
)

func displayGeneratingBill() {
	fmt.Println()
	billDisplayText := "************************************* Generating Bill *************************************"
	for _, element := range billDisplayText {
		fmt.Printf("%c", element) // if you use "%v" instead of "%c" then convert element into string, as shown in the comment below
		// fmt.Print("%v", string(element))
		time.Sleep(time.Millisecond * 15)
	}
}

// TestDisplayGeneratingBill_610181d5ed tests the displayGeneratingBill function.
func TestDisplayGeneratingBill_610181d5ed(t *testing.T) {
	// Test case 1: Check if the function prints the correct message.
	expectedMessage := "************************************* Generating Bill *************************************"
	displayGeneratingBill()
	if !strings.Contains(expectedMessage, billDisplayText) {
		t.Errorf("Expected message '%s', got '%s'", expectedMessage, billDisplayText)
	} else {
		t.Log("Test case 1 passed: Function prints the correct message.")
	}

	// Test case 2: Check if the function sleeps for the correct duration.
	expectedDuration := time.Millisecond * 15
	startTime := time.Now()
	displayGeneratingBill()
	endTime := time.Now()
	if endTime.Sub(startTime) != expectedDuration {
		t.Errorf("Expected duration '%s', got '%s'", expectedDuration, endTime.Sub(startTime))
	} else {
		t.Log("Test case 2 passed: Function sleeps for the correct duration.")
	}
}
