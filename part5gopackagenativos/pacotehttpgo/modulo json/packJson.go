package main

import (
	"fmt"
	"encoding/json"
)

type Person struct {
	Name string `json:"name"`
	Age int `json:"age"`


}

func main() {
	jsonString := `{"name":"Romário", "age":30}`
	var person Person
	json.Unmarshal([]byte(jsonString), &person)
	fmt.Printf("Nome: %s, Idade: %d\n", person.Name, person.Age)

	person.Name = "Fulano"
	person.Age = 50

	jsonData, _ := json.Marshal(person)
	fmt.Println(string(jsonData))

	json.Unmarshal([]byte(jsonData), &person)
	fmt.Printf("Nome: %s, Idade: %d\n", person.Name, person.Age)
}
