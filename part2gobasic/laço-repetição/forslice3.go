package main

import "fmt"

func main() {
	notas := []float64{8.5, 9.5, 7.5, 6.8, 9.2}

	fmt.Println(notas)
	fmt.Println("\nCalculando media de valores dentro de uma  slice")
	//calculando média
	var total float64 // variavél global nao preciso utilizar o ":="
	for _, notas := range notas {
		total += notas // total = total + notas
	}

	media := total / float64(len(notas)) // total / pelo quantidade das notas da slice  no caso / 5
	fmt.Println("Soma das notas: ", total)
	fmt.Println("\nMedia das notas: ", media)

}