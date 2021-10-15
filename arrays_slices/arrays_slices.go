package main

import (
	"fmt"
	"reflect"
)

func main() {
	var array1 [5]int
	fmt.Println(array1)

	var array2 [5]string
	array2[0] = "Position 0"
	array2[1] = "Position 1"
	array2[2] = "Position 2"
	array2[3] = "Position 3"
	array2[4] = "Position 4"
	fmt.Println(array2)

	array3 := [4]string{"Pos 0", "Pos 1", "Pos 2", "Pos 3"}
	fmt.Println(array3)

	array4 := [...]int{1, 2, 3, 4, 5, 6, 7}  // [...] is based on lenght of your declaration but not dynamic
	fmt.Println(array4)

	slice := []int{11, 12, 13, 14, 15}
	fmt.Println(slice)
	// reflect is a native lib that un can debugg the type of a variable
	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(reflect.TypeOf(array4))

	slice = append(slice, 20)
	fmt.Println(slice)

	// [1:3] takes the number of position array[1] and array[2] but not array[3]
	slice2 := array3[1:3]
	fmt.Println(slice2)

	// same as a pointer
	array3[1] = "Position changed"
	fmt.Println(slice2)

	// Internal arrays
	fmt.Println("\n------ Internal arrays ------")

	slice3 := make([]float64, 10, 11)
	fmt.Println("\n--- slice3: Slice / Lenght / Capacity ---")
	fmt.Println(slice3)
	fmt.Println(len(slice3))  // length
	fmt.Println(cap(slice3))  // capacity

	slice3 = append(slice3, 1)
	slice3 = append(slice3, 1)
	fmt.Println("\n--- slice3: Slice / Lenght / Capacity ---")
	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3)) // 2 * capacity

	slice4 := make([]string, 3) // make a slice with same len and cap by default
	fmt.Println("\n--- slice4: Slice / Lenght / Capacity ---")
	fmt.Println(slice4)
	fmt.Println(len(slice4))
	fmt.Println(cap(slice4))
	slice4 = append(slice4, "test")
	fmt.Println("\n--- slice4: Slice / Lenght / Capacity ---")
	fmt.Println(slice4)
	fmt.Println(len(slice4))
	fmt.Println(cap(slice4))
}