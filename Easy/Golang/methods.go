package main

import "fmt"

type User struct {
	name    string
	surname string
	age     uint16
}

func (u *User) getAllInfo() string {
	u.name = "Hach"
	return fmt.Sprintf("Name %s. Surname %s. Ahe is %d", u.name, u.surname, u.age)
}

func main() {
	cat := User{"Kotenochek", "Kotov", 19}
	fmt.Println(cat.getAllInfo())
	fmt.Println(cat)
}

//Попробуй создать такой метод и изменить внутри него имя.
//Если без указателя сделаешь, то имя не изменится, а если с указателем, то имя поменяетс
//Представь, у тебя есть структура User и там хранится три поля: name, surname, phone. Когда ты пишешь так:
//
//func (u User) SayHello() {}
//
//У тебя создается еще один объект юзера и происходит копирование name, surname и phone.
//Когда ты пишешь так:
//func (u *User) SayHello() {}
//Копирования всех этих объектов не происходит, ты работаешь не с копией, а с тем же объектом
