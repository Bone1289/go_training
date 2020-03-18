package main

import "fmt"

/*
Create TYPED and UNTYPED constants. Print the values of the constants.
*/

const (
	a int64 = 42
	b       = 42
)

func main() {
	fmt.Printf("%T %v \n", a, a)
	fmt.Printf("%T %v \n", b, b)
}
