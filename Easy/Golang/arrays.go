package main

import "fmt"

func main() {
	var todolist = [4]string{}
	todolist[0] = "Полить цветы"
	todolist[1] = "Выпить стакан воды"
	todolist[2] = "Принять таблетки"
	todolist[3] = "Сходить в магазин"

	fmt.Printf("Количество ваших задач на сегодня: %d\n", len(todolist))

	for _, item := range todolist {
		fmt.Printf("Какой-то айтем: %s\n", item)
	}
}
