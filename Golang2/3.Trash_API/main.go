package mainтт

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	port string  = "8080"
	db   []Pizza // база данных
)

func init() {
	pizza1:= Pizza{
		ID: 1,
		Diameter: 22,
		Price: 500,
		Title: "Margarita"
	}
	pizza2:= Pizza{
		ID: 2,
		Diameter: 22,
		Price: 550,
		Title: "BBQ"
	}
	pizza2:= Pizza{
		ID: 3,
		Diameter: 22,
		Price: 700,
		Title: "Pepperooni"
	}
	db = append(db, pizza1,pizza2,pizza3)
}

// наша модель
type Pizza struct {
	ID       int `json:"id"`
	Diameter int `json:"diameter"`
	Price    int `json:"price"`
	Title    int `json:"title"`
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
		return pizza, found
	}

}

func GetAllPizzas(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Content-Type", "application/json")
}
func GetPizzaById(w http.ResponseWriter, request *http.Request) {}

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
