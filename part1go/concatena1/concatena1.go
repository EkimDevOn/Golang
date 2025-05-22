// estratégia para concatenação de valores uma forma mais estruturada com um pacote de manipulaçao de strings
//
package main

import ("fmt"
		//"strings"
)

func main() {

	var nome string
	var altura float64
	var idade int
	
	fmt.Println("Insira seu nome:")
	fmt.Scan( &nome)

	fmt.Println("Insira sua altura:")
	fmt.Scan(&altura)

	fmt.Println("Digite sua idade:")
	fmt.Scan( &idade)

	maiorDeIdade := idade >= 18
	menorDeIdade := idade < 18

	fmt.Println("\nDados pessoais:\n")
	fmt.Printf("Nome: %s\n", nome)
	fmt.Printf("Altura: %.2f\n", altura)
	fmt.Printf("Idade: %d\n", idade)
	fmt.Printf("Maior de idade? %v\n", maiorDeIdade)
	fmt.Printf("Menor de idade? %v\n", menorDeIdade)
}
