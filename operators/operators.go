package main

import (
	"fmt"
)

func main() {
	// Arithmetics:
	sum := 5 + 3
	sub := 5 - 3
	mult := 5 * 3
	div := 5 / 3
	rest := 5 % 3
	fmt.Println(sum, sub, mult, div, rest)

	// Relations:
	fmt.Println(5 > 3)
	fmt.Println(5 < 3)
	fmt.Println(5 >= 3)
	fmt.Println(5 <= 3)
	fmt.Println(5 == 3)
	fmt.Println(5 != 3)

	// Logics:
	varTrue, varFalse := true, false
	fmt.Println(varTrue && varFalse)
	fmt.Println(varTrue || varFalse)
	fmt.Println(!varFalse)
	fmt.Println(!varTrue)

	// Unaries:
	num := 10
	num++
	fmt.Println(num)
	num--
	fmt.Println(num)
	num -= 1
	fmt.Println(num)
	num += 10
	fmt.Println(num)
	num *= 3
	fmt.Println(num)
	num /= 10
	fmt.Println(num)
	num %= 10
	fmt.Println(num)

	
}