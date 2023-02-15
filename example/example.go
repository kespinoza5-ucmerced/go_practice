package main

import (
	"fmt"

	"github.com/kespinoza5-ucmerced/go_practice/math/simplemath"
)

// add_one takes a value by reference and adds on to the value
func add_one(a *int) {
	*a = *a + 1
}

// creates a pointer with value initialized to one
func create_pointer() *int {
	a := 1
	return &a
}

func main() {
	a := 16
	b := 2
	c := a + b

	fmt.Printf("Hello, C-%d0304 Rick!\n", c)

	fmt.Printf("This is the max of a = %d and b = %d: %d\n", a, b, simplemath.Max(a, b))

	add_one(&a)

	fmt.Printf("Here is a after add_one: %d\n", a)

	d := create_pointer()

	fmt.Printf("Here is new var d after create_pointer(): %d\n", *d)
}
