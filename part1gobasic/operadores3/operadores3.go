package main

import "fmt"

func main() {
	var num1, num2 int

	fmt.Println("Informe o parimeiro número:")
	fmt.Scan(&num1)

	fmt.Println("Digite o segundo numero:")
	fmt.Scan(&num2)

	soma := num1 + num2
	subtracao := num1 - num2
	multiplica := num1 * num2
	divisao := num1 / num2
	modulo := num1 % num2

	fmt.Printf("Soma de %d e %d é %d\n", num1, num2, soma)
	fmt.Printf("Subtração de %d e %d, é %d\n", num1, num2, subtracao)
	fmt.Printf("Multiplicação de %d e %d, é %d\n", num1, num2, multiplica)
	fmt.Printf("Divisão de %d e %d, é %d\n", num1, num2, divisao)
	fmt.Printf("Modúlo de %d e %d, é %d\n", num1, num2, modulo)
}
