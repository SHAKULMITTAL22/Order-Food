package main

import (
	"testing"
	"bytes"
	"fmt"
	"os"
	"io"
	"strings"
)

type MenuItem struct {
	itemNo    int
	itemPrice float64
	itemName  string
}

var MenuItems = []MenuItem{
	{1, 15.99, "Pizza"},
	{2, 2.99, "Soda"},
}

func printMenuItems() {
	fmt.Printf("%15s\n", "Menu")
	fmt.Printf("%s\n", strings.Repeat("-", 35))
	fmt.Printf("%-7s %6s    %12s\n", "S.No.", "Price", "Item Name")
	fmt.Printf("%s\n", strings.Repeat("-", 35))
	for _, item := range MenuItems {
		fmt.Printf(" %-7d %.2f    %-4s\n", item.itemNo, item.itemPrice, item.itemName)
	}

	fmt.Printf("%s", strings.Repeat("-", 35))
	fmt.Println()
}

func TestPrintMenu(t *testing.T) {
	// Capture the standard output
	old := os.Stdout 
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Call the function
	printMenuItems()

	// Restore the standard output
	w.Close()
	out, _ := io.ReadAll(r)
	os.Stdout = old

	// Check the output
	expected := "           Menu\n-----------------------------------\nS.No.   Price    Item Name\n-----------------------------------\n 1       15.99    Pizza\n 2       2.99    Soda\n-----------------------------------\n"
	if string(out) != expected {
		t.Errorf("PrintMenu() = %v, want %v", string(out), expected)
	}
}

