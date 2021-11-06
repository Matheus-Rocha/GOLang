package main

import (
	"fmt"
	"math"
)

type Object interface {
	area() float64
}

type Rectangle struct {
	width  float64
	height float64
}

type Circle struct {
	radius float64
}

func main() {
	r := Rectangle{
		width: 10,
		height: 20,
	}

	c := Circle{
		radius: 5,
	}

	// using the interface
	printArea(&r)
	printArea(&c)

	// using generic interface
	fmt.Println("\n--- Generic interfaces ---")
	name := map[string]string {
		"Name": "Matheus",
		"Lastname": "Rocha",
	}
	Generic("String")
	Generic(12)
	Generic(12.777777)
	Generic(-12.777777)
	Generic(false)
	Generic(name)

	
	maps := map[interface{}]interface{} {
		"String": 12,
		2: "String2",
		true: 12.777777,
		-4: name,
	}
	fmt.Println(maps)
}

func printArea(obj Object) float64 {
	fmt.Printf("The area is: %.2f\n", obj.area())
	return obj.area()
}

func (r *Rectangle) area() float64 {
	return r.height * r.width
}

func (c *Circle) area() float64 {
	return math.Pi * (c.radius * c.radius)
}

// Generic interface
func Generic(interfac interface{}) {
	fmt.Println(interfac)
}
