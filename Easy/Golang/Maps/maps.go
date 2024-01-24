package main

import "fmt"

func main() {
	//1.Мапы это набор: ключ и значение
	mapper := make(map[string]int)
	fmt.Println("Empty map: ", mapper)

	//2.Добавление в мапу (Алиса-ключ, 24 значение)
	mapper["Alice"] = 24
	mapper["Mark"] = 25
	fmt.Println("Mapper after adding: ", mapper)

	//3.Инициализация мапы с указанием
	newMapper := map[string]int{
		"Alice": 100,
		"Max":   200,
	}
	newMapper["Jo"] = 3000
	fmt.Println("New Mapper: ", newMapper)

	// 4. Что может быть ключом в мапе?
	//4.1 Мапа- неупорядочена в го, не будет никакого порядка
	//4.2 Ключом в мапе может быть только сравнимый тип (==, !=)

	////5.Нулевое значение для мап
	//var mapZeroValue = map[string]int //mapZeroValue == nil
	//mapZeroValue["Alice"] = 12

	//6. Получение элементов из пар
	//6.1.Получение элемента, который представлен в мапе
	//testPerson:= "Alice"
	fmt.Println("Salary of testPerson", newMapper["Alice"])
	//6.2.Получение элемента, который НЕ представлен в мапе
	testPerson := "Kath"
	fmt.Println("Salary of new testPerson", newMapper[testPerson])
	fmt.Println(newMapper)

	//7. Проверка, есть ли ключ и значение какое-нибудь
	employee := map[string]int{
		"Den":  10,
		"Kate": 40,
		"Bob":  0,
	}
	fmt.Println("Den and value", employee["Den"])
	//7.1 При обращении по ключу возвращаетмя 2 значения
	if value, ok := employee["Den"]; ok {
		fmt.Println("Den and value:", value)
	} else {
		fmt.Println("Den does not exists in map")
	}

}
