package main

import (
	"fmt"
	"log"
	"net/http"
)

//w - responseWriter (куда писать ответ)
//r - request (откуда брать запрос) что от нас хотят?

// Обработчик
func GetGreet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi im new web-server")
}

//Смотрим, по какому адресу пришел запрос  Выбирает нужный обработчик

func RequestHandler() {
	http.HandleFunc("/", GetGreet)               //если придет запрос по такому адресу "/", то вызывай GetGreet
	log.Fatal(http.ListenAndServe(":8080", nil)) //запускаем в режиме "слушания"
}
func main() {
	RequestHandler()
}
