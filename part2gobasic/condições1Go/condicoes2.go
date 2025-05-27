package main

import "fmt"

func main() {
	var nota float64

	fmt.Println("Digite sua nota(0 a 10)")
	fmt.Scan(&nota)

	if nota < 0 || nota > 10 {
		fmt.Println("Nota inválida! Digite uma nota entre 0 e 10")
	}else if nota <= 10 && nota >= 5 {
		 fmt.Print("Parabéns você passou de ano, média: ", nota)
	}else {
		fmt.Println("Você esta reprovado: ", nota)
	}
}