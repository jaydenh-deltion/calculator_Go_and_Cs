package main 

import ("fmt")

func getNumber(prompt string)(float64){

	for {
		var input float64
		fmt.Print(prompt)
		_, err := fmt.Scanf("%f", &input)
		if err != nil{
			fmt.Print("Invalid input. Please enter a number: ")
			fmt.Scanln()
			continue
		}
		return input
		
	}
}

func getOperator() string {
	for {
		var oper string
		fmt.Print("Enter operator: +, -, *, /:\n ")
		_, err := fmt.Scanln(&oper)
		if err != nil || (oper != "+" && oper != "-" && oper != "*" && oper != "/"){
			fmt.Print("Invalid input. Please enter an operator: ")
			continue
		}
		return oper
		}
	}

func main(){

	for {

	fmt.Print("welcome to my calculator\n") // print function 

	var num1 , num2 float64
	num1 = getNumber("enter the first number: ")

	var operator string // variable declaration
	operator = getOperator()
	
	num2 = getNumber("enter the second number: ")

		if operator == "/" && num2 == 0 {
		fmt.Print("Error! Dividing by zero is not allowed.")
		continue
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
		
		fmt.Println("if you want to quit press N. Or if you want to continue press enter: " )
		var quit string
		fmt.Scanln(&quit)

		if quit == "N"  {
			fmt.Println("goodbye")
			break
		}
	
	
}}


// go run calculator.go