package main

import (
	"fmt"
)

func main()  {
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
}