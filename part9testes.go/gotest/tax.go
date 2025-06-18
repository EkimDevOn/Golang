// go mod init #NOME#, para iniciar um mod testes
// tax - realizando cálculo de uma valor específico, taxa
// package tax
package tax

func CalculateTax(amount float64) float64 {
	if amount >= 1000 {
		return 10.0
	} else {
		return 5.0
	}
}
