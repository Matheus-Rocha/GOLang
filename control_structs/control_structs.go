package main

import (
	"fmt"
)

type student struct {
	name     string
	lastname string
	age      int
	course   string
}

func main() {
	student1 := student{
		name:     "Matheus",
		lastname: "Rocha",
		age:      25,
		course:   "Análise e Desenvolvimento de Sistemas",
	}
	fmt.Println("--- If-else ---")

	if student1.age >= 18 {
		fmt.Println("higher than legal age")
	} else {
		fmt.Println("lower than legal age")
	}

	// if init
	if num := student1.age; num >= 18 {
		fmt.Println("higher than legal age")
	} else {
		fmt.Println("lower than legal age")
	}
	// variable num only exist in if scope than its value is removed from memory

	if num := 12; num >= 0 {
		fmt.Println("higher or equal than zero")
	} else if num < -100 {
		fmt.Println("lower than -100")
	} else {
		fmt.Println("Value between 0 and -100")
	}

	fmt.Println("--- Switch ---")
	fmt.Println("Function WeekDay:")
	fmt.Println(weekDay(1))
	fmt.Println(weekDay(3))
	fmt.Println(weekDay(12))

	fmt.Println("Function WeekDay2:")
	fmt.Println(weekDay2(1))
	fmt.Println(weekDay2(2))
	fmt.Println(weekDay2(5))
	fmt.Println(weekDay2(120))

}

// Switch case:
func weekDay(num int) string {
	switch num {
	case 1:
		return "Sunday"
	case 2:
		return "Monday"
	case 3:
		return "Tuesday"
	case 4:
		return "Wednesday"
	case 5:
		return "Thursday"
	case 6:
		return "Friday"
	case 7:
		return "Saturday"
	default:
		return "Invalid day"
	}
}

func weekDay2(num int) string {
	var weekday string
	switch {
	case num == 1:
		weekday = "Sunday"
		fmt.Println("More than one action in switch")
	case num == 2:
		weekday = "Monday"
	case num == 3:
		return "Tuesday"
	case num == 4:
		return "Wednesday"
	case num == 5:
		return "Thursday"
	case num == 6:
		return "Friday"
	case num == 7:
		return "Saturday"
	default:
		return "Invalid day"
	}
	return weekday
}
