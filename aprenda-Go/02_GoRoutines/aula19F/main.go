package main

import (
	"fmt"
	"runtime"

	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("Aula 19F - Concorrência –  6. Atomic")
	fmt.Println("O pacote `atomic` fornece primitivas de memória atômica \n de baixo nível úteis para implementar algoritmos de sincronização\n. Ele permite operações atômicas em variáveis compartilhadas entre goroutines, garantindo que essas operações sejam realizadas de forma segura e sem condições de corrida.")
	/*teremos 2 funções principais atomic.LoadInt64 e atomic.StoreInt64*/
	/**Repetiremos o código da aula anterior sem o Mutex*/

	/**Repetiremos o codigo da aula anterior*/

	fmt.Println()

	fmt.Println()
	sharedVariable := int64(0)
	numberOfGoroutines := 5
	wg.Add(numberOfGoroutines)

	for i := 0; i < numberOfGoroutines; i++ {

		go func() {
			/**Iniciamos aqui o atomics e passaremos um Ponteiro para a variável compartilhada*/
			/**invocando o construtor atômica*/
			/**unc atomic.AddInt64(addr *int64, delta int64) (new int64)*/
			localAtomic := atomic.AddInt64(&sharedVariable, 1)
			runtime.Gosched() // Forçando a troca de contexto para simular concorrência
			/**Imprimiremos o atômico*/
			fmt.Println("Shared Variable:\t", localAtomic)

			wg.Done()
		}() //Temos que por () para teremos a execução , criamos uma go func() anônima.

	}
	/**No final do programa ou do bloco de código teremos que esperar todas as goroutines terminarem*/
	wg.Wait()
	fmt.Println("Valor final da variável compartilhada:", sharedVariable)
	/* O valor final da variável compartilhada deve ser 5, pois temos 5 goroutines incrementando a variável compartilhada. */
	fmt.Println()

}
