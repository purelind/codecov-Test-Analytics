package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	result := Add(2, 3)
	if result != 5 {
		t.Errorf("Add(2, 3) = %d; want 5", result)
	}
}

func TestSubtract(t *testing.T) {
	result := Subtract(5, 3)
	if result != 2 {
		t.Errorf("Subtract(5, 3) = %d; want 2", result)
	}
}

func TestMultiply(t *testing.T) {
	result := Multiply(2, 3)
	if result != 6 {
		t.Errorf("Multiply(2, 3) = %d; want 6", result)
	}
}

func TestFlakyDivide(t *testing.T) {


	result, err := FlakyDivide(6, 2)
	if err != nil {
		t.Logf("FlakyDivide failed: %v", err)
		t.Error("This is flaky test")
	}
	if result != 3 {
		t.Errorf("FlakyDivide(6, 2) = %d; want 3", result)
	} 


	// Test division by zero
	_, err = FlakyDivide(6, 0)
	if err == nil {
		t.Error("FlakyDivide(6, 0) should return an error for division by zero")
	}

}