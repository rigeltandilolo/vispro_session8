package ui

import (
	"fmt"
	"visprosession8/vispro_session8/calculator" // Sesuaikan dengan struktur direktori dan nama proyek Anda
)

// ShowMenu menampilkan menu pilihan operasi kalkulator
func ShowMenu() {
	fmt.Println("Calculator Operations:")
	fmt.Println("1. Addition (+)")
	fmt.Println("2. Subtraction (-)")
	fmt.Println("3. Multiplication (*)")
	fmt.Println("4. Division (/)")
	fmt.Println("5. Power (^)")
	fmt.Println("6. Square Root (√)")
}

// GetUserInput menerima input dari pengguna untuk dua bilangan dan operasi
func GetUserInput() (float64, float64, string, error) {
	var num1, num2 float64
	var op string

	// Input pertama
	fmt.Print("Enter the first number: ")
	fmt.Scanln(&num1)

	// Pilihan operasi
	fmt.Print("Choose operation (1-5): ")
	fmt.Scanln(&op)

	// Input kedua
	fmt.Print("Enter the second number: ")
	fmt.Scanln(&num2)

	// Validasi input numerik
	if isNaN(num1) || isNaN(num2) {
		return 0, 0, "", fmt.Errorf("invalid input: please enter numeric values")
	}

	return num1, num2, op, nil
}

// isNaN memeriksa apakah suatu nilai adalah NaN (Not a Number)
func isNaN(num float64) bool {
	return num != num
}

// RunCalculator menjalankan kalkulator berdasarkan input pengguna
func RunCalculator() {
	ShowMenu()

	num1, num2, op, err := GetUserInput()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	var result float64
	switch op {
	case "1":
		result = calculator.Add(num1, num2)
	case "2":
		result = calculator.Subtract(num1, num2)
	case "3":
		result = calculator.Multiply(num1, num2)
	case "4":
		result, err = calculator.Divide(num1, num2)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
	case "5":
		result = calculator.Power(num1, num2)
	case "6":
		result, err = calculator.SquareRoot(num1)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
	default:
		fmt.Println("Invalid operation")
		return
	}

	fmt.Printf("Result: %v\n", result)
}

func main() {
	RunCalculator()
}
