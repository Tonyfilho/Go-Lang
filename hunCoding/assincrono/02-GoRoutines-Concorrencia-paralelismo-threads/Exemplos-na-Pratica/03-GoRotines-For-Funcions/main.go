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
	fmt.Println("GoRotines-For-Funcions")
	printMessage("Sem GoRoutine")
	/**Ilustraremos 2 GoRoutines rodando em concorrência*/

	go printMessage("GoRoutine 1")
	/**Recebendo a função como variável*/
	fgx := printMessage
	go fgx("GoRoutine 2")

	/*Veja  ordem de saida na tela, ela são concorrentes e não se conhecem, por isto ficará assim */
	/**Da mesma forma que se não pormos um tempo de espera a mesma não seria executada*/
	// GoRoutine 2
	// GoRoutine 1
	// GoRoutine 1
	// GoRoutine 2
	// GoRoutine 2
	// GoRoutine 1
	time.Sleep(5 * time.Millisecond)

	/*Veremos mais para frente como o scheduler do Go trabalha com GoRoutines*/

	fmt.Println("Fim do programa")

}
