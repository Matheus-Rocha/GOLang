package main

import (
	"fmt"
)

type Lista struct {
	qtd    int
	inicio *Registro
}

type Registro struct {
	valor   int
	proximo *Registro
}

func main() {

	fmt.Println(("Startando o programa de Grafos..."))

	lista1 := Lista{
		qtd:    1,
		inicio: &Registro{},
	}

	lista1.InsereOrdenado(4)
	lista1.InsereOrdenado(1)
	lista1.InsereOrdenado(2)
	lista1.InsereOrdenado(6)
	lista1.InsereOrdenado(5)
	lista1.InsereOrdenado(10)
	lista1.MostrarLista()

}

func (l *Lista) InsereOrdenado(valor int) {

	if (&Lista{}) == l {
		fmt.Println("Lista vazia")
	}

	var aux *Registro = &Registro{}
	var ant *Registro = &Registro{}
	var novo *Registro = &Registro{}

	novo.valor = valor

	if l.inicio.valor == 0 {
		l.inicio.valor = novo.valor

	} else {

		aux = l.inicio
		is_insert := false
		for {
			if (aux == nil || is_insert == true) {
				break
			} else {

				if novo.valor == aux.valor {
					fmt.Println("Valor repetido")
				}

				if aux.valor < novo.valor{
					ant = aux
					aux = aux.proximo
				} else {


					if ant.valor == 0 {
						l.inicio = novo
					} else {
						ant.proximo = novo
					}
					novo.proximo = aux
					is_insert = true
				}
			}
		}
		if is_insert == false {
			ant.proximo = novo
			is_insert = true
		}
	}
	l.qtd++
}

func (l *Lista) MostrarLista() {

	var aux *Registro
	if (&Lista{}) == l {
		fmt.Println("Lista não existe")
	}
	if l.inicio == nil {
		fmt.Println("Lista vazia.")
	}

	aux = l.inicio
	fmt.Println("Mostrando a lista...")

	for i := 0; i < l.qtd; i++ {
		if aux != nil {
			fmt.Printf("%v \n", *&aux.valor)
			aux = aux.proximo
		}
	}
}