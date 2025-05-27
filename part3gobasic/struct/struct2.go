package main

import "fmt"

// Definindo struct, sempre fora da main()
type Pessoa struct {
	Nome string
	Idade    int
	Endereco    string
}

func main() {
	//Criar instância da struct Pessoa
	pessoa1 := Pessoa{
		Nome: "Mariana",
		Idade:    26,
		Endereco:   "Rua la paz 459",
	}
	fmt.Println("Informações da Pessoa:")
	fmt.Printf("Nome: %s\n", pessoa1.Nome)
	fmt.Printf("Idade: %d\n", pessoa1.Idade)
	fmt.Printf("Endereço: %s", pessoa1.Endereco)
}
