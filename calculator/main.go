package main

import "fmt"

func main() {
	var operation string
	var num1, num2 int

	fmt.Println("Calculago 1.0")
	fmt.Println("==========================")
	fmt.Println("Which operation do you want to perform? (add, substract, multiply, divide)")
	fmt.Scanf("%s", &operation)
	fmt.Println("Enter first number")
	fmt.Scanf("%d", &num1)
	fmt.Println("Enter second number")
	fmt.Scanf("%d", &num2)

	switch operation {
	case "add":
		fmt.Println("ADD + ADD2")
	case "substract":
		fmt.Println(num1 - num2)
	case "multiply":
		fmt.Println(num1 * num2)
	case "divide":
		fmt.Println(num1 / num2)
	}

}
