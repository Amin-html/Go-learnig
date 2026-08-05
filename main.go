package main

import (
	"fmt"
)

// ⭐ Дополнительное задание (для портфолио)
var tasks []string

func menegerTasks() {
	var operator string

	fmt.Print("Operator? ")
	fmt.Scan(&operator)

	switch operator {
	case "add":
		task := ""
		fmt.Print("Enter to tasks: ")
		fmt.Scan(&task)

		tasks = append(tasks, task)
		fmt.Println("Ok")
	case "get":
		fmt.Println(tasks)
	case "remove":
		if len(tasks) > 0 {
			tasks = tasks[:len(tasks)-1]
			fmt.Println("Ok")
		} else {
			fmt.Println("Err")
		}
	case "exit":
		return
	}

}

func main() {
	// Мини-проект №1
	numbers := []int{10, 20, 30, 40, 50}

	sum := 0

	for _, num := range numbers {
		sum += num
	}

	fmt.Println(sum/len(numbers))

	//  Мини-проект №2
	num := []int{2, 7, 15, 2, 9}

	max := num[0]

	for _, num := range num {
		if num > max {
			max = num
		}
	}
	fmt.Println(max)

	// Мини-проект №3
	names := []string{
		"Amin",
		"Ali",
		"Jenni",
	}

	search := "Jenni"

	found := false

	for _, name := range names {

		if name == search{
			found = true
			break
		}
	}

	fmt.Println(found)

	// Задание 1
	nums := [5]int{2, 3, 6, 7, 9,}

	for _, num := range nums {
		fmt.Println(num)
	}

	// Задание 2, 3, 4
	numSlice := []int{}

	sumSkice := 0

	numSlice = append(numSlice, 10, 20, 30, 40)

	for _, num := range numSlice{
		fmt.Println(num)
	}

	for _, nums := range numSlice{
		sumSkice += nums
	}

	fmt.Println(numSlice)

	min := numSlice[0]

	for _, num := range numSlice {
		if num < min {
			min = num
		}
	}

	fmt.Println(min)

	// ⭐ Задание 5
	save := []int{}

	fmt.Println("Введите 5 чисел (через пробел или Enter): ")

	for i := 0; i < 5; i++ {
		_, err := fmt.Scan(&save[i])
		if err != nil {
			fmt.Println("Ошибка ввода! Пожалуйста, введите корректное число.")
			return
		}
	}

	sumSave := 0
	maxNum := save[0]

	for _, num := range save {
		sumSave += num
		if num > maxNum {
			maxNum = num
		}
	}

	mid := sumSave / len(save)

	fmt.Println(sumSave)
	fmt.Println(maxNum)
	fmt.Println(mid)

	for {
		menegerTasks()
		fmt.Println("Continue managing tasks? (y/n): ")
		var choice string
		fmt.Scan(&choice)
		if choice == "n" {
			break
		}
	}
}
