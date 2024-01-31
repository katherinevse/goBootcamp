package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var (
	//Порт запуска приложения
	port string = "8080"
	//Наша "база данных"
	//db []Pizza
)

func main() {
	log.Println("trying to start REST Api pizza")
	//Инициализация маршрутизатора
	route := mux.NewRouter()
	route.HandleFunc("/pizzas", GetAllPizza).Methods("GET")
	log.Fatalln(http.ListenAndServe(":"+port, nil))
}
