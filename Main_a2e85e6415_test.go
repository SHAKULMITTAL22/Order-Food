package main

import "testing"

func greet(name string) string {
	return "Hello, " + name
}

func TestGreet(t *testing.T) {
	result := greet("Test")
	if result != "Hello, Test" {
		t.Errorf("greet('Test') = %s; want 'Hello, Test'", result)
	}
}
