package main

import "fmt"

func main() {
	var numberOne, numberTwo float32
	var operator string

	fmt.Print("Please provide the 1st number \n")
	fmt.Scanln(&numberOne)

	fmt.Print("Please provide the 2nd number \n")
	fmt.Scanln(&numberTwo)

	fmt.Print("Please provide the operator (+, -, *, /) \n")
	fmt.Scanln(&operator)

	switch operator {
	case "+":
		fmt.Printf("%f %s %f = %f", numberOne, operator, numberTwo, numberOne+numberTwo)
	case "-":
		fmt.Printf("%f %s %f = %f", numberOne, operator, numberTwo, numberOne-numberTwo)
	case "*":
		fmt.Printf("%f %s %f = %f", numberOne, operator, numberTwo, numberOne*numberTwo)
	case "/":
		//check if numberOne is not 0
		if numberOne == 0.0 {
			fmt.Println("Cannot divide by zero")
		} else {
			fmt.Printf("%f %s %f = %f", numberOne, operator, numberTwo, numberOne/numberTwo)
		}
	default:
		fmt.Println("Invalid operator")

	}
}
