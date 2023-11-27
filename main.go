package main

import (
	"fmt"
	"visprosession8/ui"
)

func main() {
	fmt.Println("Welcome to the Rigel's Calculator!")

	for {
		ui.RunCalculator()

		var choice string
		fmt.Print("Do you want to perform another calculation? (yes/no): ")
		fmt.Scanln(&choice)

		if choice != "yes" {
			fmt.Println("Thank you for using the calculator. Goodbye! <3")
			break
		}
	}
}
