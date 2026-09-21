package main

import (
	"fmt"
	"sync"
)
/**O WaitGroup funciona como um contador que precisa chegar a ZERO para liberar o Wait().*/
var wg sync.WaitGroup

func main() {
	fmt.Println("#### Aula21 D  Canais – 3. Range e close")
	fmt.Println(" O Close é usado para indicar que não há mais valores a serem enviados em um canal\n. Ele é útil quando você deseja sinalizar aos receptores que não haverá mais dados a serem recebidos.")
	fmt.Println("Temos q usar o WaitGroup para esperar as GoRoutines terminarem")
	fmt.Println()
	fmt.Println("Sem o WaitGroup, as GoRoutines são executadas em paralelo, mas o main não espera por elas")
	fmt.Println()


	/**Mostraremos o exemplo de range e close, onde a falta do close pode causar um deadlock**/
	fmt.Println()
	fmt.Println("Mostraremos o exemplo de range e close, onde a falta do close pode causar um deadlock")
	/*1º Cria um canal**/
	canal := make(chan int)

	/**Mostrando o exemplo com close, temos a saida normal*/

	wg.Add(2) // 1️⃣ ADD: Incrementa o contador (diz: "tem mais 2 tarefas a serem concluídas")
	go meuLoopComClosed(10, canal)
	go printarValores(canal)
	// ⚠️ O programa CHEGA AQUI e termina imediatamente!
	// As goroutines estão executando em paralelo, mas o main não espera por elas

	wg.Wait() // 3️⃣ WAIT: Bloqueia até o contador chegar a ZERO
	fmt.Println("Programa finalizado!")


}



/*2º Cria um loop q recebe os valores do canal de saída**/
func meuLoopComClosed(total int, canal chan<- int) {
	defer wg.Done() // 2️⃣ DONE: Decrementa o contador (diz: "terminei uma tarefa")
	for i := 0; i < total; i++ {
		canal <- i
	}
	/**É aqui que evitamos o deadlock */
	/**Lê valores enquanto o canal estiver aberto*/
	/**Quando o canal é fechado com close(), o range termina*/
	/**Se o canal não for fechado, o range espera para sempre → deadlock!*/
	close(canal) 
}

/**Opcionamente podemos criar uma canal de entrada**/
func printarValores(localCanal <-chan int) {
	defer wg.Done() // 2️⃣ DONE: Decrementa o contador (diz: "terminei uma tarefa")
	for i := range localCanal {
		fmt.Println("Canal de entrada: ", i)
	}
}
