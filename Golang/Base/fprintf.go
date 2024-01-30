package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("output.txt")
	if err != nil {
		fmt.Println("Ошибка при создании файла:", err)
		return
	}
	defer file.Close()
	fmt.Println(file.Name())

	name := "John"
	age := 25
	fmt.Fprintf(file, "Привет, меня зовут %s и мне %d лет\n", name, age)

}
