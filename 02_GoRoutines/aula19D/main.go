package main

import (
	"fmt"
	"runtime"
	"sync"
)

/**Criando a variável de espera*/
var wg sync.WaitGroup

func main() {
	fmt.Println("Aula 19D - Concorrência – 4. Na prática: Run Condition ou condição de corrida")
	fmt.Println()
	/**teremos um exemplo prático de condição de corrida, onde compartilharemos uma variável entre goroutines*/
	sharedVariable := 0

	/**vamos mediar quantos processadores estão disponíveis*/
	numOfCpu := runtime.NumCPU()
	fmt.Println("Número de CPUs disponíveis:", numOfCpu)
	/**vamos mediar quantas goroutines usadas*/
	fmt.Println("Número de Go Routines Antes:", runtime.NumGoroutine())

	numberOfGoroutines := 5

	wg.Add(numberOfGoroutines)

	/**Criaremos uma função para disparar as goroutines*/
	for i := 0; i < numberOfGoroutines; i++ {
		/**lançando a goroutine*/
		go func() {
			/*teremos uma variável para armazenar o valor da variável compartilhada*/
			v := sharedVariable
			/**faremos um Yield usaremos runtime.Gosched()*/
			runtime.Gosched()
			fmt.Println("Valor da variável compartilhada: Antes do incremento", v)
			v++
			//	fmt.Println("Valor da variável compartilhada:", v)
			/**salvando o valor da variável compartilhada*/
			sharedVariable = v
			/**indicando que a goroutine terminou*/
			wg.Done()
		}() //Temos que por () para teremos a execução , criamos uma go func() anônima.
		fmt.Println("Número de Go Routines dentro do loop:", runtime.NumGoroutine())

	}
	/**No final do programa ou do bloco de código teremos que esperar todas as goroutines terminarem*/
	wg.Wait()
	fmt.Println("Número de Go Routines No Final:", runtime.NumGoroutine())
	fmt.Println("Valor final da variável compartilhada:", sharedVariable)
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println("Tivemos 5 goroutines incrementando a variável compartilhada\n, mas o valor final da variável compartilhada é ", sharedVariable, ", pois tivemos uma condição de corrida.")
	fmt.Println()
	fmt.Println("Resulmo vais goroutines leram uma variavel compartilhada e não conseguiram salvar na variavel")
	fmt.Println("Para evitar a condição de corrida, podemos usar o Mutex do pacote sync, que é uma estrutura que permite bloquear o acesso a uma variável compartilhada entre goroutines.")
}
