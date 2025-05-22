package main

import "fmt"

func main() {
	var numeros[5]int
	var frutas = [3]string{"Maça", "Banana", "Laranja"}
	var notas = [5]float64{8.75, 4.5, 9.8, 7.6, 3.9}

	soma := notas[0] + notas[1] + notas[2] + notas[3] + notas[3]
	//fmt.Println(Len(notas))
	media := soma / float64(len(notas))

	numeros[0] = 10
	numeros[1] = 20
	numeros[2] = 30
	numeros[3] = 40
	numeros[4] = 50

	fmt.Println("\n Arrays de números: ", numeros)
	fmt.Println("Array de frutas: ", frutas[0:2])
	fmt.Println("\n Soma das notas é: ", soma)
	fmt.Printf("Média das notas: %.2f", media)
}