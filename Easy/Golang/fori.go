//Требуется написать программу, при выполнении которой с клавиатуры считываются два натуральных числа
//A и B (каждое не более 100, A < B).
//Вывести сумму всех чисел от A до B  включительно.

//1. Объявить переменные a, b
//
//2. Объявить и инициализировать переменную s (сумма чисел от a до b). Стартовое значение = 0
//
//3. Написать функцию ввода значений для a, b с клавиатуры
//
//4. Создать цикл, инициализировать счетчик i, который равен a. Вписать условие i < b. Вписать оператор инкремента i++
//
//5. В качестве действия ввести формулу увеличения переменной s на значение i. s += i
//
//6. Вне цикла написать функцию вывода суммы чисел от a до b

package main

import (
	"fmt"
)

func main() {
	var a, b int
	s := 0
	fmt.Scan(&a)
	fmt.Scan(&b)
	for i := a; i < b; i++ {
		s += i
	}
	fmt.Println(s)

}
