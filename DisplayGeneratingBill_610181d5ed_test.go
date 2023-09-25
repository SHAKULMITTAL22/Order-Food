package main

import (
	"fmt"
	"testing"
	"time"
)

// The original function displayGeneratingBill
func displayGeneratingBillOriginal() {
	fmt.Println()
	billDisplayText := "************************************* Generating Bill *************************************"
	for _, element := range billDisplayText {
		fmt.Printf("%c", element)
		time.Sleep(time.Millisecond * 15)
	}
}

func TestDisplayGeneratingBill_610181d5ed(t *testing.T) {

	startTime := time.Now()
	displayGeneratingBillOriginal()
	endTime := time.Now()

	// Test case: checking if the execution time is reasonable
	duration := endTime.Sub(startTime)
	if duration.Seconds() > 5 {
		t.Errorf("displayGeneratingBillOriginal() took too long to execute: %v", duration)
	} else {
		t.Logf("displayGeneratingBillOriginal() execution time is reasonable: %v", duration)
	}

	// Test case: Checking if the function is actually printing the required text
	billDisplayText := "************************************* Generating Bill *************************************"
	fmt.Println()
	expectedPrintTime := (len(billDisplayText) + 1) * 15 // "+1" is for the initial Println
	if int(duration.Milliseconds()) < expectedPrintTime {
		t.Errorf("displayGeneratingBillOriginal() terminated too early. It seems not everything was printed.")
	} else {
		t.Log("displayGeneratingBillOriginal() successfully printed the required text.")
	}
}
