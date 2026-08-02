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
	
}
