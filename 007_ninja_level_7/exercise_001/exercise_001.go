package main

import "fmt"

type person struct {
	name string
}

func main() {
	x := 42
	fmt.Println(x)
	fmt.Println(&x)

	p1 := person{
		name: "Tom And Jerry",
	}

	fmt.Println(p1)
	changeMe(&p1)
	fmt.Println(p1)
}

func changeMe(p *person) {
	p.name = "Creed Alex"
}
