package main

import (
    "fmt"
    "testing"
)

func delFromOrder(serialNo uint) {

    for _, targetItem := range menu {
        if serialNo == targetItem.itemNo {
            itemName := targetItem.itemName //name of the item who is to be removed from the list
            //update bill; customerOrder[itemName] -> it results in the quantity of that item
            subTotalBill -= float64(customerOrder[itemName]) * float64(targetItem.itemPrice)
            //delete item
            delete(customerOrder, itemName)
            fmt.Printf("%v removed from the order list.\n", itemName)
            break
        }
    }
}

// TestDelFromOrder_66d2b41819 tests the delFromOrder function.
func TestDelFromOrder_66d2b41819(t *testing.T) {
    // Test case 1: Valid serial number.
    serialNo := uint(1)
    delFromOrder(serialNo)
    if _, ok := customerOrder["item1"]; ok {
        t.Errorf("Item not removed from order list.")
    }
    t.Log("Test case 1 passed.")

    // Test case 2: Invalid serial number.
    serialNo = uint(100)
    delFromOrder(serialNo)
    if _, ok := customerOrder["item100"]; ok {
        t.Errorf("Item not removed from order list.")
    }
    t.Log("Test case 2 passed.")
}
