package main

import (
	"fmt"
)

func main() {
	// Задание 1
	students := map[string]int{
		"Amin": 17,
		"Ali": 18,
		"Alina": 16,
	}

	for student, age := range students {
		fmt.Println(student, "-", age)
	}

	// Задание 2
	var name string
	fmt.Print("Enter to your name: ")
	fmt.Scan(&name)

	search, ok := students[name]

	if ok {
		fmt.Println(search)
	} else {
		fmt.Println("Student not found")
	}

	// Задание 3
	phones := map[string]string{
		"Any": "+996700111111",
		"Alex": "+996700222222",
		"Emmi": "+996700333333",
		"Ren": "+996700444444",
		"Ali": "+986700555555",
	}

	phone, ok := phones[name]

	if ok {
		fmt.Println(phone)
	} else {
		fmt.Println("Student not found")
	}

	// Задание 4
	products := map[string]float64 {
		"Apple": 50.0,
		"Milk": 77.9,
		"Bread": 87.3,
		"Water": 13.9,
		"Coffee": 43.0,
	}

	sum := 0.0

	for _, price := range products {
		sum += float64(price)
	}
	fmt.Println("Sum: ", sum)

	// ⭐ Задание 5
	dictionary := map[string]string {
		"Go": "язык программирования",
		"Docker": "контейнеризация",
		"Git": "система контроля версий",
	}

	var word string

	fmt.Print("Enter to words: ")
	fmt.Scan(&word)

	words, ok := dictionary[word]

	if ok {
		fmt.Println(words)
	} else {
		fmt.Println("Not found")
	}

	db()
}

// ⭐ Дополнительное задание (для портфолио)
func db() {
	mini_DB := map[string]int{}

	for {
			var operator int

		fmt.Println("1 add \n2 get \n3 delete \n4 list \n0 exit")

		fmt.Print("Emter to operator(1-0): ")
		fmt.Scan(&operator)

		switch operator {
		case 1:
			var newName string
			var newAge int

			fmt.Println("Enter to name: ")
			fmt.Scan(&newName)
			fmt.Println("Enter to age: ")
			fmt.Scan(&newAge)

			mini_DB[newName] = newAge
			fmt.Println("Ok ✅")
			fmt.Println()
		case 2:
			var searchName string
			fmt.Print("Enter to name: ")
			fmt.Scan(&searchName)

			age, ok := mini_DB[searchName]
			if ok {
				fmt.Println(age)
			} else {
				fmt.Println("map")
			}
		case 3:
			var name string
			fmt.Println("Введите имя для удаления: ")
			fmt.Scan(&name)

			_, ok := mini_DB[name]
			if ok {
				delete(mini_DB, name)
				fmt.Println("Пользователь удален.")
			} else {
				fmt.Println("Пользователь не найден.")
			}
		case 4:
			if len(mini_DB) == 0 {
				fmt.Println("База данных пуста.")
				continue
			}
			fmt.Println("Список пользователей:")
			for name, age := range mini_DB {
				fmt.Printf("- %s: %d\n", name, age)
			}
		case 0:
			fmt.Println("Программа завершена.")
			return
		default:
			fmt.Println("Неверный ввод. Попробуйте снова.")
		}
	}
}