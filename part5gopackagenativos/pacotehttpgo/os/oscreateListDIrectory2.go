package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

// 1 retornar a pasta atual
func getCurrentDirectory() string {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Erro ao obter o diretório: ", err)
		return ""
	}
	return dir
}

// 2 Listar arquivos e pastas
func listFilesAndDirectories() {
	files, err := os.ReadDir(".") // ".", vai ler tudo que esta dentro do diretóro
	if err != nil {
		fmt.Println("Erro ao ler o diretório: ", err)
		return
	}
	for _, file := range files { // trás nome do arquivo
		fmt.Println(file.Name())
	}
}

// 3 Versão SO, no cmd
func getOsVersion() {
	var cmd *exec.Cmd

	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/C", "ver")
	} else if runtime.GOOS == "Linux" {
		cmd = exec.Command("uname", "-a")
	} else if runtime.GOOS == "darwin" {
		cmd = exec.Command("s2_vers")
	} else {
		fmt.Println("SO não suportado!")
	}
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Erro ao executar comando: ", err)
		return
	}
	fmt.Println(string(out)) // converto pra string e chamo meu out
}

// 4 Configuração da máquina
func getSystemInfo() {
	var cmd *exec.Cmd

	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/C", "systeminfo")
	} else {
		fmt.Println("Comando 'systeminfo' não disponível neste SO")
	}
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Erro ao executar comando: ", err)
		return
	}
	fmt.Println(string(out))
}

// 5  delisgar o computador em 1 hr
func shutdownInOneHour(){
	cmd := exec.Command("shutdown", "/s", "/t", "3600")
	err := cmd.Run()
	if err != nil {
		fmt.Println("Erro ao agendar o desligamento")
	}
}

// 6 cancelar o desligamento
func cancelShutdownI(){
	cmd := exec.Command("shutdown", "/a")
	err := cmd.Run()
	if err != nil {
		fmt.Println("Erro ao cancelar o desligamento")
	}
}

func main() {
	fmt.Println("Pasta atual: ", getCurrentDirectory())
	fmt.Println("Arquivos e pastas:")
	listFilesAndDirectories()

	fmt.Println("Versao do SO: ")
	getOsVersion()

	fmt.Println("configuração da Máquina: ")
	getSystemInfo()

	//shutdownInOneHour()

	//cancelShutdownI()
}
