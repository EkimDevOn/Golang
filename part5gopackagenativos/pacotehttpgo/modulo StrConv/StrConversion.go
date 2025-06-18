// Este package transforma tipos de dados converções de dados
package main

import (
	"fmt"
	"strconv"
)

func main() {
	// String para int
	numeroStr := "123"
	numero, err := strconv.Atoi(numeroStr)
	if err != nil {
		fmt.Println("Erro: ", err)
		return
	}
	fmt.Println("Número: ", numero)

	// int para string
	numeroInt := 456
	numeroStr2 := strconv.Itoa(numeroInt)
	fmt.Println("numeroInt2: ", numeroStr2)

	// string para float
	floatStr := "16.45"
	valorFloat, err := strconv.ParseFloat(floatStr, 64) // bit size = 64 é o formato 64, 32...
	if err != nil {
		fmt.Println("Erro especifique tipo de dados: ", err)
		return
	}
	fmt.Println("Valor float: ", valorFloat)
}

// porque nao tem tratamento de erro quando e de int para string :

//Conversão de Int para String

//Quando você converte um inteiro para uma string usando strconv.Itoa,
//a operação sempre será bem-sucedida, desde que o valor seja um inteiro válido.
// Isso porque um inteiro sempre pode ser convertido para uma string que representa esse número.
