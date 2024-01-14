package main

import "fmt"

func main() {
	var n int
	for fmt.Scan(&n); n != 0; fmt.Scan(&n) {
		fmt.Println(n)
	}

}
