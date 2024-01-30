package main

import "fmt"

type Employee struct {
	name     string
	position string
	salary   int
	currency string
}

// Методы- функции, привязанные к определенным структурам
//func(s Struct) MethodName( parametrs type) typereturn{}
// reciever -получаетель метода

func (e Employee) displayInfo() {
	fmt.Println("Name ", e.name)
	fmt.Println("Position", e.position)
	fmt.Printf("Salary %d %s\n", e.salary, e.currency)
}

func main() {
	emp := Employee{"Bob", "Middle", 1700, "USD"}
	emp.displayInfo() //вызов метода
}
