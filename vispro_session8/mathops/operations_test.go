package mathops

import (
	"testing"
)

func TestAddition(t *testing.T) {
	result := Addition(3, 5)
	if result != 8 {
		t.Errorf("Expected 8, got %v", result)
	}
}

func TestSubtraction(t *testing.T) {
	result := Subtraction(7, 4)
	if result != 3 {
		t.Errorf("Expected 3, got %v", result)
	}
}

func TestMultiplication(t *testing.T) {
	result := Multiplication(2, 6)
	if result != 12 {
		t.Errorf("Expected 12, got %v", result)
	}
}

func TestDivision(t *testing.T) {
	result, err := Division(8, 2)
	if err != nil || result != 4 {
		t.Errorf("Expected result: 4, got result: %v, error: %v", result, err)
	}

	// Test division by zero
	_, err = Division(10, 0)
	if err == nil {
		t.Error("Expected error for division by zero, but got nil")
	}
}
