package main

import (
	"fmt"
	"time"
)

func main() {
	delay := 1 * time.Second

	loop_base(0, delay)

	loop_base2(delay)

	names := []string{"Matheus", "Rocha", "Marcos", "Jake"}
	loop_slice(delay, names)

	loop_string("GOWAY")

	mapper := map[string]string {
		"name": "Matheus",
		"lastname": "Rocha",
		"college": "IESB",
		"couse": "ADS",
		"campus": "Asa sul",
	}
	loop_map(mapper)

	while_loop()
}

func loop_base(x int, delay time.Duration) {
	for x < 7 {
		fmt.Println(x)
		x++
		time.Sleep(delay)
	}
	fmt.Println(x)
}

func loop_base2(delay time.Duration) {
	for j := 0; j < 7; j++ {
		fmt.Println("Adding j...", j)
		time.Sleep(delay)
	}
}

func loop_slice(delay time.Duration, stringsIterable []string) {
	for key, value := range stringsIterable {
		fmt.Println("Key", key, "Value", value)
		time.Sleep(delay)
	}
}

func loop_string(obj_string string) {
	for key, value := range obj_string {
		fmt.Println("Key", key, "Value", string(value))
	}
}

func loop_map(mapper map[string]string) {
	for key, value := range mapper {
		fmt.Println("Key", key, "Value", string(value))
	}
}

func while_loop() {
	i := 0
	for {
		if i < 10 {
			fmt.Println(i)
			i++
			time.Sleep(time.Second)
		} else {
			fmt.Println("Condition finished")
			return
		}
	}
}