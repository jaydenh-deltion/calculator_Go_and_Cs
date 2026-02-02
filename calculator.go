package main 

import ("fmt")

func main(){
	fmt.Print("welcome to my calculator\n") // print function 
	var operator string // variable declaration
	fmt.Print("Enter operator: +, -, *, /:\n ")
	fmt.Scan(&operator) // input for operator 

	var num1 , num2 float64
	fmt.Print("Enter first number: ")
	fmt.Scan(&num1) // input num1

	fmt.Print("Enter second number: ")
	fmt.Scan(&num2) //input num2

	if operator == "/" && num2 == 0 {
		fmt.Print("Error! Dividing by zero is not allowed.")
	}
	switch operator{
	case "+":
		fmt.Println(num1 + num2)
	case "-":
		fmt.Println(num1 - num2)
	case "*":
		fmt.Println(num1 * num2)
	case "/":
		fmt.Println(num1 / num2)
	default:
		fmt.Println("error! operator is not correct")
	}
}

// go run calculator.go