package main

import "fmt"

func main() {

	nomes := [4]string{"João", "Maria", "Pedro", "Camila"}
	fmt.Println("Estrutura de Repetição For")

	for _, nome := range nomes { // a primeira posição que pega é o indíce
		fmt.Println("Nome:", nome)
	}
}
