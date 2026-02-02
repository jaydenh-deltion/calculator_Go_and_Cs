package main 

import ("fmt")

func getNumber(prompt string)(float64){
// controleert of een getal is ingevoerd anders vraagt hij opnieuw om een getal
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
		if err != nil || (oper != "+" && oper != "-" && oper != "*" && oper != "/"){ // controleert of het 1 van deze operators is: +, -, *, /
			fmt.Print("Invalid input. Please enter an operator: ")
			continue
		}
		return oper
		}
	}

func main(){

	for {
// heet de gebruiker welkom 
	fmt.Print("welcome to my calculator\n") 

// vraagt om het eerst getal
	var num1 , num2 float64
	num1 = getNumber("enter the first number: ")

// vraagt om de operator
	var operator string // variable declaration
	operator = getOperator()
	// vraagt om het 2de getal
	num2 = getNumber("enter the second number: ")

	// delen door 0 is niet toegestaan
		if operator == "/" && num2 == 0 {
		fmt.Print("Error! Dividing by zero is not allowed.")
		continue
	}
// doet de berekening
		switch operator{
	case "+":
		fmt.Println(num1 + num2)
	case "-":
		fmt.Println(num1 - num2)
	case "*":
		fmt.Println(num1 * num2)
	case "/":
		fmt.Println(num1 / num2)

// geeft foutmelding
	default:
		fmt.Println("error! operator is not correct")
	}
		// vraagt of je door wilt gaan 
		fmt.Println("if you want to quit press N. Or if you want to continue press enter: " )
		var quit string
		fmt.Scanln(&quit)

// sluit het programma als input N is 
		if quit == "N"  {
			fmt.Println("goodbye")
			break
		}
	
	
}}


// go run calculator.go om te starten 