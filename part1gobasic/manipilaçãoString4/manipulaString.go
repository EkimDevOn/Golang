package main

import (
	"fmt"
	"strings"
)

func main() {
	movieName := "Top Gun"
	movieName2 := "top gun"

	fmt.Println(movieName == movieName2)

	movieDescription := `
	Top Dun marveric é um filme de aviação
	`
	line := "-"
	fmt.Println(strings.Repeat(line, 40))
	//fmt.Println("========================================")
	fmt.Println(movieDescription)

	//Verifica se uma palavra existe dentro de uma string
	fmt.Println(strings.Contains(movieDescription, "Top"))
	fmt.Println(strings.Contains(movieDescription,"Filme"))

	// Convertendo texto para maiúsculo
	fmt.Println(strings.ToUpper(movieDescription))
	//Convertendo texto para minúsculo
	fmt.Println(strings.ToLower(movieDescription))
	// Primeira letra em maiúsculo
	fmt.Println(strings.Title(movieDescription))

	// encontrar a posição de uma caractere
	fmt.Println(strings.Index(movieDescription, "p"))

	// Contando o número de caracteres de uma string
	fmt.Println(strings.Count(movieDescription, "a"))
	fmt.Println(strings.Count(movieDescription, "e"))

	// Substituir um elemento por outro
	fmt.Println(strings.ReplaceAll(movieDescription, "filme", "série")) 
}
