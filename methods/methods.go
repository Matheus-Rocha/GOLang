package main

import "fmt"

type user struct {
	name     string
	lastname string
	age      uint8
}

func main() {
	user1 := user{
		name: "Matheus",
		lastname: "Rocha",
		age: 25,
	}
	user2 := user{
		name: "Marcos",
		lastname: "Rocha",
		age: 43,
	}
	user3 := user{
		name: "Deborah",
		lastname: "Coelho",
		age: 47,
	}

	// using the struct method
	user1.save()
	user2.save()
	user3.save()

	fmt.Println(user1.name, user1.LegalAge())
	fmt.Println(user2.name, user2.LegalAge())
	fmt.Println(user3.name, user3.LegalAge())

	fmt.Println(user1.age)
	user1.birth()
	fmt.Println(user1.age)
}

func (u user) save() {
	fmt.Printf("Saving the user...\nname: %s\nlastname: %s\nage: %d", u.name, u.lastname, u.age)
}

func (u user) LegalAge() string {
	if u.age >= 18 {
		return "Higher or equal than legal age"
	}
	return "Lower than legal age"
}

func (u *user) birth() {
	fmt.Println("Happy birthday!")
	u.age++
}
