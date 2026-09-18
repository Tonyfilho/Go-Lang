package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Aula 19 Especial - Entendendo Go Routines")
	fmt.Println()
	fmt.Println("Aqui entenderemos que: Enquando não for usando o valor que está em uma GoRoutines\n, não será atribuido um novo valor")
	fmt.Println()

	/*01º *Criamremos uma Fila onde teremos um Canal de numeros inteiros*/
	queue := make(chan int)

	/*02º Criremos um Goroutine dentro de uma função, onde teremos um FOr e um Time, e atribuiremos o canal **/
	/*02B O Valor de "I" que será  Atribuido, só entra na variavel se o valor anterior for CONSUMIDO**/

	fmt.Println("02B O Valor de I que será  Atribuido, só entra na variavel se o valor anterior for CONSUMIDO")
	fmt.Println()
	/***Com uso de Lupe Infinito não se usa o Close() */

	go func() {
		i := 0
		for {
			/*Se ninguem usar esta Variavel QUEUE em outra GoRoutine, a tribuição fica Bloqueada**/
			/*Ou seja o FOR não funciona**/
			queue <- i
			i++
			time.Sleep(time.Second)
		}
	}()
	// 3. Consumidor: lê valores e para quando chegar a 20
	// ⚠️ NÃO feche o canal! Quem envia é quem fecha, e essa goroutine é infinita
	fmt.Println("O for ficará párado enquato um outra GoRoutine  não usa o valor dentro de Queue")
	fmt.Println()
	lerGoRoutines(queue)
	fmt.Println()	
	fmt.Println("✅ Programa finalizado!")

}

func lerGoRoutines(a chan int) {
	for x := range a {
		if x < 20 {
			fmt.Println("Canal Usado e Lido: ", x)

		} else {
			fmt.Println("🛑 Limite atingido, saindo...")
			return
		}
	}
}
