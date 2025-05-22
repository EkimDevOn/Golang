// estratégia para concatenação de valores uma forma mais estruturada com um pacote de manipulaçao de strings
package main

import (
	"fmt"
	"strings"
)

func main() {

	var nome string
	var altura float64
	var idade int

	fmt.Println("Insira seu nome:")
	fmt.Scan(&nome)

	fmt.Println("Insira sua altura:")
	fmt.Scan(&altura)

	fmt.Println("Digite sua idade:")
	fmt.Scan(&idade)

	maiorDeIdade := idade >= 18
	menorDeIdade := idade < 18

	var sb strings.Builder

	fmt.Println("\nDados pessoais:\n")
	sb.WriteString(fmt.Sprintf( "Nome: %s\n", nome))
	sb.WriteString(fmt.Sprintf("Altura: %.2f\n", altura))
	sb.WriteString(fmt.Sprintf("Idade: %d\n", idade))
	sb.WriteString(fmt.Sprintf("Maior de idade? %v\n", maiorDeIdade))
	sb.WriteString(fmt.Sprintf("Menor de idade? %v\n", menorDeIdade))

	fmt.Println(sb.String())
}
