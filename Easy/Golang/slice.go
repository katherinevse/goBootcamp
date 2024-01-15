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

	for index, item := range todoList {
		fmt.Printf("%d %s\n", index+1, item)

	}

}
