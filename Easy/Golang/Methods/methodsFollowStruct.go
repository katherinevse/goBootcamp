package main

import "fmt"

type University struct {
	city string
	name string
}

// данный метод связан только с структурой University
func (u *University) fullInfoAboutUni() {
	fmt.Printf("Uni name %s and city %s\n", u.name, u.city)
}

type Professor struct {
	fullName   string
	age        int
	University // объявляем структуру
}

func main() {
	p := Professor{
		fullName: "Bubnov Igor",
		age:      56,
		University: University{
			city: "Moscow",
			name: "MSU",
		},
	}
	p.fullInfoAboutUni()
}
