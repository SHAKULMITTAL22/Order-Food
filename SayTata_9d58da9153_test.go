package main

import (
	"fmt"
	"testing"
)

func sayTata(customerName string) string {
	return fmt.Sprintf("\n                 _/\\_ Thank you %v for visiting Jaipur Bhojanalya! _/\\_\n                                     We hope to see you again!\n                                                Have a nice day! ", customerName)
}

func TestSayTata_9d58da9153(t *testing.T) {
	testCases := []struct {
		customerName   string
		expectedOutput string
	}{
		{
			customerName:   "John Doe",
			expectedOutput: "\n                 _/\\_ Thank you John Doe for visiting Jaipur Bhojanalya! _/\\_\n                                     We hope to see you again!\n                                                Have a nice day! ",
		},
		{
			customerName:   "",
			expectedOutput: "\n                 _/\\_ Thank you  for visiting Jaipur Bhojanalya! _/\\_\n                                     We hope to see you again!\n                                                Have a nice day! ",
		},
	}

	for _, tc := range testCases {
		result := sayTata(tc.customerName)
		if result != tc.expectedOutput {
			t.Errorf("sayTata(%s) = %s; expected %s", tc.customerName, result, tc.expectedOutput)
		} else {
			t.Logf("TestSayTata_9d58da9153 on %s PASSED", tc.customerName)
		}
	}
}
