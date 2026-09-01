package main

import (
	"fmt"
	"runtime"
	"sync"
)

var mutex sync.Mutex /**teremos 2 métodos , o lock e o unlock*/
var wg sync.WaitGroup

func main() {
	fmt.Println("Aula 19E - Concorrência – 5. Mutex")
	fmt.Println("Mutex é mutual exclusion, exclusão mútua.")
	/**Utilizando mutex somente uma thread poderá utilizar
	* a variável contador de cada vez, e as outras deve aguardar sua vez "na fila.
	 */
	fmt.Println("Utilizando mutex somente uma thread poderá utilizar a variável contador de cada vez\n, e as outras deve aguardar sua vez na fila de espera.	")
	/**O Mutex traca um trecho de codigo que deve ser executado de forma segura */

	/**Repetiremos o codigo da aula anterior*/
	numOfCpu := runtime.NumCPU()
	fmt.Println("Número de CPUs disponíveis:", numOfCpu)
	fmt.Println()
	fmt.Println("Número de Go Routines Antes:", runtime.NumGoroutine())
	fmt.Println()
	sharedVariable := 0
	numberOfGoroutines := 5
	
	wg.Add(numberOfGoroutines)
	
	for i := 0; i < numberOfGoroutines; i++ {
		
		go func() {
			/**1º Colocaremos o mutex Lock antes de acessar a variável compartilhada*/
			mutex.Lock()
			v := sharedVariable
			runtime.Gosched()
			fmt.Println("Valor da variável compartilhada: Antes do incremento", v)
			v++
			sharedVariable = v
			/**2º Colocaremos o mutex Unlock antes de terminar*/
			mutex.Unlock()
			
			wg.Done()
			}() //Temos que por () para teremos a execução , criamos uma go func() anônima.
			fmt.Println("Número de Go Routines dentro do loop:", runtime.NumGoroutine())
			
		}
		/**No final do programa ou do bloco de código teremos que esperar todas as goroutines terminarem*/
		wg.Wait()
		fmt.Println("Número de Go Routines No Final:", runtime.NumGoroutine())
		fmt.Println("Valor final da variável compartilhada:", sharedVariable) 
		/* O valor final da variável compartilhada deve ser 5, pois temos 5 goroutines incrementando a variável compartilhada. */
		fmt.Println()
		fmt.Println("Agora sim, temos o valor final da variável compartilhada correto\n, pois utilizamos o mutex para garantir que apenas uma goroutine acessasse a variável compartilhada de cada vez.")

}
