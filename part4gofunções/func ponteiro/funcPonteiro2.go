package main

import "fmt"

type Cliente struct {
	Nome  string
	Idade int
}

func atualizarCliente(c *Cliente, NovoNome string, NovaIdade int) {
	c.Nome = NovoNome
	c.Idade = NovaIdade
}

func main() {
	cliente := Cliente{Nome: "João", Idade: 28}
	fmt.Println("Antes da alteração: ", cliente)
	atualizarCliente(&cliente, "Calos", 38)
	fmt.Println("Após a alteração: ", cliente)
}
