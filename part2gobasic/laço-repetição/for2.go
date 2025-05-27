package main

import "fmt"

func main() {
	nomes := [4]string{"João", "Maria", "Gulherme", "Pedro"}
	fmt.Println(nomes)

	for id, nome := range nomes {
		//fmtprintln("Id: ", id)
		//fmt.Println("Nome: ", nome)
		fmt.Println(id+1, "->", nome) // +1 começa na posiçao 1. (pula posição 0)
		if len(nome) < 5 {
			fmt.Println(nome, "Tem mais de 5 caracteres")
		}
	}
}
