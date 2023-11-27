package mathops

import "errors"

// Addition returns the sum of two numbers
func Addition(a, b float64) float64 {
	return a + b
}

// Subtraction returns the result of subtracting b from a
func Subtraction(a, b float64) float64 {
	return a - b
}

// Multiplication returns the product of two numbers
func Multiplication(a, b float64) float64 {
	return a * b
}

// Division returns the result of dividing a by b
func Division(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero is not allowed")
	}
	return a / b, nil
}
