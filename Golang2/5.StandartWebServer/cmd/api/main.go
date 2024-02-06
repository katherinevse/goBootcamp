package main

import (
	"Downloads/IT/goBootcamp/Golang2/5.StandartWebServer/internal/app/api"
	"log"
)

var ()

func init() {

}

func main() {
	//server instance initialization
	server := api.New

	if err := server.Start(); err != nil {
		log.Fatal(err)

	}
}
