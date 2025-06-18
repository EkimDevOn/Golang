package main

import (
	"fmt"
	"math"
)

//Acessar o número de Pi
func acessarPi() {
	fmt.Printf("PI arredondado (2 casas decimais): %.2f\n", math.Pi)
}

//Arredondando números para cima e para baixo
func arredondar() {
	num := 10.4
	fmt.Println("Arredondando para cima: ", math.Ceil(num))
	fmt.Println("Arredondando para baixo: ", math.Floor(num))
}

func main() {
	acessarPi()
	arredondar()

	fmt.Println("Potência de 5 elevado a 5: ", math.Pow(5, 5))
	fmt.Println("Raiz quadrada de 169: ", math.Sqrt(169))
	fmt.Println("Logartimo de 10: ", math.Log(10))
	
}