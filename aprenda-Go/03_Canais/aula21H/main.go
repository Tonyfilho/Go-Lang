package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)
/*Então imagine q temos um trabalho para fazer onde NÓS determinamos quantas Threads ou GoRoutine irá fazer o mesmo trabalho
* esta é a moral deste exercicio.
**/
/**01º Começamos com 1 canal e terminaremos com 1 canal, ate aqui não temos divergencia*/
/**02º Madamos X numeros para 1º canal Func: Manda*/
/**03º Pega cada numero do 1º canal  e manda 10 Thread ou 10(GoRoutine) que abastece o canal02
* ou seja temos 10 itens no canal01 e com range criaremos 10 Goroutines ou threads no canal02.
* Ou seja dividimos ou espalhamos o TRABALHO do processador em 10 GoRoutines e todas foram processadas 
* concorrentemente e pois colocamos tudo no canal02
  */
func main() {
	fmt.Println("#### Aula21H   Canais Divergencia Exemplo 01 ")
	fmt.Println()
	fmt.Println("# 1. Um stream vira centenas de go funcs que depois convergem.")
	fmt.Println()
	// 01 Dois canais.
	canal01 := make(chan int)
	canal02 := make(chan int)

	go manda(10, canal01)
	go outra(canal01, canal02)

	// 05  Por fim, range canal dois demonstra os valores.
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

// 03 Outra func faz um range deste canal, e para cada ítem lança uma go func
// que poe o retorno de trabalho() no canal dois.
/*Como dependemos do canal01 da função Manda(), precisamos de Wait.Group*/
func outra(canal1, canal2 chan int) {
var wg  sync.WaitGroup
	for v := range canal1 {
		/*para cada uma iteração um Add()*/
		wg.Add(1)
		go func(x int) {
			canal2 <- trabalho(x)
			/*para cada um Done()*/
			wg.Done()
		}(v)
	}
	// aqui o canal02 recebeu 10 GoRoutine ou Threads
	/**aqui antes do close um Wait()*/
	wg.Wait()
	close(canal2)
}

// 04 Trabalho() é um timer aleatório pra simular workload.
func trabalho(n int) int {
	/**time de um tempo de 0s a 1s*/
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(1e3)))
	return n * 10

}
