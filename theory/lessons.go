package main

import (
	"fmt"
)

func main()  {
	fmt.Println("Hello, World")

	// День 2. Переменные, типы данных и ввод
	// Задание 1, 2
	name := "Amin"
	age := 17
	height := 1.64
	isStudent := true

	fmt.Printf("Name: %s\n", name)
	fmt.Printf("Age: %d\n", age)
	fmt.Printf("Height: %.2f\n", height)
	fmt.Printf("Im Student: %t\n", isStudent)

	// Задание 3
	const id = 2254
	const Country = "Kyrgystan"
	fmt.Printf("My ID: %d\n", id)
	fmt.Printf("My Country: %s\n", Country)

	// Задание 4, 5 ⭐
	name2 := ""
	age2 := 0
	favoriteLanguage := ""

	fmt.Print("Enter your name: ")
	fmt.Scan(&name2)
	fmt.Print("Enter your age: ")
	fmt.Scan(&age2)
	fmt.Print("Enter your favorite language❤️ : ")
	fmt.Scan(&favoriteLanguage)

	fmt.Println("Hello! 👋", name2)
	fmt.Println("You", age2, "years old.")
	fmt.Println("Your favorite language", favoriteLanguage)

	// 📚 День 3. Условия (if, else, switch)
		// 🎯 Мини-проект №1, Задание 1
	age := 0

	fmt.Print("Введите возраст:")
	fmt.Scan(&age)

	if age >= 18 {
		fmt.Println("Доступ разрешен ✅")
	} else {
		fmt.Println("Доступ запрещен ❌")
	}

	// 🎯 Мини-проект №2, Задание 5 ⭐
	num1 := 0
	num2 := 0
	operator := ""

	fmt.Print("Введите первое число: ")
	fmt.Scan(&num1)
	fmt.Print("Введите второе число: ")
	fmt.Scan(&num2)
	fmt.Print("Введите оператор (+, -, *, /): ")
	fmt.Scan(&operator)

	switch operator {
	case "+":
		fmt.Println("Результат:", num1 + num2)

	case "-":
		fmt.Println("Результат:", num1 - num2)

	case "*":
		fmt.Println("Результат:", num1 * num2)

	case "/":
		if num2 !=0 {
			fmt.Println("Результат:", num1 / num2)
		} else {
			fmt.Println("Ошибка: деление на ноль ❌")
		}
	
	default:
		fmt.Println("Ошибка: неверный оператор ❌")
	}

	// Задание 2
	userName := ""
	password := ""

	fmt.Print("Введите имя пользователя: ")
	fmt.Scan(&userName)
	fmt.Print("Введите пароль: ")
	fmt.Scan(&password)

	if userName == "admin" && password == "admin123" {
		fmt.Println("Доступ разрешен ✅")
	} else {
		fmt.Println("Ошибка: неверное имя пользователя или пароль ❌	")
	}

	// Задание 3
	day := 0

	fmt.Print("Введите номер дня недели (1-7): ")
	fmt.Scan(&day)

	switch day {
	case 1:
		fmt.Println("Понедельник")
	case 2:
		fmt.Println("Вторник")
	case 3:
		fmt.Println("Среда")
	case 4:
		fmt.Println("Четверг")
	case 5:
		fmt.Println("Пятница")
	case 6:
		fmt.Println("Суббота")
	case 7:
		fmt.Println("Воскресенье")
	default:
		fmt.Println("Ошибка: неверный день недели ❌")
	}

	// Задание 4
	UserID := 0
	isTicketAvailable := false

	fmt.Print("Введите ваш ID: ")
	fmt.Scan(&UserID)
	fmt.Print("Доступен ли билет? (true/false): ")
	fmt.Scan(&isTicketAvailable)

	if UserID == 3 && isTicketAvailable == true {
		fmt.Println("Есть билет ✅")
	} else {
		fmt.Println("Доступ запрещен ❌")
	}

	// 📚 День 4. Циклы (for, break, continue)
	// Цели урока

	// После урока ты сможешь:

	// использовать for;
	// применять break;
	// применять continue;
	// создавать вложенные циклы;
	// написать несколько небольших программ.
		// 🎯 Мини-проект №1 — Таблица умножения
	for i := 1; i <= 10; i++ {
		result := 5 * i
		fmt.Printf("5 x %d = %d\n", i, result)
	}

	// 🎯 Мини-проект №2 — Сумма чисел
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Printf("Sum: %d\n", sum)

	// 🎯 Мини-проект №3 — Четные числа
	for i := 1; i <= 50; i++ {
		if i%2 == 0 {
			fmt.Printf("%d is even\n", i)
		}
	}

	// Задание 1
	for i := 1; i <= 20; i++ {
		fmt.Printf("№: %d\n", i)
	}

	// Задание 2
	for i := 20; i >= 1; i-- {
		fmt.Printf("№: %d\n", i)
	}

	// Задание 3, 4
	for i := 1; i <= 30; i++ {
		if i == 15 {
			continue
		} else if i == 18 {
			break
		}
		fmt.Printf("№: %d\n", i)
	}

	// Задание 5 ⭐
	for i := 1; i<= 10; i++ {
		for j := 1; j <= 10; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	// ⭐ Дополнительное задание (для портфолио)
	usernum := 0

	fmt.Print("Введите число: ")
	fmt.Scan(&usernum)

	for i := 1; i <= 10; i++ {
		result := usernum * i
		fmt.Printf("%d x %d = %d\n", usernum, i, result)
	}

	// 📚 День 5 — Функции (func)
	// 🎯 Цели урока

	// После этого урока ты сможешь:

	// создавать свои функции;
	// передавать параметры;
	// возвращать значения;
	// разбивать программу на части;
	// писать более чистый код.

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
