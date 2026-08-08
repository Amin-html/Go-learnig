package main

import (
	"fmt"
)

// 🎯 Мини-проект №1 — User
type User struct {
	Name string
	Age int
	Country string
}

// 🎯 Мини-проект №2 — Car
type Car struct {
	Brand string
	Model string
	Year int
}

// 🎯 Мини-проект №3 — Game
type Game struct {
	Name string
	Genre string
	Year int
}

// Задание 1
type Student struct {
	Name string
	Age int
	Course int
}

// Задание 3
func printStudent(student Student) {
	fmt.Println("Hello! 👋👋", student.Name, student.Age, student.Course)
}

// Задание 4
type Product struct {
	Name string
	Price float64
	Stock int
}

// ⭐ Задание 5
type UserCheckAge struct {
	Name string
	Age int
}

func main() {
	// 🎯 Мини-проект №1 — User
	user := User{"Amin", 17, "Kyrgyzstan"}

	fmt.Println(user.Name)
	fmt.Println(user.Age)
	fmt.Println(user.Country)

	// 🎯 Мини-проект №2 — Car
	car := Car{"Toyota", "Camry", 2020}

	fmt.Println(car.Brand)
	fmt.Println(car.Model)
	fmt.Println(car.Year)

	// 🎯 Мини-проект №3 — Game
	game1 := Game{"The Legend of Zelda", "Action-Adventure", 1986}
	game2 := Game{"Minecraft", "Sandbox", 2011}
	game3 := Game{"Tetris", "Puzzle", 1984}

	games := []Game{
		game1,
		game2,
		game3,
	}

	for _, game := range games{
		fmt.Println(game.Name, game.Genre, game.Year)
	}

	// Задание 1, 2
	student1 := Student{"Алексей", 19, 1}
	student2 := Student{"Мария", 21, 3}
	student3 := Student{"Иван", 20, 2}

	students := []Student{
		student1,
		student2,
		student3,
	}

	for _, student := range students{
		fmt.Println(student.Name, student.Age, student.Course)
	}

	// Задание 3
	printStudent(student1)

	// Задание 4
	product1 := Product{"Ноутбук", 1299.99, 15}
	product2 := Product{"Смартфон", 699.50, 42}
	product3 := Product{"Наушники",  89.00, 120}

	products := []Product{
		product1,
		product2,
		product3,
	}

	maxPrice := products[0]

	for _, product := range products {
		if product.Price > maxPrice.Price {
			maxPrice = product
		}
	}
	fmt.Println(maxPrice.Name)

	// ⭐ Задание 5

	users := []UserCheckAge{
		{"Amin", 17},
		{"Ali", 18},
		{"Renni", 20},
		{"Lol", 16},
		{"Max", 19},
	}

	for _, user := range users {
		if user.Age >= 18 {
			fmt.Println(user.Name)
		}
		
	}
	
}
