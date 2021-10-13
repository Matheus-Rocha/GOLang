package main

import (
	"fmt"
)

// objects
type Car struct {
	model string
	year uint
	brand Brand
}

type Brand struct {
	name string
}

type SUV struct {
	Car
	brand Brand
	color string
}

func main() {
	var c Car
	c.model = "Ferrari"
	c.year = 2020
	fmt.Println(c)

	brand1 := Brand{"Mercedes"}
	car1 := Car{"BMW", 2021, brand1}
	fmt.Println(car1)

	car2 := Car{model: "BMW"}
	fmt.Println(car2)

	car2 = Car{year: 2021}
	fmt.Println(car2)

	car3 := Car{}
	fmt.Println(car3)

	// GOway of object orientation but its not a real OO
	Jeep := Brand{name: "Jeep"}
	car10 := Car{}
	car10.model = "Renegade"
	car10.year = 2020
	car10.brand = Jeep

	mySUV := SUV{car10, Jeep, "Black"}

	fmt.Println(mySUV.color)
	fmt.Println(mySUV.model)
	fmt.Println(mySUV.year)
	fmt.Println(mySUV.brand)

	// methods
	car10.moved()
	car10.stopped()
	car10.parked()
	mySUV.moved()
	mySUV.stopped()
	mySUV.parked()

	// pointers
	var variable1 int = 10
	var variable2 int = variable1

	fmt.Println(variable1, variable2)

	variable1++
	fmt.Println(variable1, variable2)

	var variable3 int
	var pointer *int
	fmt.Println(variable3, pointer)

	variable3 = 12
	pointer = &variable3
	fmt.Println(variable3, pointer)

	fmt.Println(variable3, *pointer)

	variable3 = 1212
	fmt.Println(variable3, *pointer)
}

// objects methods
func (c *Car) moved() string {
	return "Car moved"
}

func (c *Car) stopped() string {
	return "Car Stopped"
}

func (c *Car) parked() string {
	return "Car parked"
}