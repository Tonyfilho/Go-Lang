package main

import (
	"fmt"
	"time"
)

/**Uma função simples que imprime uma mensagem e simula um atraso*/

func printMessage(message string) {
	for i := 0; i < 3; i++ {
		fmt.Println(message)
		time.Sleep(1 * time.Millisecond) // Simula um atraso de 1 segundo
	}
}

func main() {
	fmt.Println("GoRotines-Anonymous-Call")
	printMessage("Sem GoRoutine")
	/**Ilustraremos 2 GoRoutines rodando em concorrência*/

	go printMessage("GoRoutine 1")
	/**Recebendo a função como variável*/
	fgx := printMessage
	go fgx("GoRoutine 2")

	/**Função anônima*/
   go func() {
		printMessage("GoRoutine 3 Func")
	}()



	time.Sleep(5 * time.Millisecond)


	fmt.Println("Fim do programa")

}