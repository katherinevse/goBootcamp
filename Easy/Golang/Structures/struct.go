package main

import "fmt"

type Student struct {
	firstName string
	lastName  string
	age       int
}

type anotherStudent struct {
	firstName, lastName, groupName string
	age, courseNumber              int
}

func printStudent(gg Student) {
	fmt.Println("===============")
	fmt.Println("FirstName: ", gg.firstName)
	fmt.Println("LastName: ", gg.lastName)
	fmt.Println("Age : ", gg.age)

}

func main() {
	//Создание представителей структуры
	stud1 := Student{
		firstName: "Fedor",
		lastName:  "Petrov",
		age:       15,
	}
	stud2 := Student{"Max", "Ivanov", 14}
	printStudent(stud1)
	printStudent(stud2)
	// что если не все поля структуры определить?
	stud3 := Student{
		firstName: "kate",
	}
	printStudent(stud3)
	fmt.Println("=====Anon ==========")

	//Анонимные структуры, нет имени
	anonStud := struct {
		age      int
		groupId  int
		profName string
	}{
		age:      23,
		groupId:  2,
		profName: "Alexeev",
	}
	fmt.Println("anon student", anonStud)

	//7.Доступ к состояниям
	fmt.Println("====================")

	studVova := Student{"Vova", "ivanov", 44}
	fmt.Println(studVova.age)
	studVova.age += 2
	fmt.Println("New age", studVova.age)

	//Указатели на экземпляры структур
	studPointer := &Student{
		firstName: "Igor",
		lastName:  "Ivlev",
		age:       45,
	}
	fmt.Println("Value pointer: ", studPointer)
	secondPointer := studPointer
	(*secondPointer).age += 20
	fmt.Println("Value pointer Mod: ", studPointer)

	//Работа с доступом к полям через указатель
	fmt.Println("Age via (*...)", (*studPointer).age)
	fmt.Println("Age via .age ", studPointer.age) //неявно происходит разыменование указателя и запрос соответсвующего поля

}
