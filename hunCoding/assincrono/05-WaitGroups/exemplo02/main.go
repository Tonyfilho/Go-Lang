package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Aula05 de GoRoutines e  WaitGroup.")
	fmt.Println("Com GoRoutines mas SEM WaitGroup.")
	fmt.Println()
	/**Criaremos um Metodo q gera um delay na execução do codigo**/
	go DemoraParaExecutar(2 * time.Second)
	go DemoraParaExecutar(2 * time.Second)
	go DemoraParaExecutar(2 * time.Second)
	go DemoraParaExecutar(2 * time.Second)
	go DemoraParaExecutar(2 * time.Second)
	
	fmt.Println("Terminou o Main()....")
}

func DemoraParaExecutar(t time.Duration) {
	fmt.Println("Começando a Execução. Salvando no Banco")
	fmt.Println()
	time.Sleep(t)
	fmt.Println("Terminou a Execução. Dados Salvo no Banco")
	fmt.Println()

}