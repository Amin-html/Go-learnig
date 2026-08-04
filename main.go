package main

import (
	"fmt"
)

	// Мини-проект №1
func greet(name string) {
	fmt.Println("Hello 👋", name)
}

// Мини-проект №2
func square(num int) int {
	return num * num
}
// Мини-проект №3
func isAdult(age int) bool {
	if age >= 18 {
		return true
	}
	return false
}
// Задание 1
func sayHello(name string) {
	fmt.Println("Welcome to the Go!", name)
}

// Задание 2
func sum(a int, b int) int {
	return a + b
}
// Задание 3
func multiply(a int, b int) int {
	return a * b
}
// Задание 4
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
// ⭐ Задание 5
func table(number int) {
	for i := 1; i <= 10; i++ {
		result := number * i
		fmt.Printf("%d x %d = %d\n", number, i, result)
	}
}

func calculator(a, b int, operator string) int {
	switch operator {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		if b != 0 {
			return a / b
		}
	}
	return 0
}

func main() {
	// Мини-проект №1
	greet("Amin")

	// Мини-проект №2
	result := square(5)
	fmt.Println("Result: ", result)

	// Мини-проект №3
	isAdult := isAdult(17)
	if isAdult {
		fmt.Println("Access")
	} else {
		fmt.Println("Denied")
	}
	// Задание 1
	sayHello("Amin")

	// Задание 2
	resultSum := sum(5, 10)
	fmt.Println("Sum: ", resultSum)

	// Задание 3
	resultMultiply := multiply(6, 6)
	fmt.Println("Multiply: ", resultMultiply)

	// Задание 4
	resultMax := max(15, 16)
	fmt.Println("Max num: ", resultMax)

	// ⭐ Задание 5
	table(7)

	// ⭐ Дополнительное задание (для портфолио)
	resultCalculator := calculator(10, 0, "/")
	fmt.Println("Calculator result: ", resultCalculator)
}