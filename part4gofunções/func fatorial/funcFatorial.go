// 2 -> 2 * 1
// 3 -> 3 * 2 * 1
// 4 -> 4 * 3 * 2 * 1
package main

import "fmt"

func fatorial (num int) int {
	if num == 1 { // ponto de partida pra calcular o fatorial chegando a 1
		return 1
	}else {
		return num * fatorial(num-1)
	}
}

func totalSum(num int) int {
	if num == 1 {
		return 1
	} else {
		return num + totalSum(num-1) // 5 + 4 + 3 + 2 + 1
	}
}

func main() {
	var number int
	fmt.Println("Digite o número para o fatorial:")
	fmt.Scan(&number)
	fmt.Printf("O fatorial de %d é %d\n", number, fatorial(number)) // placeholder

	var num int
	fmt.Println("Digite um número para a soma:")
	fmt.Scan(&num)
	fmt.Printf("A soma de %d é: %d\n", num, totalSum(num))
}