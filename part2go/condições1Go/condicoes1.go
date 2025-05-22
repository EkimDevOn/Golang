package main

import "fmt"

func main() {
	var idade int

	fmt.Println("Digite sua idade:")
	fmt.Scan(&idade)

	if idade >= 18 {
		fmt.Printf("Você é maior de idade: %d", idade)
	}else{
		fmt.Printf("Você é menor de idade: %d", idade)
	}
}