package main

import "fmt"

//1. Для отправки данных в канал a (chan int) используем синтаксис
//a <- dataInt

//2. Для получения данных из канала используем синактсис
// dataInt := <- a

//3. Отправка и получения данных из канала - блокирующая операция!
// Это означает, что если данные отправлены в канал, то выполнение текущей программы останавливается до тех пор, пока с другой
// стороны из этого канала кто-то не считает данные.

// Аналогично и в обратную сторону. Если кто-то читает из канала, то выполнение текущей программы (горутины) останавливается до тех пор,
// пока кто-то в этот канал не отправит данные.

// 4. Пример использования каналов.
func newGoRoutine(done chan bool) {
	fmt.Println("Hey, I'm new Gorutine!")
	done <- true
}

func main() {
	done := make(chan bool)
	go newGoRoutine(done)
	<-done // в этой точке выполнение main горутины останавливается до тех пор, пока в канал кто-нибудь не запишет данные!
	fmt.Println("Main goroutine work!")
}

//Пока не получим данные из горутины 	go newGoRoutine(done) - дальше программа не работает ! останавливается
