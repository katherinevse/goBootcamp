package main

import (
	"fmt"
)

func newGoRoutine(done chan string) {
	fmt.Println("Hey, I'm new Goroutine!")
	done <- "Work is done!"
}

func main() {
	done := make(chan string) //make(chan string) создает новый канал, который предназначен для передачи значений строкового типа (string)
	//Теперь переменная done представляет собой канал, который может использоваться для отправки и получения значений типа string.
	go newGoRoutine(done)
	result := <-done // в этой точке выполнение main горутины останавливается до тех пор, пока в канал кто-нибудь не запишет данные!
	fmt.Println("Main goroutine work!" + result)
}
