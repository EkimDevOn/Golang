package main

// Basicamente uma é uma variável que gaurda endereço de memória de um outro objeto, manipular dinamicamente execuçao de códigos e outras vantagens a mais
import "fmt"

// Função que recebe um ponteiro como argumento
// e altera o valor da variável original
func alterValue(x *int) {
	//alterando valor do ponteiro
	*x *= 2
}

func main() {
	num := 5
	fmt.Printf("Valor inicial: %d\n", num)
	//Passando o ponteiro para a função
	alterValue(&num)
	fmt.Printf("Valor após a alteração: %d\n", num)
}
