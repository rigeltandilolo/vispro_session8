package calculator

import (
	"visprosession8/vispro_session8/mathops"
)

// Add performs addition using the mathops package
func Add(a, b float64) float64 {
	return mathops.Addition(a, b)
}

// Subtract performs subtraction using the mathops package
func Subtract(a, b float64) float64 {
	return mathops.Subtraction(a, b)
}

// Multiply performs multiplication using the mathops package
func Multiply(a, b float64) float64 {
	return mathops.Multiplication(a, b)
}

// Divide performs division using the mathops package
func Divide(a, b float64) (float64, error) {
	return mathops.Division(a, b)
}

// Power calculates the power of a number using the mathops package
func Power(base, exponent float64) float64 {
	result := 1.0
	for i := 0; i < int(exponent); i++ {
		result *= base
	}
	return result
}
