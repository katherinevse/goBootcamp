package main

import "fmt"

func TypeFinder(i interface{}) {
	switch i.(type) {
	case string:
		fmt.Println("String!!!")
	case int:
		fmt.Println("INT!!!")
	default:
		fmt.Println("heze")

	}
}

func main() {
	TypeFinder("hello")
}
