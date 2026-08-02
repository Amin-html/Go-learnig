package main

import (
	"fmt"
)

func main()  {
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
}