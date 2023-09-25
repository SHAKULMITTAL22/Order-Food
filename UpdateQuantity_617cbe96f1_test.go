package main

import (
	"fmt"
	"testing"
)

// Dummy structure to simulate the test
type Item struct {
	itemNo    uint
	itemName  string
	itemPrice uint
}

// Dummy variables to simulate the test
var testMenu []Item
var testCustomerOrder map[string]uint
var testSubTotalBill float64

// This is just a test function to simulate 'delFromOrder'. Actual implementation may vary
func testDelFromOrder(serialNo uint) {
	// Simulating deletion of an item by setting its quantity to zero in testCustomerOrder map
	for _, targetItem := range testMenu {
		if serialNo == targetItem.itemNo {
			testCustomerOrder[targetItem.itemName] = 0
			return
		}
	}
}

// TestUpdateQuantity_617cbe96f1 is a function that will handle the test cases for the 'updateQuantity' function
func TestUpdateQuantity_617cbe96f1(t *testing.T) {

	testMenu = []Item{
		{1, "apple", 10},
		{2, "banana", 5},
		{3, "grapes", 15},
	}
	testCustomerOrder = map[string]uint{
		"apple":  5,
		"banana": 2,
		"grapes": 1,
	}
	testSubTotalBill = 70

	testUpdateQuantity(1, 7)
	if testCustomerOrder["apple"] != 7 || testSubTotalBill != 80 { // 2 new apples added => 20 should be added to bill
		t.Errorf("Fail: Expected quantity: 7, actual quantity: %v, Expected bill: 80, actual bill: %v", testCustomerOrder["apple"], testSubTotalBill)
	} else {
		t.Logf("Success: Set quantity test - Quantity: %v, Bill: %v", testCustomerOrder["apple"], testSubTotalBill)
	}

	testUpdateQuantity(2, 0)
	if testCustomerOrder["banana"] != 0 || testSubTotalBill != 70 { // 2 bananas removed => 10 should be subtracted from bill
		t.Errorf("Fail: Expected quantity: 0, actual quantity: %v, Expected bill: 70, actual bill: %v", testCustomerOrder["banana"], testSubTotalBill)
	} else {
		t.Logf("Success: Delete item test - Quantity: %v, Bill: %v", testCustomerOrder["banana"], testSubTotalBill)
	}

	testUpdateQuantity(4, 1) // serialNo 4 doesn't exist. So, no changes expected
	if testSubTotalBill != 70 {
		t.Errorf("Fail: No action expected. But the bill was modified to %v", testSubTotalBill)
	} else {
		t.Logf("Success: No action test - Bill: %v", testSubTotalBill)
	}
}

func testUpdateQuantity(serialNo uint, newQuantity uint) {
	for _, targetItem := range testMenu {
		if serialNo == targetItem.itemNo {
			itemName := targetItem.itemName        //name of the item whose quantity to be updated
			oldQuantity := testCustomerOrder[itemName] //current quantity of that item
			fmt.Printf("Current quantity of %v is %v.\n", itemName, oldQuantity)
			fmt.Printf("Now, set the updated quantity of %v to be ordered: %v\n", itemName, newQuantity)
			//in case new quantity is '0' then delete that item from the order
			if newQuantity == 0 {
				testDelFromOrder(serialNo)
				return
			}
			fmt.Printf("")
			//update quantity of item
			testCustomerOrder[targetItem.itemName] = newQuantity
			fmt.Printf("Updated the quantity of %v from %v to %v.\n", itemName, oldQuantity, newQuantity)
			//update bill
			testSubTotalBill -= float64(oldQuantity) * float64(targetItem.itemPrice) //delete the item price for old quantity
			testSubTotalBill += float64(newQuantity) * float64(targetItem.itemPrice) //add the item price for updated quantity
			break
		}
	}
}
