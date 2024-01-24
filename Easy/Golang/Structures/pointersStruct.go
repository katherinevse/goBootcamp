package main

type Employee struct {
	name   string
	salary int
}

// Метод, в котором пользователь копируется и в его теле происходит работа с локальной копией
func (e Employee) setName(newName string) {
	e.name = newName
}

// Метод, в котором получатель передается по ссылке и в его теле происходит работа с ссылкой )
func (e *Employee) setSalary(newSalary int) {
	e.salary = newSalary
}

func main() {
	e := Employee{"Kathe", 500}
	//fmt.Println("Before setting parametrs", e.)

}
