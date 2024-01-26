package main

import (
	"fmt"
	"unicode/utf8"
)

// Определение интерфейса
type bigWord interface {
	IsBig() bool
}

// Наш претендент, которого нужно научить делать isbig
type Mystring string

// Реализация IsBig у претендента mystring
func (ms Mystring) IsBig() bool {
	return utf8.RuneCountInString(string(ms)) > 10

}

func main() {
	sample := Mystring("AbdaOOOm")
	var interfaceSample bigWord //Объявление переменной, типа интерфейс Bigword
	interfaceSample = sample    //присваивание возможно потому что  sample( Mystring) удовлетворяет интерфейсу BigWord
	fmt.Println("Isbig", interfaceSample.IsBig())

	////попытка присвоить значение с типажом, неудовлетворяющеему интерфейсу
	//var interfaceSample bigWord
	//interfaceSample = "Helloychinka!" //отсутсвует метод, нужен именно mystring2

}
