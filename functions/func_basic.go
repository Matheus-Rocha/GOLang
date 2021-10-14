package main

import (
	"fmt"
	"net/http"
)

func main() {
	add := add(10, 20)
	fmt.Println(add)

	// Declaring a variable of type function:
	var f = func (w int, z int) int {
		return w * z
	}

	f(3, 4)

	mult := f(4, 5)
	fmt.Println(mult)

	mathAdd, mathSub, mathMult, mathDiv := mathOperations(20, 10)
	fmt.Println(mathAdd, mathSub, mathMult, mathDiv)

	resp, err := getResponse("github.com")
	if err != nil {
		fmt.Println(resp, err)
	}
	fmt.Println(resp)


}

// simple functions
func add(x int, y int) int {
	return x + y
}

// Multiple returns
func mathOperations(x, y int) (int, int, int, int) {
	add := x + y
	sub := x - y
	mult := x * y
	div := x / y

	return add, sub, mult, div
}

// Error returns
func getResponse(site string) (string, error) {
	response, err := http.Get(site)
	if err != nil {
		return "Error", err
	}

	fmt.Println(response)

	if response.StatusCode == 200 {
		fmt.Println("Got an sucessfully")
	}

	return "Got an response", nil
}
