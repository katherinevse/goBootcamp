package main

import "fmt"

func main() {
	todoList := []string{
		"Полить цветы",
		"Выпить стакан воды",
		"Принять таблетки",
		"Сходить в магазин",
	}
	fmt.Println("Длина списка: ", len(todoList))
	fmt.Println("Емкость списка: ", cap(todoList))

	todoList = append(todoList, "Сходить к бабушке")

	fmt.Println("Длина списка после: ", len(todoList))
	fmt.Println("Емкость списка после: ", cap(todoList))

	for index, item := range todoList {
		fmt.Printf("%d %s\n", index+1, item)

	}

}
