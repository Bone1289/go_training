package main

import "fmt"

type person struct {
	first      string
	last       string
	favFlavors []string
}

func main() {
	p1 := person{
		first: "James",
		last:  "Bond",
		favFlavors: []string{
			"chocolate",
			"martini",
		},
	}

	fmt.Println(p1)

	fmt.Println("First name", p1.first)
	fmt.Println("Last name", p1.last)
	fmt.Println("Favorit Flavors:")
	for i, v := range p1.favFlavors {
		fmt.Printf("\t%v %v", i, v)
	}
}
