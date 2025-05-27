package main

import "fmt"

func welcome() {
	fmt.Println("Bem-vindo ao sistema de filmes!")
}

func createMovie() {
	var name string
	var yearRelease int
	var moviePrice float64

	fmt.Println("Digite o nome do filme:")
	fmt.Scanln(&name)

	fmt.Println("Digite o ano de lançamento:")
	fmt.Scanln(&yearRelease)

	fmt.Println("Digite o valor do filme:")
	fmt.Scanln(&moviePrice)

	fmt.Printf("%s (%d) - R$ %.2f\n", name, yearRelease, moviePrice)
}

func calcularAverage() float64{
	var numRatings int
	fmt.Println("Digite quantas avaliações deseja fazer para o filme!")
	fmt.Scanln(&numRatings)

	var total float64
	for i := 0; i< numRatings; i++{
		var note float64
		fmt.Println("Digite a nota para o filme:")
		fmt.Scanln(&total)
		total += note
	}

	var average float64

	if numRatings > 0 {
		average = total / float64(numRatings) // convertendo numRatings de int para float64
	} else {
		average = 0
	}

	return average
}

func main() {
	welcome()
	fmt.Println("Utilizando Função")
	//createMovie()
	//fmt.Println("\nListe o Próximo")
	//createMovie()
	media := calcularAverage()
	fmt.Printf("A média das avaliações é: %.2f\n", media)
}