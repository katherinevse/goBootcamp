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
	fmt.Println("==============================")
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
	//8. Перебор элементов мапы(итерации)
	fmt.Println("==============================")
	for key, value := range employee {
		fmt.Printf("%s i value %d\n", key, value)
	}

	//9.Как удалять пары
	//9.1 Удаление существующец пары
	fmt.Println("==============================")

	fmt.Println("Before deleting", employee)
	delete(employee, "Den")
	fmt.Println("After deleting", employee)
	//9.1 Удаление несуществующей пары
	if _, ok := employee["Anna"]; ok {
		delete(employee, "Name") //ОЧЕНЬ дорогая операция!!!!
	} // так лучше, не нужно итерироваться просто так, делаем лучше прямое обращение
	fmt.Println("After deleting Name", employee)

	//10. Длина мапы== кол-во пар
	fmt.Println("==============================")
	fmt.Println("Pair amount in map: ", len(employee))

	//11. Мапа(как и слайс) ссылочный тип
	fmt.Println("==============================")

	words := map[string]string{
		"One": "Один",
		"Two": "Два",
	}
	newWords := words
	newWords["Three"] = "Три"
	delete(newWords, "One")
	fmt.Println("Words map:", words)

	//12. Сравнение мап
	//12.1 Сравнение массивов( массив можно использовать как ключ в мапе, попробовать!!)
	fmt.Println("==============================")
	if [3]int{1, 1, 1} == [3]int{1, 1, 1} {
		fmt.Println("Equal")
	} else {
		fmt.Println("Not equal")

	}
	////12.2  Сравнение слайсов (Можно сравнивать только с nil но не между собой изза того что тип ссылочный
	//if []int{1,1,1} == []int{1,1,1}{ //
	//	fmt.Println("Equal")
	//
	//}
	////12.3 Сравнение мап
	//aMap:= map[string]int{
	//	"a":=1
	//}
	//bMap:= map[string]int{
	//	"b":=1
	//}
	//НЕЛЬЗЯ СРАВНИВАТЬ МЕЖДУ СОБОЙ ТАКЖЕ
	//
}
