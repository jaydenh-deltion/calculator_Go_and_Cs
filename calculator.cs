using System;

class Calculator
{
    static void Main()
    {
        while (true)
        {
            Console.WriteLine("Welcome to my calculator");

            double num1 = getNumber("Enter first number: ");
            string op = getOperator("Enter operator (+,-,*,/): ");
            double num2 = getNumber("Enter second number: ");

            if (op == "/" && num2 == 0)
            {
                Console.WriteLine("Error: Division by zero is not allowed.");
                continue;
            }
        

            double result = 0;

            switch (op)
            {
                case "+":
                    result = num1 + num2;
                    Console.WriteLine("Result: " + result);
                    break;
                case "-":
                    result = num1 - num2;
                    Console.WriteLine("Result: " + result);
                    break;
                case "*":
                    result = num1 * num2;
                    Console.WriteLine("Result: " + result);
                    break;
                case "/":
                    result = num1 / num2;
                    Console.WriteLine("Result: " + result);
                    break;
                default:
                    Console.WriteLine("Invalid operator");
                    break;
            }
                Console.Write("Do you want to perform another calculation? press 'n' to quit otherwise press enter to continue: ");
                string? choice = Console.ReadLine()?.ToLower();

                if (choice == "n")
                {
                    Console.WriteLine("Thank you for using my calculator! goodbye");
                    return;
                }
            
        }
    }


    static double getNumber(string prompt)
    {
        while (true)
        {
            Console.Write(prompt);
            if (double.TryParse(Console.ReadLine(), out double result))
            {
                return result;
            }
            Console.WriteLine("Invalid input. Please try a number");
        }
    }

    static string getOperator(string prompt)
    {
        while (true)
        {
            Console.Write(prompt);
            string? input = Console.ReadLine();

            if (input == "+" || input == "-" || input == "*" || input == "/")
            {
                return input;
            }
            Console.WriteLine("Invalid input. Please try again");
        }
    }

}
