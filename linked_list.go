package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type List struct {
	qtd    int64
	init *Register
}

type Register struct {
	value   int64
	next *Register
}

func main() {

	fmt.Println(("Starting linked list..."))

	list := List{
		qtd:    1,
		init: &Register{},
	}
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("1 - Add item in list ")
		fmt.Println("1 - Show the list ")
		fmt.Println("9 - leave ")
		scanner.Scan()
		option := scanner.Text()
		switch option {
		case "1":
			fmt.Println("Item: ")
			scanner.Scan()
			item, _ := strconv.ParseInt(scanner.Text(), 10, 64)
			list.OrdelyInsert(item)
		case "2":
			list.ShowList()
			break
		case "9":
			os.Exit(99)
		default:
			fmt.Println("Invalid input")
		}

	}

}

func (l *List) OrdelyInsert(value int64) {

	if (&List{}) == l {
		fmt.Println("Empty list")
	}

	var aux *Register = &Register{}
	var ant *Register = &Register{}
	var new *Register = &Register{}

	new.value = value

	if l.init.value == 0 {
		l.init.value = new.value

	} else {

		aux = l.init
		is_insert := false
		for {
			if (aux == nil || is_insert == true) {
				break
			} else {

				if new.value == aux.value {
					fmt.Println("Value already in list")
				}

				if aux.value < new.value{
					ant = aux
					aux = aux.next
				} else {


					if ant.value == 0 {
						l.init = new
					} else {
						ant.next = new
					}
					new.next = aux
					is_insert = true
				}
			}
		}
		if is_insert == false {
			ant.next = new
			is_insert = true
		}
	}
	l.qtd++
}

func (l *List) ShowList() {

	var aux *Register
	if (&List{}) == l {
		fmt.Println("List does not exist")
	}
	if l.init == nil {
		fmt.Println("Empty list")
	}

	aux = l.init
	fmt.Println("Printing the list...")

	for i := int64(0); i < l.qtd; i++ {
		if aux != nil {
			fmt.Printf("%v \n", *&aux.value)
			aux = aux.next
		}
	}
}