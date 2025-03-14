package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var a float64
	var b float64
	var input string
	var err error
	for {
		for {
			fmt.Print("Введите первое число:")
			fmt.Scan(&input)
			a, err = isValidNumber(input)
			if err == nil {
				break
			} else {
				fmt.Println("Ошибка:", err)
			}
		}

		for {
			fmt.Print("Введите второе число:")
			fmt.Scan(&input)
			b, err = isValidNumber(input)
			if err == nil {
				break
			} else {
				fmt.Println("Ошибка:", err)
			}
		}

		for {
			fmt.Println("Введите математический знак (+, -, *, /):")
			fmt.Scanln(&input)

			if isValidOperator(input) {
				break
			}
			fmt.Println("Ошибка: введен неверный знак. Пожалуйста, используйте +, -, * или /.")
		}

		switch input {
		case "+":
			fmt.Println(a + b)
		case "-":
			fmt.Println(a - b)
		case "*":
			fmt.Println(a * b)
		case "/":
			fmt.Println(a / b)
		}

	}
}

func isValidNumber(input string) (float64, error) {
	input = strings.TrimSpace(input)

	num, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return 0, fmt.Errorf("неверное число: %s", input)
	}

	return num, nil
}

func isValidOperator(input string) bool {
	validOperators := []string{"+", "-", "*", "/"}

	input = strings.TrimSpace(input)
	for _, op := range validOperators {
		if input == op {
			return true
		}
	}
	return false
}
