package main

import "fmt"

func doSomething(pretendent interface{}) {
	switch pretendent.(type) {
	case string:
		fmt.Println("This is string")
	case int:
		fmt.Println("This is int")
	default:
		fmt.Println("heze")
	}
}

func main() {
	doSomething(10)
	doSomething("Hello")
	doSomething(func(a, b int) int { return a + b })

}
