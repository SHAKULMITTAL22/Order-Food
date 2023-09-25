package main

import (
    "fmt"
    "testing"
)

type args struct {
    serialNo uint
}

var tests = []struct {
    name string
    args args
}{
    {"UpdateQuantity_617cbe96f1", args{1}},
    {"UpdateQuantity_617cbe96f2", args{2}},
    {"UpdateQuantity_617cbe96f3", args{3}},
    {"UpdateQuantity_617cbe96f4", args{4}},
    {"UpdateQuantity_617cbe96f5", args{5}},
}

func TestUpdateQuantity(t *testing.T) {
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            updateQuantity(tt.args.serialNo)
            t.Log("Test Case Passed")
        })
    }
}

func updateQuantity(serialNo uint) {

	var newQuantity uint

	for _, targetItem := range menu {

		if serialNo == targetItem.itemNo {

			itemName := targetItem.itemName        //name of the item whose quantity to be updated
			oldQuantity := customerOrder[itemName] //current quantity of that item

			fmt.Printf("Current quantity of %v is %v.\n", itemName, oldQuantity)
			fmt.Printf("Now, enter the updated quantity of %v to be ordered: \n", itemName)
			fmt.Scan(&newQuantity)

			//in case new quantity is '0' then delete that item from the order
			if newQuantity == 0 {
				delFromOrder(serialNo)
				return
			}
			fmt.Printf("")

			//update quantity of item
			customerOrder[targetItem.itemName] = newQuantity
			fmt.Printf("Updated the quantity of %v from %v to %v.\n", itemName, oldQuantity, newQuantity)

			//update bill
			subTotalBill -= float64(oldQuantity) * float64(targetItem.itemPrice) //delete the item price for old quantity
			subTotalBill += float64(newQuantity) * float64(targetItem.itemPrice) //add the item price for updated quantity

			break
		}
	}

}
