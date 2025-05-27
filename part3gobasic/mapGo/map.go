package main

import "fmt"

func main() {
	// Extraindo informaçoes de uma map
	//Criando  um map com nome do aluno como chave e a nota como valor
	estudantes := map[string]float64{
		"Ana": 7.5,
		"Carlos": 5.9,
		"Beatriz": 9.5,
		"João": 6.6,
	}
	fmt.Println("Classificação dos alunos: ", estudantes)
	for nome, nota := range estudantes { // índice, valor
		//fmt.Println(nome)
		//fmt.Println(nota)
		status := "Reprovado"
		if nota >= 6.0 {
			status = ("Aprovado")
		}
		fmt.Printf("Aluno: %s | Nota: %.2f | Status: %s\n", nome, nota, status)
	}
	
}