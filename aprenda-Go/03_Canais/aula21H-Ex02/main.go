package main

import (
	"fmt"

	"sync"
	"time"
)
/*Resulmo da Lição: 
01º cada Trabalho leva 1S
02º e criamos 5 GoRoutines que no daremos resposta
03º ou seja no terminal teremos 5 em 5, veja.
*Ou seja na Func Outra, criar grupos de 5 goRoutines para resolver 1 Trabalho por vez
* isto prova que podemos criar um grupo de GoRoutines para execultar 1 trabalho concorrentemente.
* Imagine que temos 1 stream e dividimos o trabalho desta stream em varias Treads ou GoRoutine
*/

/*
* Ídem acima, mas a func que lança go funcs é assim:
* Na função Outra Cria X go funcs, cada uma com um range no primeiro canal que,
* para cada item, poe o retorno de trabalho() no canal dois.
*/
func main() {
	fmt.Println("#### Aula21H   Canais Divergencia Exemplo 01 ")
	fmt.Println()
	fmt.Println("# 1.Criamos uma Função que recebe o X numeros de GoRoutines que irão funcionar para resolver um Trabalho, neste ex: 05 por Trabalho.")
	fmt.Println()

	canal01 := make(chan int)
	canal02 := make(chan int)
	totalGoFunc := 5

	go manda(20, canal01)
	go outra(totalGoFunc, canal01, canal02)

	for v := range canal02 {
		fmt.Println("Range de Canal02: ", v)
	}

}

// 02  Uma func manda X números ao primeiro canal.
func manda(n int, canal chan int) {
	for i := 0; i < n; i++ {
		canal <- i
	}
	close(canal)
}

func outra(totalGoF int, canal1, canal2 chan int) {
	var wg sync.WaitGroup
	// Cria X go funcs, cada uma com um range do canal01 que, para cada item, poe o retorno de trabalho() no canal02

	for i := 0; i < totalGoF; i++ {
		wg.Add(1)
		go func() {
			for v := range canal1 {
				canal2 <- trabalho(v)
			}
			wg.Done()
		}()

	}
	wg.Wait()
	close(canal2)

}

func trabalho(n int) int {
	/**Cada trabalho leva 1S para execução*/
	time.Sleep(time.Second)
	return n * 10

}
