package main

import (
	"fmt"
)

func main() {
	fmt.Println(aproveStudent(6, 6))
	fmt.Println("After function...")
}

// panic if media = 6
func aproveStudent(n1, n2 int) bool {
	defer recoverExec()

	average := (n1 + n2) / 2

	if average > 6 {
		return true
	} else if average < 6 {
		return false
	}

	panic("Average equal to 6!!")
}

func recoverExec() {
	fmt.Println("Trying to recover the program")
	if r := recover(); r != nil {
		fmt.Println("Recovered sucessfully!")
	}
}
