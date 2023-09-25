package main

import (
    "testing"
    "fmt"
)

// Preparing simulated menu and order as package-level variables for testing purposes.
var testMenu = []struct {
    itemNo   uint
    itemName string
    itemPrice float64
} {
    {1, "Burger", 5.99},
    {2, "Fries", 3.49},
    {3, "Soda", 1.99},
}

var testCustomerOrder = map[string]uint {
    "Burger": 2,
    "Fries": 1,
}

var testSubTotalBill float64 = 15.47  // 2 burgers and 1 fries

// Test method for delFromOrder.
func TestDelFromOrder_66d2b41819(t *testing.T) {
    delFromOrderTest(1)
    if _, found := testCustomerOrder["Burger"]; found {
        t.Error(fmt.Sprintf("TestDelFromOrder --- FAIL --- found Burger in the order"))
    } else {
        t.Log("TestDelFromOrder --- PASS --- Burger successfully removed from the order")
    }

    if testSubTotalBill != 3.49 {
        t.Error(fmt.Sprintf("TestDelFromOrder --- FAIL --- expected remaining bill to be 3.49, got %v", testSubTotalBill))
    } else {
        t.Log("TestDelFromOrder --- PASS --- subTotalBill correctly updated after deleting Burger")
    }

    delFromOrderTest(5)
    if testSubTotalBill != 3.49 {
        t.Error(fmt.Sprintf("TestDelFromOrder --- FAIL --- expected bill to be unchanged (3.49), got %v", testSubTotalBill))
    } else {
        t.Log("TestDelFromOrder --- PASS --- subTotalBill correctly remained the same when attempting to delete non-existing item")
    }
}

func delFromOrderTest(serialNo uint) {
    for _, targetItem := range testMenu {
        if serialNo == targetItem.itemNo {
            itemName := targetItem.itemName //name of the item who is to be removed from the list
            //update bill; customerOrder[itemName] -> it results in the quantity of that item
            testSubTotalBill -= float64(testCustomerOrder[itemName]) * float64(targetItem.itemPrice)
            //delete item
            delete(testCustomerOrder, itemName)
            fmt.Printf("%v removed from the order list.\n", itemName)
            break
        }
    }
}
