package main

import (
	"errors"
	"fmt"
	"time"
)

// Data types:
type DataTypes struct {
	u uint
	v int
	a int8
	b int16
	c int32
	d int64
	e float32
	f float64
	g string
	h time.Time
	i rune  // alias of int32
	j byte  // alias of int8
	k bool
	l error
}

func main() {
	// Declaration:
	var variable1 string
	var variable2 string = "string2"
	variable1 = "string1"
	fmt.Println(variable1, variable2)

	var (
		variable3 string
		variable4 string = "string4"
	)
	variable3 = "string3"
	fmt.Println(variable3, variable4)

	variable5 := "string5"
	fmt.Println(variable5)

	variable6, variable7 := "string6", "string7"
	fmt.Println(variable6, variable7)

	variable6, variable7 = variable7, variable6
	fmt.Println(variable6, variable7)

	// Number of ASCII table
	char := 'M'
	fmt.Println(char)

	// Initial value
	var var1 string  // Empty string
	fmt.Println(var1)

	// int and float values are initialized with value = 0
	var var2 int
	fmt.Println(var2)

	var var3 int32
	fmt.Println(var3)

	var var4 int64
	fmt.Println(var4)

	var var5 float32
	fmt.Println(var5)

	var var6 bool  // false
	fmt.Println(var6)

	var var7 error // nil
	fmt.Println(var7)
	
	var8 := errors.New("New error")
	fmt.Println(var8)

}
