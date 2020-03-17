package main

import "fmt"

func main() {
	defer foo()
	fmt.Println("Hello, Gutys")
}

func foo() {
	defer func() {
		fmt.Println("foo Defer run")
	}()
	fmt.Println("foo run")
}
