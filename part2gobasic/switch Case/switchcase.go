package main

import "fmt"

func main() {
	var nota int

	fmt.Println("Digite sua nota: ")
	fmt.Scan(&nota)

	switch{
	case nota >= 9 && nota <= 10:
		fmt.Printf("Excelente sua nota é: %d", nota)
	case nota >=7 && nota <= 10:
		fmt.Println("Parabéns sua nota é: ", nota)
	case nota >=5 && nota <= 10:
		fmt.Println("Razoável sua nota é: ", nota)
	case nota >=3 && nota <= 10:
		fmt.Println("Estude mais, sua nota é: ", nota)
	case nota >=0 && nota <= 10:
		fmt.Println("Seu burro!! Sua nota é: ", nota)
	default:
		fmt.Println("Nota inválida! digite uma nota entre 0 e 10")
	}
}