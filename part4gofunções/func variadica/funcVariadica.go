// func variádico

package main

import "fmt"

// 1 - função para soma de números com variádicos
func sum(nums ... int){
	sumTotal := 0
	for _, n := range nums {
		sumTotal += n
	}
	fmt.Printf("Soma é: %d\n", sumTotal)
}

// 2 - func para apresentação com variádicos
 func presentation(data map[string]string){// map[key]value
	for key, value := range data {
		fmt.Printf("%s - %s\n", key, value)
	}
 } 

func main() {
	sum(1, 2, 45, 50)

	presentation(map[string]string{
		"\nName": "Golang",
		"Category": "FullStack",
		"Level": "Novice",
	})
	presentation(map[string]string{
		"\nName": "C",
		"Category": "FullStack",
		"Level": "Intermediate",
	})
}