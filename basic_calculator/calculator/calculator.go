package calculator

import "fmt"

func Calculate(a int, b int, operator string) int {
	var res int
	switch operator {
	case "+":
		res = a + b
	case "-":
		res = a - b
	case "*":
		res = a * b
	case "/":
		res = a / b
	default:
		fmt.Println("Invalid operator")
	}
	
	return res
}