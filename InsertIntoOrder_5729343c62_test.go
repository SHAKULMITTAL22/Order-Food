package main

import (
	"testing"
)

// Mock orderItems function
var mockOrderItems = func() {}

// Test method for insertIntoOrder function
func TestInsertIntoOrder(t *testing.T) {

	// Test case where mockOrderItems function is successful
	{
		mockOrderItems = func() {}

		t.Run("Success", func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					t.Errorf("The code panicked: %v", r)
				}
			}()

			mockInsertIntoOrder()

			t.Log("Test passed successfully")
		})
	}

	// Test case where mockOrderItems function panics
	{
		mockOrderItems = func() {
			panic("Error while ordering items")
		}

		t.Run("Failure", func(t *testing.T) {
			defer func() {
				if r := recover(); r == nil {
					t.Error("The code did not panic")
				} else {
					t.Logf("The code panicked as expected: %v", r)
				}
			}()

			mockInsertIntoOrder()

			t.Log("Test passed successfully")
		})
	}
}

// mockInsertIntoOrder function
func mockInsertIntoOrder() {
	//add an item in the order
	mockOrderItems()
}
