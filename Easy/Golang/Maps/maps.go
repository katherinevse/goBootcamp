package main

import "fmt"

func main() {
	//1.Мапы это набор: ключ и значение
	mapper := make(map[string]int)
	fmt.Println("Empty map: ", mapper)

	//2.Добавление в мапу
	mapper["Alice"] = 24
	mapper["Mark"] = 25
	fmt.Println("Mapper after adding: ", mapper)
	//Инициализация мапы с указанием
	newMapper := map[string]int{
		"Alice": 100,
		"Max":   200,
	}
	fmt.Println("New Mapper: ", newMapper)

}
