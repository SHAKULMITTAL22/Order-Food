package main

import (
	"fmt"
	"math/rand"
	"strings"
	"testing"
	"time"
)

var subTotalBill float64

func generateBill() {
	//TODO: Implement bill generation logic
	subTotalBill = 1000.0
}

func printFinalBill() {
	for _, element := range "Here is your final bill:-" {
		fmt.Printf("%c", element)
		time.Sleep(time.Millisecond * 50)
	}
	fmt.Println()

	fmt.Printf("\n%52s\n", "JAIPUR BHOJANALYA")
	time.Sleep(time.Millisecond * 200)
	fmt.Printf("%s\n", strings.Repeat("*", 91))
	time.Sleep(time.Millisecond * 200)
	fmt.Printf("%86s\n", "Bhawani Singh Road, First Floor, Jaipur Bhojanalya, Jaipur, Jaipur 302005, Bharat")
	time.Sleep(time.Millisecond * 200)
	fmt.Printf("%50s\n", "Tel: 92145623XX")
	fmt.Printf("%60s\n\n", "Email: jaipur.bhojanalaya@gmail.com")
	time.Sleep(time.Millisecond * 200)
	fmt.Printf("%s", strings.Repeat("-", 42))
	fmt.Printf("%s", "INVOICE")
	fmt.Printf("%s\n", strings.Repeat("-", 42))
	time.Sleep(time.Millisecond * 200)

	rand.Seed(time.Now().Unix()) //necessary to produce random integers
	fmt.Printf(" Ticket No: %d\n", rand.Intn(550)+1)

	fmt.Printf(" Date: %v\n", time.Now().Local().Format("06-Jan-02")) //display date
	fmt.Printf(" Time: %v", time.Now().Local().Format("3:4 PM"))      //display time
	time.Sleep(time.Millisecond * 200)

	generateBill() //prints details of the bill,like, item name, price, quantity and total price and sub total amount.

	tax := 18 * subTotalBill / (100)
	grandTotal := subTotalBill + tax

	time.Sleep(time.Millisecond * 200)
	fmt.Printf("%71s: ₹%.2f\n", "GST", tax) //display tax amount
	fmt.Printf("+%s+\n", strings.Repeat("-", 90))
	time.Sleep(time.Millisecond * 200)
	fmt.Printf("%71s: ₹%.2f\n", "Grand Total", grandTotal) //display final bill
	fmt.Printf("+%s+\n", strings.Repeat("-", 90))

}

// TestPrintFinalBill_42388c9d9f tests the printFinalBill function.
func TestPrintFinalBill_42388c9d9f(t *testing.T) {
	t.Log("Given that the printFinalBill function is called")
	{
		// When the printFinalBill function is called
		printFinalBill()

		// Then the expected output should be displayed
		expectedOutput := `Here is your final bill:-

JAIPUR BHOJANALYA
****************************************************************************************************
Bhawani Singh Road, First Floor, Jaipur Bhojanalya, Jaipur, Jaipur 302005, Bharat
Tel: 92145623XX
Email: jaipur.bhojanalaya@gmail.com

------------------------------------------------INVOICE------------------------------------------------
 Ticket No: 345
 Date: 06-Jan-02
 Time: 3:4 PM
Item Name                                 Price  Quantity  Total
----------------------------------------------------------------------------------------------------
                                                100.00       1    100.00
                                                200.00       2    400.00
                                                300.00       3    900.00
                                                400.00       4   1600.00
                                                500.00       5   2500.00
----------------------------------------------------------------------------------------------------
Sub Total                                               5500.00
                                                
GST: ₹18.00
++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
Grand Total: ₹5518.00
++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
`
		if diff := compareOutput(expectedOutput, t); diff != "" {
			t.Errorf("Expected output:\n%s\nActual output:\n%s\nDiff:\n%s", expectedOutput, output, diff)
		} else {
			t.Log("Test passed")
		}
	}
}

func compareOutput(expectedOutput string, t *testing.T) string {
	t.Helper()
	output := fmt.Sprintf("%s", captureOutput(printFinalBill))
	diff := compare(expectedOutput, output)
	return diff
}

func compare(expected string, actual string) string {
	if expected == actual {
		return ""
	}
	var diff []string
	for i := 0; i < len(expected); i++ {
		if expected[i] != actual[i] {
			diff = append(diff, fmt.Sprintf("Expected '%c' at position %d, but got '%c'", expected[i], i, actual[i]))
		}
	}
	if len(diff) == 0 {
		return ""
	}
	return strings.Join(diff, "\n")
}

func captureOutput(f func()) string {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	f()
	log.SetOutput(os.Stderr)
	return buf.String()
}
