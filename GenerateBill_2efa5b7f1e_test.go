package main

import (
	"testing"
	"time"
)

func TestGenerateBill_2efa5b7f1e(t *testing.T){
	testCases := []struct {
		itemName string
		price    float64
		quantity int
		expectedTotalPrice float64
	}{
		{itemName: "TestItem1", price: 120.5, quantity: 2, expectedTotalPrice: 241},
		{itemName: "TestItem2", price: 150.5, quantity: 3, expectedTotalPrice: 451.5},
		// you can add more test cases here
	}

	for _, tc := range testCases {
		t.Run(tc.itemName, func(t *testing.T) {
			// calling the function generateBill with test case input
			actualTotalPrice := generateBill(tc.itemName, tc.price, tc.quantity)

			// comparing the expected output with the actual output
			if actualTotalPrice != tc.expectedTotalPrice {
				t.Errorf("Expected total price for %s to be %.2f but got %.2f", tc.itemName, tc.expectedTotalPrice, actualTotalPrice)
			}
		})
		time.Sleep(2 * time.Second)
	}
}
