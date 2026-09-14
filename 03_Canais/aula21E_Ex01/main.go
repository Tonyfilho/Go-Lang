package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("#### Aula21E   Canais – 4. Select # Exemplo 1:")
	fmt.Println()
	fmt.Println("O que temos de aprender q podemos ter uma FUNÇÂO com SELECT que recebe valores de varios CANAIS")
	fmt.Println()
	fmt.Println("Select é como switch, só que pra canais, e não é sequencial.")
	fmt.Println()
	fmt.Println("Ele vai esperar até que algum canal esteja pronto para enviar ou receber dados\n, e então executa o case correspondente.")
	fmt.Println()
	fmt.Println("Duas go funcs enviando X/2 numeros cada uma pra um canal")
	fmt.Println()

	canalA := make(chan int)
	canalB := make(chan int)
	x := 50
	/**Temos q por sync.Add, sync.wait e sync.done, senão o main termina antes das goroutine*/
	var wg sync.WaitGroup
	wg.Add(2) /**01 temos 2 Goroutines temos q ter 2 no Add*/

	/**1ª Função*/
	go func(localNumero int) {
		defer wg.Done()
		//wg.Done() // 02 Quando terminar, decrementa o contador Ex:   contador: 1 → 0
		for i := 0; i < localNumero; i++ {
			canalA <- i
		}
		close(canalA) // Fecha canalA quando terminar
	}(x / 2) // ⚠️ Executa em paralelo com a outra goroutine

	/**2º Função onde teremos o SELECT CASE*/
	/**O quetemos de aprender q podemos ter uma FUNÇÂO que recebe valores de varios CANAIS**/
	go func(localNumero int) {
		defer wg.Done()
	//	wg.Done() // 02 Quando terminar, decrementa o contador
		for i := 0; i < localNumero; i++ {
			/**Aqui teremos os cases */
			select {
			/**Se eu receber do canalA faço algo Ex: chamo a API da europa**/
			case v := <-canalA:
				fmt.Println("Sou do canalA: ", v)

				/**Se eu receber do canalB faço algo Ex: chamo a API da america**/
			case v := <-canalB:
				fmt.Println("Sou do canalB: ", v)

			}
		}
		close(canalB) // Fecha canalA quando terminar
	}(x / 2) // ⚠️ Executa em paralelo com a outra goroutine

	// 03 Espera ambas as goroutines terminarem
	wg.Wait()
	fmt.Println("✅ Programa finalizado!")

}
