package main

import "fmt"

type Rectangle struct {
	length, width int
}

//для имен конструкторов в го договорились использовать функции со следующим названием
// *New() если данный конструктор на один файл( в файле присутсвует описание олько одной структуры)
// *NewStructName- если в файле присутсвтуют еще какие-либо структуры

func New(newLength, newWidth int) Rectangle {
	return Rectangle{length: newLength, width: newWidth}
}

func main() {
	r := New(10, 20)
	fmt.Printf("Type %T and value %v", r, r)
}
