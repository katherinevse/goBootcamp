package main

import "fmt"

//Структура- явно декларированный заименованный набор состояний
//Структура исходя из своего описания, отвечает на вопрос- из чего я должен состоять, чтобы считаться тем типом, который декларируется структурой
//структура описывает паттерн состояния ( из чего состоит)

//Интерфейсы - это явнот декларированный набор сигнатур поведений ( чаещ всего в виде набора методов)
//Интерфейсы в го декларируют полу-абстрактный тип
//Интерфейс - описывает ПАТТЕРН поведения ( отчечает на вопрос: что я должен УМЕТЬ ДЕЛАТЬ, чтобы считаться тем типом, который декларирует интерфейс

// Объявление интерфейса
type figureBuilder interface {
	//Набор сигнатур методов, которые необходимо реализовать в структуре-претенденте
	// у претендента должен быть метод area(), возвращающий тип int
	Area() int
	// у претендента должен быть метод perimeter(), возвращающий тип int
	Perimeter() int
}

// Декларируем претендентов
//Первый претендент - это прямоугольник,
//У НЕГО ЕСТЬ 2 МЕТОДА area and perimeter, когда эти методы реализованы,
//говорят, что  Rectangle удовлетворяет условиям интерфейса figureBuilder
//Rectangle реализует интерфейс figureBuilder

type Rectangle struct {
	length, width int
}

func (r Rectangle) Area() int {
	return r.length * r.width
}

func (r Rectangle) Perimeter() int {
	return 2 * (r.length + r.width)
}

// второй претендент - это окружность
type Circle struct {
	radius int
}

// площадь окружности
func (c Circle) Area() int {
	return 3 * c.radius * c.radius
}
func (c Circle) Perimeter() int {
	return 2 * 3 * c.radius
}

func main() {
	//4. Создадим 2 экземпляра
	r1 := Rectangle{10, 20}
	r2 := Rectangle{10, 20}
	r3 := Rectangle{10, 20}
	c1 := Circle{5}
	c2 := Circle{5}
	c3 := Circle{5}

	//5.Посчитаем общую площадь этих 2 фигур
	//total:= r.Area() + c.Area()
	rectangles := []Rectangle{r1, r2, r3}
	totalSumRec := 0
	for _, val := range rectangles {
		totalSumRec += val.Area() // подсчет площади в val.Area()
	}
	fmt.Printf("Посчет всех площадей Rectangle %d\n", totalSumRec)

	circleses := []Circle{c1, c2, c3}
	totalSumCir := 0
	for _, val := range circleses {
		totalSumCir += val.Area()
	}
	fmt.Printf("Посчет всех площадей Circle %d", totalSumCir)

}
