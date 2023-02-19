package main

import (
	"fmt"

	"github.com/kespinoza5-ucmerced/go_practice/go_tutorial/math"
	"github.com/kespinoza5-ucmerced/go_practice/go_tutorial/math/simplemath"
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

type IPoint interface {
	DistanceToOrigin() float64
}

type specialPoint struct {
	x float64
}

func (p specialPoint) DistanceToOrigin() float64 {
	return p.x * p.x
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

	fmt.Printf("Creating new point...\n")

	point := math.NewPoint(3, 4)
	point_2 := math.NewPoint(2, 6)

	fmt.Printf("New point created!\n")
	fmt.Printf("Distance to origin: %f\n", point.DistanceToOrigin())
	point.Double()
	fmt.Printf("Distance to origin: %f\n", point.DistanceToOrigin())

	fmt.Printf("New point created!\n")
	fmt.Printf("Distance to origin: %f\n", point_2.DistanceToOrigin())
	point_2.Double()
	fmt.Printf("Distance to origin: %f\n", point_2.DistanceToOrigin())

	// stuff with interface... I guess this is polymorphism?
	var point_3 IPoint
	point_3 = math.NewPoint(2, 3)

	fmt.Printf("Here is the distance to the IPoint: %f", point_3.DistanceToOrigin())

	point_3 = specialPoint{3}
	fmt.Printf("Special point created!\n")
	fmt.Printf("Distance to origin: %f\n", point_3.DistanceToOrigin())

	// stuff with arrays, slice, dict
	//simple list
	var list [3]int
	list[0] = 10
	list[1] = 20
	list[2] = 30

	fmt.Println(list)

	// working with slice, aka vector
	var list_2 []int

	list_2 = append(list_2, 10)
	list_2 = append(list_2, 20)
	list_2 = append(list_2, 30)
	list_2 = append(list_2, 40)
	list_2 = append(list_2, 50)

	fmt.Println(list_2)

	// with slices, random way to init them for performance reasons
	// preallocates space, up to a listed amount
	list_3 := make([]int, 4, 6)

	fmt.Println(list_3)

	// implement a map
	dict := make(map[string]int)

	dict["Jan"] = 1
	dict["Feb"] = 2
	dict["Mar"] = 3

	fmt.Println(dict)
	fmt.Println(dict["Jan"])

	// iterate through dict
	for key, value := range dict {
		fmt.Println(key, value)
	}

	// fun ways to init specialPoint
	p := specialPoint{3}

	fmt.Printf("New point created!\n")
	fmt.Printf("Distance from origin: %f\n", p.DistanceToOrigin())

	p_2 := specialPoint{
		x: 4,
	}

	fmt.Printf("New point created!\n")
	fmt.Printf("Distance from origin: %f\n", p_2.DistanceToOrigin())
}
