package main

import "fmt"

func main() {
	//Criando um slice de inteiros
	var numeros[] int
	var frutas = []string{"Maçã", "Banana", "Laranja", "Uva", "Morango"}
	// Adicionando elementos no slice
	numeros = append(numeros, 10)
	numeros = append(numeros, 120, 30, 40)
	numeros = append(numeros, 50)
	//Criando subslice pegando o índice de 1 a 3
	fmt.Println("Slice antes da modificação: ", frutas)
	subslice := frutas[1:4]
	fmt.Println("Subslice", subslice)
	subslice[0] = "Manga"
	fmt.Println("Slice de frutas após a modificação: ", frutas)
	fmt.Println("\nSlice de números", numeros)
}