package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Println("I'am", p.first, p.last, "and my age is", p.age)
}

func main() {
	p1 := person{
		first: "Tom",
		last:  "Mouse",
		age:   32,
	}
	p1.speak()
}
