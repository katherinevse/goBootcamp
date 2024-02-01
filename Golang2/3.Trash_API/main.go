package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var (
	port string  = "8080"
	db   []Pizza // база данных
)

func init() {
	pizza1 := Pizza{
		ID:       1,
		Diameter: 22,
		Price:    500,
		Title:    "Margarita",
	}
	pizza2 := Pizza{
		ID:       2,
		Diameter: 22,
		Price:    550,
		Title:    "BBQ",
	}
	pizza3 := Pizza{
		ID:       3,
		Diameter: 22,
		Price:    700,
		Title:    "Pepperooni",
	}
	db = append(db, pizza1, pizza2, pizza3)
}

// наша модель
type Pizza struct {
	ID       int    `json:"id"`
	Diameter int    `json:"diameter"`
	Price    int    `json:"price"`
	Title    string `json:"title"`
}

// вспомогательная функция для модели  (модельный метод)
func FindPizzaById(id int) (Pizza, bool) {
	var pizza Pizza
	var found bool
	for _, value := range db {
		if value.ID == id {
			pizza = value
			found = true
			break
		}
	}
	return pizza, found

}

func GetAllPizzas(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	log.Println("Кто-то вызвал данный запрос обо всех пиццах")
	w.WriteHeader(200)            // StatusCode для запроса ЗАЧЕМ ЭТО ПРОПИСЫВАТЬ?
	json.NewEncoder(w).Encode(db) //Сериализация+ запись в w NewEncoder- подготавливает экземляр записи Encode- записывает в байты, записывает в формате json в w, после этого w закрывается
}

func GetPizzaById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// считаем id из строки запроса и конвертируем его в инт
	//получаем значение из id
	vars := mux.Vars(r)                 // извлекает переменные маршрута из объекта http.Request.{id} - это переменная маршрута, эти значения могут быть извлечены с использованием mux.Vars(r)
	id, err := strconv.Atoi(vars["id"]) // отдаем то, что лежит в поле id
	if err != nil {
		log.Fatal(err) // пока что так
	}
	log.Println("Trying to send to client pizza with id", id)
	pizza, ok := FindPizzaById(id)
	if ok {
		w.WriteHeader(200)
	}

}

func main() {
	log.Println("Trying to start REST API pizza!")
	// Инициализируем маршрутизатор
	router := mux.NewRouter()
	//1. Если на вход пришел запрос /pizzas
	router.HandleFunc("/pizzas", GetAllPizzas).Methods("GET")
	//2. Если на вход пришел запрос вида /pizza/{id}
	router.HandleFunc("/pizza/{id}", GetPizzaById).Methods("GET")
	log.Println("Router configured successfully! Let's go!")
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
