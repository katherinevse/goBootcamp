package main

import "fmt"

type Rectangle struct {
	length, width int
}

//для имен конструкторов в го договорились использовать функции со следующим названием
// *new() если данный конструктор на один файл( в файле присутсвует описание олько одной структуры)
// *NewStructName- если в файле присутсвтуют еще какие-либо структуры

func New(newLenght, newWidth int) Rectangle {
	return Rectangle{newLenght, newLenght}

}

func main() {
	r := New(10, 20)
	fmt.Println("Type %T and value %v\n", r, r)
}
