package main

import (
	"fmt"
)

var (
	var_init  int
	var_init2 int
	var_init3 string
)

func main() {
	fmt.Println("Using normal return:")
	sum, sub := mathOperations(10, 20)
	fmt.Println("Sum:", sum, "Sub:", sub)

	fmt.Println("Using named return:")
	sum1, sub1 := mathOperations1(10, 20)
	fmt.Println("Sum:", sum1, "Sub:", sub1)

	fmt.Println("\n--- Variable functions ---")
	sum2 := variableSum(1, 2, 3, 4, 5)
	fmt.Println(sum2)
	sum3 := variableSum(1, 2)
	fmt.Println(sum3)
	sum4 := variableSum()
	fmt.Println(sum4)
	sum5 := variableSum(1, 2, 3, 4, 5, 100, 200)
	fmt.Println(sum5)

	// Anonimous functions
	func() {
		fmt.Println("Anonimous function")
	}()

	func(str string) {
		fmt.Println(str)
	}("Anonimous function")

	anonimous_return := func(str string) string {
		return str
	}("Anonimous function")
	fmt.Println(anonimous_return)

	// fibonacci using recursive functions
	fmt.Println("\n--- Recursive functions ---")
	position := uint(3)
	fib_sum := fibonacci(position)
	fmt.Println(fib_sum)

	for i := uint(1); i < position; i++ {
		fmt.Println(fibonacci(i))
	}

	// using closure function
	fmt.Println("\n--- Closure functions ---")
	text := "Inside main function"
	fmt.Println(text)
	f_closure := closure()
	f_closure()

	// using pointers
	fmt.Println("\n--- Functions using pointers---")
	n1 := 12

	n2 := inverseSignal(n1)
	fmt.Println(n1)
	fmt.Println(n2)

	inverseSignalWithPointers(&n2)
	fmt.Println(n2)

	// testing init attr
	fmt.Println(var_init)
	fmt.Println(var_init2) // not attr by init funct
	fmt.Println(var_init3)
}

// normal return
func mathOperations(x int, y int) (int, int) {
	sum := x + y
	sub := x - y
	return sum, sub
}

// named return
func mathOperations1(x int, y int) (sum int, sub int) {
	sum = x + y
	sub = x - y
	return
}

// variable func
func variableSum(nums ...int) int { // ...int build a slice of int
	total := 0
	for _, value := range nums {
		total += value
	}
	return total
}

// recursive func
func fibonacci(posicao uint) uint {
	if posicao <= 1 {
		return posicao
	}

	return fibonacci(posicao-2) + fibonacci(posicao-1)
}

// closure func
func closure() func() {
	text := "Inside de closure function"

	f := func() {
		fmt.Println(text)
	}
	return f
}

func inverseSignal(x int) int {
	return x * -1
}

func inverseSignalWithPointers(x *int) {
	*x = *x * -1 // this is how to acess a pointer value
}

// init function
func init() {
	fmt.Println("Passing by init function")
	var_init = 777
	var_init3 = "Init atribution"
}
