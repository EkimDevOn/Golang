package main
//argumentos ou parametros
import "fmt"

func fullName(fristName, lastName string){
	fmt.Printf("Nome completo: %s, %s\n", fristName, lastName)
}

func sumNumbers(a, b int) int{ // tenho que colar a funçao dentro de um print pois so esta retornando 
 return a + b
}

//parametro peenchido
func address (country string){
	if country == ""{
		country = "Brasil"
	}
	fmt.Printf("Eu moro no %s\n", country)
}

func main() {
	fmt.Println("Utilizando função com parametros")
	fullName("Rodrigo","Macedo")
	fullName("Fulano","Ciclano")
	fmt.Printf("Soma: %d\n", sumNumbers(10, 50))
	address("")
	address("Canadá")
} 