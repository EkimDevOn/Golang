// Automatizar tarefas no sistema operacional os -operating system

package main

// Imagine que queremos trabalhar com a criação de arquivo de texto como exemplo usando o package os
//e com a leitura de arquivo de texto
import (
	"fmt"
	"os"
)

func main() {
	// Criando um arquivo
	arquivo, err := os.Create("exemplo.txt")
	if err != nil {
		fmt.Println("Erro ao criar arquivo: ", err)
		return
	}

	defer arquivo.Close()

	//lendo um arquivo
	conteudo, err := os.ReadFile("exemplo.txt")
	if err != nil {
		fmt.Println("Erro ao ler arquivo: ", err)
		return
	}

	fmt.Println("Conteudo do arquivo: ", string(conteudo))
}
