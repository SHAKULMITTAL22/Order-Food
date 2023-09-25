package main

import (
	"fmt"
	"strings"
	"testing"
)

type item struct {
	itemNo     int
	itemName   string
	itemPrice  float64
}

var menu = []item{
	{1, "Dosa", 12.99},
	{2, "Idly", 2.99},
	{3, "Vada", 1.99},
	{4, "Uttapam", 8.99},
	{5, "Puri Bhaji", 10.99},
}

func printMenu() {
	fmt.Printf("%15s\n", "Menu")
	fmt.Printf("%s\n", strings.Repeat("-", 35))
	fmt.Printf("%-7s %6s    %12s\n", "S.No.", "Price", "Item Name")
	fmt.Printf("%s\n", strings.Repeat("-", 35))

	for _, element := range menu {
		fmt.Printf(" %-7d %.2f    %-4s\n", element.itemNo, element.itemPrice, element.itemName)
	}
	fmt.Printf("%s", strings.Repeat("-", 35))
	fmt.Println()
}

func TestPrintMenu_8422b57a1c(t *testing.T) {
	t.Log("Test case: Print Menu")
	expectedOutput := `Menu
-----------------------------------
S.No.  Price    Item Name
-----------------------------------
    1  12.99    Dosa
    2   2.99    Idly
    3   1.99    Vada
    4   8.99    Uttapam
    5  10.99    Puri Bhaji
-----------------------------------
`
	printMenu()
	if output := captureOutput(printMenu); output != expectedOutput {
		t.Errorf("Expected output:\n%s\nActual output:\n%s", expectedOutput, output)
	} else {
		t.Log("Test case passed")
	}
}

func captureOutput(f func()) string {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	f()
	log.SetOutput(os.Stdout)
	return buf.String()
}
