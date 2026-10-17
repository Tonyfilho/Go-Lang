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
	printMessage("Mensagem sem GoRoutines")
	/*usando assim de forma básica temos a mesma saida do exemplo anterior*/
	/*as GoRoutines executam concorrentemente, elas não si conhecem aqui neste caso por isto q temos esta comportamento*/
	/*Quando vc executa uma goRoutine dentro da Função Main que tb é uma goRoutine, ela só vai esperar a conclusão da execução se fizermos algo*/
	/**Ou Seja se usarmos waitGroup, chanels e outros*/
	/*GoRoutines Function Call Não vai imprimir a mensagem, pois o programa principal termina antes de executar a função printMessage*/
	go printMessage("Mensagem com GoRoutines")
	/**A função Main vai chamar o codigo assim, printar 3x e depois vai falar com a GoRoutine, lamento mas não tenho mais nada para fazer tenho que fechar*/
	/**Ou seja ninguem me dou ordem de espera e então vou finalizar o programa*/

	/**como por fazer neste caso? podemos dar um Delay ou usar waitGroup, nesta caso ficaremos com o delay de um time.Sleep(2 * time.Second)*/
	time.Sleep(5 * time.Millisecond) // Simula um atraso de 2 segundos para dar tempo da GoRoutine executar
/*O Grande desafio é fazer com que um GoRoutine espere pela outra e tb passe a conhecê-la uma a outra */
	fmt.Println("Fim do programa")
}
