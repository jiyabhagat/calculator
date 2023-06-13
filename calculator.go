package main

import (
	"fmt"
	"here/add"
	"here/div"
	multi "here/mul"
	"here/sub"
)

func main() {
	var num1, num2 float64
	var operator string
	var res float64 = 0.0

	fmt.Print("Enter the first number: ")
	fmt.Scanln(&num1)

	fmt.Print("Enter the second number: ")
	fmt.Scanln(&num2)

	fmt.Print("Enter the operator (+, -, *, /): ")
	fmt.Scanln(&operator)

	switch operator {
	case "+":
		res = add.Details(num1, num2)
		fmt.Println(num1, operator, num2, "=", res)
	case "/":
		res = div.Details(num1, num2)
		fmt.Println(num2, operator, num2, "=", res)
	case "*":
		res = multi.Details(num1, num2)
		fmt.Println(num1, operator, num2, "=", res)
	case "-":
		res = sub.Details(num1, num2)
		fmt.Println(num1, operator, num2, "=", res)
	}

}
