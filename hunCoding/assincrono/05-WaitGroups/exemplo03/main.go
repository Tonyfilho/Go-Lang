package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("Aula05 de GoRoutines e  WaitGroup.")
	fmt.Println("Com GoRoutines e WaitGroup.")
	fmt.Println()
	/**Criaremos uma var WG que recebe o construtor do WaitGroup*/
	var wg sync.WaitGroup

	/**Ao invez de repetirmo 5x criamos um FOR*/
	for i := 0; i < 5; i++ {
		/**Para cada Item ou seja para cada loop daremos um Add(1)*/
		/** 1 è numero que tem de esperar, ou seja 1 goRoutine*/
		wg.Add(1)
		/**Criaremos um função anomima onde dentro teremos defer e o Done()*/
		go func() {
			/**Quer dizer assim, que esta GoRoutine quando terminar não vai precisa mais de que o Main() espere*/
			defer wg.Done()
			DemoraParaExecutar(2 * time.Second)
		}()

	}
	/**Agora para avisar o Go q temos GoRoutines, senão da acontece*/
	/**O Wait() dizem: Todos os numeros colocado no Add(1), eu vou segurar o Main(), ate estes acabarem*/
	wg.Wait()

	fmt.Println("O Main() terminou.....")

	/**Ou Seja executou as 5 de uma vez e terminou as 5 de uma vez*/
	// 	Começando a Execução. Salvando no Banco
	// Começando a Execução. Salvando no Banco
	// Começando a Execução. Salvando no Banco
	// Começando a Execução. Salvando no Banco
	// Começando a Execução. Salvando no Banco

	/**Tudo terminou em 2 Segundos, pois tu começo junto e terminou junto*/
	// Terminou a Execução. Dados Salvo no Banco
	// Terminou a Execução. Dados Salvo no Banco
	// Terminou a Execução. Dados Salvo no Banco
	// Terminou a Execução. Dados Salvo no Banco
	// Terminou a Execução. Dados Salvo no Banco
}

func DemoraParaExecutar(t time.Duration) {
	fmt.Println("Começando a Execução. Salvando no Banco")
	fmt.Println()
	time.Sleep(t)
	fmt.Println("Terminou a Execução. Dados Salvo no Banco")
	fmt.Println()

}
