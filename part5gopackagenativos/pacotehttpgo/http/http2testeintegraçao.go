package main

//aqui usaremos http para consumir as informaçoes e testa integrações
import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	URL := "https://jsonplaceholder.typicode.com/todos/2" //endpoint público pra test, 1/ 2/ 3/
	resp, err := http.Get(URL)
	if err != nil {
		fmt.Println("Erro ao fazer a requisição", err)
		return
	}
	defer resp.Body.Close()

	//Usando io.ReadAll() para ler o corpo da reposta
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Erro ao ler o corpo da requisição", err)
		return
	}

	fmt.Println("Resposta da API ", string(body))
}
