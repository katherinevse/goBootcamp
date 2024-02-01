package main

import (
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

const (
	apiPrefix string = "api/v1"
)

var (
	port                    string
	bookResoursePrefix      string = apiPrefix + "/book"
	manyBooksResoursePrefix string = apiPrefix + "/books"
)

// func init() - это специальная функция в Go, которая вызывается автоматически при инициализации пакета.
func init() {
	err := godotenv.Load() // Загружаем переменные окружения из файла с именем ".env" с использованием пакета godotenv.
	if err != nil {
		log.Fatalln("could not find .env file!")
	}
	// Получаем значение переменной окружения "app_port".
	port = os.Getenv("app_port")
}

func main() {
	log.Println("Starting RestApi server on port", port)
	router := mux.NewRouter()
	log.Println("Router is ready to go !")
	log.Fatal(http.ListenAndServe(":"+port, router)) //Здесь происходит запуск HTTP-сервера. Функция http.ListenAndServe слушает HTTP-запросы на указанном порту.

}
