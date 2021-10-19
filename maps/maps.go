package main

import (
	"fmt"
)

func main() {
	user := map[string]string {
		"nome": "Matheus",
		"lastname": "Rocha",
		"from": "Brasil",
	}

	fmt.Println(user)
	fmt.Println(user["nome"])
	fmt.Println(user["lastname"])
	fmt.Println(user["from"])

	users := map[string]map[string]int{
		"Matheus": {
			"id": 1,
			"age": 25,
		},
		"Marcos": {
			"id": 2,
			"age": 40,
		},
	}
	fmt.Println(users)

	delete(users, "Matheus")
	fmt.Println(users)
	
	users["Rocha"] = map[string]int{
		"id": 3,
		"age": 0,
		"cpf": 1234,
	}
	fmt.Println(users)
}