package main

import (
	"fmt"
	"strings"
)

func main() {
	// Criando um map para contar as palavras
	text := "Golang é uma linguagem de Programação muito rápida, muito veloz"
	words := strings.Fields(text)
	fmt.Println(words)

	wordCount := make(map[string]int) // make() para criar o map estruturado com chave de string  eo valor int

	// Contagem da frequência de palavras
	for _, word := range words {
		wordCount[word]++
	}

	//Imprimir as frequências
	fmt.Println("Contagem de palavras:")
	for word, count := range wordCount {
		fmt.Printf("Palavra %s | Frequência %d\n", word, count)
	}
}
