package main

import (
	"fmt"
	"math"
)

//Преимущество методов над функциями
//1. Наличие методов улучшает консистентность кода( легче понять и увидеть в чем заключается метод)

//2. В го запрещается создавать функции с одинаковыми назыаниями,
//в то время как для методы для разных структур с одинаковыми названиями разрешены

type Circle struct {
	radius float64
}

type Rectangle struct {
	length, width int
}

// 7.5 Создадим функцию и метод для структуры Circle
func Perimeter(c Circle) float64 {
	return c.radius * 2 * math.Pi
}

func (c Circle) Perimeter() float64 {
	return c.radius * 2 * math.Pi
}

// 7.5 Создадим функцию и метод для структуры Rectangle
func (r Rectangle) Perimeter() int {
	return 2 * (r.width + r.length)
}

//ЗАПРЕЩЕНО ИСПОЛЬЗОВАТЬ ОДНИ И ТЕ ЖЕ НАЗВАНИЯ В ФУНКЦИИ ! А В МЕТОДЕ МОЖНО К РАЗНЫМ СТРУКТУРАМ
// В го разрешено создавать методы с одинаковыми именами,  в пределах одной структуры.
//Главное, чтобы получать методы в разных структурах

//func Perimeter(r)  {
//
//}
func main() {
	c := Circle{10.5}
	fmt.Println("Call function", Perimeter(c))
	fmt.Println("Call Method", c.Perimeter())
	//same
	r := Rectangle{10, 20}
	fmt.Println(r.length)
	fmt.Println(r.width)
	fmt.Println(r.Perimeter())

}
