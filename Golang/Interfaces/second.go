//создадим пустой интерфейс

//Если интерфейс пустой, то ему удовлетворет вообще кто угодно

package main

import "fmt"

type Empty interface {
}

func Describer(pretendent interface{}) {
	fmt.Printf("Value: %v\n", pretendent)
}

func main() {
	strSample := "Hello "
	Describer(strSample)
	Describer("Lol")
}
