package main

import "fmt"

func main() {
	for {
		var a float32
		var b float32
		var value string
		fmt.Print("Введите первое число:")
		fmt.Scan(&a)
		fmt.Print("Введите второе число:")
		fmt.Scan(&b)
		fmt.Print("Введите математическую операцию ( + , - , * , / ) :")
		fmt.Scan(&value)

		switch value {
		case "+":
			fmt.Println(a + b)
		case "-":
			fmt.Println(a - b)
		case "*":
			fmt.Println(a * b)
		case "/":
			fmt.Println(a / b)
		default:
			fmt.Println("unknown literal")
		}

	}
}
